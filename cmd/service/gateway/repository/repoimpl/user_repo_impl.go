package repoimpl

import (
	"finalProject/cmd/service/gateway/errorResponse"
	"finalProject/cmd/service/grpc-model/user"
	grpc "finalProject/cmd/service/user/service"

	"context"
	"finalProject/cmd/service/gateway/repository"
	modelUser "finalProject/cmd/service/grpc-model/user"
)

type UserRepoImpl struct {
	GrpcClient grpc.UserServiceClient
}

func NewUserRepo(grpcClient grpc.UserServiceClient) repository.UserRepo {
	return &UserRepoImpl{GrpcClient: grpcClient}
}

func (i *UserRepoImpl) CreateUser(ctx context.Context, req *user.CreateUserReq) (res *user.CreateUserRes, err error) {
	addReq := modelUser.CreateUserReq{
		Name: req.Name,
		Email: req.Email,
		Password: req.Password,
	}
	addRes, err := i.GrpcClient.CreateUser(ctx, &addReq)
	if err != nil {
		return nil, errorResponse.AddFail
	}
	return addRes, nil
}

func (i *UserRepoImpl) GetUser(ctx context.Context, req *user.GetUserReq) (res *user.GetUserRes, err error) {
	findReq := modelUser.GetUserReq{
		Id: req.Id,
	}
	findRes, err := i.GrpcClient.GetUser(context.TODO(), &findReq)
	if err != nil {
		return nil, errorResponse.GetFail
	}
	return findRes, nil
}