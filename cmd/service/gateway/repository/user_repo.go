package repository

import (
	"context"
	"finalProject/cmd/service/gateway/models"
	proto_v1_model_user "finalProject/cmd/service/grpc-model/user"
)


type UserRepo interface {
	RegisterUser(context.Context, *proto_v1_model_user.CreateUserReq) (*proto_v1_model_user.CreateUserRes, error)
	GetUser(context.Context, *proto_v1_model_user.GetUserReq) (*proto_v1_model_user.GetUserRes, error)
	Login(context.Context, *proto_v1_model_user.LoginReq) (*models.AuthenticationRes, error)
}