package repoimpl

import (
	"context"
	"encoding/json"
	errorGrpc "finalProject/cmd/service/grpc-model/error"
	"finalProject/cmd/service/grpc-model/user"
	"finalProject/cmd/service/user/model"
	"finalProject/cmd/service/user/repository"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)



type UserRepoImpl struct {
	DB *gorm.DB
	RedisDb *redis.Client
}

func (i *UserRepoImpl) Login(ctx context.Context, req *user.LoginReq) (res *user.GetUserRes, err error) {
	userRet := &user.User{}
	res = &user.GetUserRes{}
	p := model.User{}
	notfound := i.DB.Where("email = ?", req.Email).First(&p).RecordNotFound()
	p.Fill(userRet)
	if notfound == true {
		res.Error = &errorGrpc.Error{
			Code: int32(codes.NotFound),
			Message: fmt.Sprintf("Not found user with email= %s", req.Email),
		}
		return res, nil
	}
	hashError := bcrypt.CompareHashAndPassword([]byte(p.Password), []byte(req.Password))
	if hashError != nil || err != nil {
		res.Error = &errorGrpc.Error{
			Code: int32(codes.Unauthenticated),
			Message: "Wrong Password!",
		}
		return res, nil
	}
	res.User = userRet
	res.Error = nil
	return res, nil
}

func (i *UserRepoImpl) DeleteUser(context.Context, *user.DeleteUserReq) (*user.DeleteUserRes, error) {
	panic("implement me")
}

func NewUserRepo(db *gorm.DB, redisDb *redis.Client) repository.UserRepo {
	return &UserRepoImpl{
		DB: db,
		RedisDb: redisDb,
	}
}

func (i *UserRepoImpl) RegisterUser(ctx context.Context, req *user.CreateUserReq) (res *user.CreateUserRes, err error) {
	p := model.User{}
	p.Set(req)
	errInsert := i.DB.Create(&p).Error
	res = &user.CreateUserRes{}
	if errInsert != nil {
		res.Error = &errorGrpc.Error{
			Code: int32(codes.AlreadyExists),
			Message: "Register failed!",
		}
		return res, nil
	}
	userRet := &user.User{}
	p.Fill(userRet)
	//Remove password response
	userRet.Password = ""
	res.User = userRet
	return res, nil
}

func (i *UserRepoImpl) GetUser(ctx context.Context, req *user.GetUserReq) (res *user.GetUserRes, err error) {
	val, err := i.RedisDb.Get(string(req.Id)).Result()
	userRet := &user.User{}
	res = &user.GetUserRes{}
	if err != nil {
		p := model.User{}
		notfound := i.DB.First(&p, req.Id).RecordNotFound()
		p.Fill(userRet)
		if notfound == true {
			return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Not found user with id= %d", req.Id))
		}
		json, err := json.Marshal(user.User{
			Id:       userRet.Id,
			Name:      userRet.Name,
			Email:     userRet.Email,
		})
		err = i.RedisDb.Set(string(req.Id), json, 0).Err()
		if err != nil {
			return nil, err
		}
		fmt.Println("Get from db", res)
		res.User = userRet
		return res, nil
	}
	fmt.Println("Get from redis", res)
	err = json.Unmarshal([]byte(val), &res.User)
	if err != nil {
		return nil, err
	}
	return res, nil
}


