package repoimpl

import (
	"context"
	"encoding/json"
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
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Not found user with email= %s", req.Email))
	}
	hashError := bcrypt.CompareHashAndPassword([]byte(p.Password), []byte(req.Password))
	if hashError != nil {
		return nil, status.Errorf(codes.Unavailable, fmt.Sprintf("Wrong password!"))
	}
	res.User = userRet
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
	if errInsert != nil {
		return nil, status.Errorf(codes.AlreadyExists, fmt.Sprint("Register failed!"))
	}
	userRet := &user.User{}
	p.Fill(userRet)
	res = &user.CreateUserRes{}
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


