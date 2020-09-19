package main

import (
	configUser "finalProject/cmd/service/user/config"
	"finalProject/cmd/service/user/service"

	"finalProject/cmd/service/product/driver"
	"finalProject/cmd/service/product/model"
	userRepo "finalProject/cmd/service/user/repository/repoimpl"

	"fmt"
	"google.golang.org/grpc"
	"net"
)

func main() {
	configUser := configUser.SetupConfig()
	// Connect database
	db := driver.Connect(configUser.DbHost, configUser.Port, configUser.Username, configUser.Password, configUser.DbName).DB
	redisDb := driver.ConnectRedis(configUser.RedisHost, configUser.PortRedis).DB
	db.DropTableIfExists(&model.Product{})
	db.AutoMigrate(&model.Product{})
	redisDb.FlushAll()
	defer db.Close()
	lisUser, _ := net.Listen("tcp", configUser.HostServer)
		// User service
		grpcUserService := grpc.NewServer()
		userRepositoryImpl := userRepo.NewUserRepo(db, redisDb)
		service.RegisterUserServiceServer(grpcUserService, userRepositoryImpl)
		fmt.Println("Start user service at 5001")
		fmt.Println("====================================")
		grpcUserService.Serve(lisUser)
	fmt.Println("done")

}
