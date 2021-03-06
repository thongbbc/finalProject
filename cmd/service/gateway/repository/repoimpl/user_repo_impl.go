package repoimpl

import (
	"context"
	"finalProject/cmd/service/gateway/models"
	"finalProject/cmd/service/gateway/repository"
	"finalProject/cmd/service/gateway/services"
	grpcError "finalProject/cmd/service/grpc-model/error"
	"finalProject/cmd/service/grpc-model/user"
	modelUser "finalProject/cmd/service/grpc-model/user"
	grpc "finalProject/cmd/service/user/service"
	"google.golang.org/grpc/codes"
)

type UserRepoImpl struct {
	GrpcClient grpc.UserServiceClient
	JWtService services.JWTService
}


func NewUserRepo(grpcClient grpc.UserServiceClient, jwtService services.JWTService) repository.UserRepo {
	return &UserRepoImpl{GrpcClient: grpcClient, JWtService: jwtService}
}

func (i *UserRepoImpl) Login(ctx context.Context, req *modelUser.LoginReq) (*models.AuthenticationRes, *grpcError.Error) {
	findReq := modelUser.LoginReq{
		Email: req.Email,
		Password: req.Password,
	}
	findRes, err := i.GrpcClient.Login(context.TODO(), &findReq)
	if err != nil || findRes.Error != nil  {
		messageError := "Login Failed"
		if  findRes.Error != nil && findRes.Error.Message != "" {
			return nil ,  &grpcError.Error{
				Code: findRes.Error.Code,
				Message: findRes.Error.Message,
			}
		}
		return nil, &grpcError.Error{
			Code: int32(codes.Unauthenticated),
			Message: messageError,
		}
	}
	loginResponse := &models.AuthenticationRes{}
	loginResponse.Email = findRes.GetUser().Email
	jwt := i.JWtService.GenerateToken(loginResponse.Email, true)
	loginResponse.JWT = jwt
	return loginResponse, nil
}


func (i *UserRepoImpl) RegisterUser(ctx context.Context, req *user.CreateUserReq) (res *user.CreateUserRes, err error) {
	addReq := modelUser.CreateUserReq{
		Name: req.Name,
		Email: req.Email,
		Password: req.Password,
	}
	addRes, err := i.GrpcClient.RegisterUser(ctx, &addReq)
	if err != nil {
		return nil, err
	}
	return addRes, nil
}

func (i *UserRepoImpl) GetUser(ctx context.Context, req *user.GetUserReq) (res *user.GetUserRes, err error) {
	findReq := modelUser.GetUserReq{
		Id: req.Id,
	}
	findRes, err := i.GrpcClient.GetUser(context.TODO(), &findReq)
	if err != nil {
		return nil, err
	}
	return findRes, nil
}