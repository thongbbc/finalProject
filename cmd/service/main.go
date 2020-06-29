package main

import (
	"finalProject/cmd/service/config"
	"finalProject/cmd/service/driver"
	"finalProject/cmd/service/model"
	product "finalProject/cmd/service/proto"
	"finalProject/cmd/service/repository/repoimpl"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"google.golang.org/grpc"
	"net"
)

func main() {
	config := config.SetupConfig()
	// Connect database
	db := driver.Connect(config.DbHost, config.Port, config.Username, config.Password, config.DbName).DB
	db.DropTableIfExists(&model.Product{})
	db.AutoMigrate(&model.Product{})
	defer db.Close()
	lis, _ := net.Listen("tcp", config.HostServer)

	grpcServer := grpc.NewServer()

	productRepositoryImpl := repoimpl.NewProductRepo(db)

	product.RegisterProductServiceServer(grpcServer, productRepositoryImpl)
	fmt.Println("Start service at 5000")
	grpcServer.Serve(lis)
}