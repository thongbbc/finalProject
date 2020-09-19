package main

import (
	configProduct "finalProject/cmd/service/product/config"
	"finalProject/cmd/service/product/service"

	"finalProject/cmd/service/product/driver"
	"finalProject/cmd/service/product/model"
	productRepo "finalProject/cmd/service/product/repository/repoimpl"

	"fmt"
	"google.golang.org/grpc"
	"net"
)

func main() {
	configProduct := configProduct.SetupConfig()
	// Connect database
	db := driver.Connect(configProduct.DbHost, configProduct.Port, configProduct.Username, configProduct.Password, configProduct.DbName).DB
	redisDb := driver.ConnectRedis(configProduct.RedisHost, configProduct.PortRedis).DB
	db.DropTableIfExists(&model.Product{})
	db.AutoMigrate(&model.Product{})
	redisDb.FlushAll()
	defer db.Close()
	lisProduct, _ := net.Listen("tcp", configProduct.HostServer)

	// Product service
	grpcProductService := grpc.NewServer()
	productRepositoryImpl := productRepo.NewProductRepo(db, redisDb)
	service.RegisterProductServiceServer(grpcProductService, productRepositoryImpl)
	fmt.Println("Start product service at 5000")
	fmt.Println("====================================")
	grpcProductService.Serve(lisProduct)
	fmt.Println("done")

}
