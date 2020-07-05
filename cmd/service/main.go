package main

import (
	"finalProject/cmd/service/product"
	"finalProject/cmd/service/product/config"
	"finalProject/cmd/service/product/driver"
	"finalProject/cmd/service/product/model"
	"finalProject/cmd/service/product/repository/repoimpl"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

func main() {
	config := config.SetupConfig()
	// Connect database
	db := driver.Connect(config.DbHost, config.Port, config.Username, config.Password, config.DbName).DB
	redisDb := driver.ConnectRedis(config.RedisHost, config.PortRedis).DB
	db.DropTableIfExists(&model.Product{})
	db.AutoMigrate(&model.Product{})
	redisDb.FlushAll()
	defer db.Close()
	lis, _ := net.Listen("tcp", config.HostServer)

	grpcServer := grpc.NewServer()

	productRepositoryImpl := repoimpl.NewProductRepo(db, redisDb)

	product.RegisterProductServiceServer(grpcServer, productRepositoryImpl)
	fmt.Println("====================================")
	fmt.Println("Start product service at 5000")
	fmt.Println("====================================")
	grpcServer.Serve(lis)
}