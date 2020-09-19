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
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)



type UserRepoImpl struct {
	DB *gorm.DB
	RedisDb *redis.Client
}


func (i *UserRepoImpl) DeleteUser(context.Context, *user.DeleteUserReq) (*user.DeleteUserRes, error) {
	panic("implement me")
}

func (i *UserRepoImpl) Get(context.Context, *user.GetUserReq) (*user.GetUserRes, error) {
	panic("implement me")
}

func NewUserRepo(db *gorm.DB, redisDb *redis.Client) repository.UserRepo {
	return &UserRepoImpl{
		DB: db,
		RedisDb: redisDb,
	}
}

func (i *UserRepoImpl) CreateUser(ctx context.Context, req *user.CreateUserReq) (res *user.CreateUserRes, err error) {
	p := model.User{}
	p.Set(req)
	errInsert := i.DB.Create(&p).Error
	if errInsert != nil {
		return nil, status.Errorf(codes.AlreadyExists, fmt.Sprint("Cannot add user"))
	}
	userRet := &user.User{}
	p.Fill(userRet)
	res = &user.CreateUserRes{}
	res.User = userRet
	return res, nil}

func (i *UserRepoImpl) GetUser(ctx context.Context, req *user.GetUserReq) (res *user.GetUserRes, err error) {
	val, err := i.RedisDb.Get(string(req.Id)).Result()
	userRet := &user.User{}
	res = &user.GetUserRes{}
	if err != nil {
		p := model.User{}
		notfound := i.DB.First(&p, req.Id).RecordNotFound()
		p.Fill(userRet)
		if notfound == true {
			return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Not found product with id= %d", req.Id))
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
