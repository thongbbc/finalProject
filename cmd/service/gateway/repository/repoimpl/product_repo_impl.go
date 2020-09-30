package repoimpl

import (
	"finalProject/cmd/service/gateway/errorResponse"
	"finalProject/cmd/service/gateway/repository"
	"finalProject/cmd/service/gateway/services"
	modelProduct "finalProject/cmd/service/grpc-model/product"
	grpc "finalProject/cmd/service/product/service"
	"fmt"
	"github.com/gin-gonic/gin"
)

type ProductRepoImpl struct {
	GrpcClient grpc.ProductServiceClient
	JWtService services.JWTService
}


func NewProductRepo(grpcClient grpc.ProductServiceClient, jwtService services.JWTService) repository.ProductRepo {
	return &ProductRepoImpl{GrpcClient: grpcClient, JWtService: jwtService}
}
func (i *ProductRepoImpl) Add(c *gin.Context, req *modelProduct.AddReq) (res *modelProduct.AddRes, err error) {
	addReq := modelProduct.AddReq{
		Name: req.Name,
		Sku: req.Sku,
		Price: req.Price,
		Quantity: req.Quantity,
	}
	addRes, err := i.GrpcClient.Add(c, &addReq)
	if err != nil {
		return nil, errorResponse.AddFail
	}
	return addRes, nil
}

func (i *ProductRepoImpl) Get(c *gin.Context, req *modelProduct.GetReq) (res *modelProduct.GetRes, err error) {
	findReq := modelProduct.GetReq{
		Id: req.Id,
	}
	findRes, err := i.GrpcClient.Get(c, &findReq)

	if err != nil {
		return nil, err
	}

	fmt.Println("Get data", findRes)
	fmt.Println("Get data", err)
	return findRes, nil
}