package main

import (
	"context"
	"finalProject/cmd/service/model"
	product "finalProject/cmd/service/proto"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"net"
	"os"
)

type productService struct {
	DB *gorm.DB
}
func (s *productService) Add(ctx context.Context, req *product.AddReq) (res *product.AddRes, err error) {
	p := model.Product{}
	p.Set(req)
	s.DB.Create(&p)

	productRet := &product.Product{}
	p.Fill(productRet)
	res = &product.AddRes{}
	res.Product = productRet
	return res, nil
}

func  (s *productService) Update(context.Context, *product.UpdateReq) (*product.UpdateRes, error) {
	return nil, nil
}

func (s *productService) Delete(context.Context, *product.DeleteReq) (*product.DeleteRes, error) {
	return nil, nil
}

func (s *productService) Get(ctx context.Context, req *product.GetReq) (res *product.GetRes, err error) {
	p := model.Product{}
	s.DB.Find(&p, req.Id)

	productRet := &product.Product{}
	p.Fill(productRet)

	res = &product.GetRes{}
	res.Product = productRet
	return res, nil
}

func main() {
	e := godotenv.Load("./cmd/service/.env") //Load .env file
	if e != nil {
		fmt.Print(e)
		panic(e)
	}
	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	port := os.Getenv("port")
	// Connect database
	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)
	db, err := gorm.Open("postgres", dbUri)
	if err !=nil {
		panic(err)
	}
	db.DropTableIfExists(&model.Product{})
	db.AutoMigrate(&model.Product{})
	defer db.Close()


	lis, _ := net.Listen("tcp", port)

	grpcServer := grpc.NewServer()
	service := productService{
		DB: db,
	}
	product.RegisterProductServiceServer(grpcServer, &service)


	fmt.Println("Start service at 5000")
	grpcServer.Serve(lis)
}