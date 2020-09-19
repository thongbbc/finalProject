package driver

import (
	grpcProductService "finalProject/cmd/service/product/service"
	grpcUserService "finalProject/cmd/service/user/service"
	"fmt"
	"google.golang.org/grpc"
)


func ConnectProductService(host string, port string) grpcProductService.ProductServiceClient {
	grpcUrl := fmt.Sprintf("%s:%s", host, port)
	conn, err := grpc.Dial(grpcUrl, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := grpcProductService.NewProductServiceClient(conn)
	return client
}


func ConnectUserService(host string, port string) grpcUserService.UserServiceClient {
	grpcUrl := fmt.Sprintf("%s:%s", host, port)
	conn, err := grpc.Dial(grpcUrl, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := grpcUserService.NewUserServiceClient(conn)
	return client
}

