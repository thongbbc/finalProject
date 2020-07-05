package repoimpl

import (
	"context"
	"finalProject/cmd/service/gateway/errorResponse"
	"finalProject/cmd/service/gateway/repository"
	"finalProject/cmd/service/grpc-model/product"
	modelProduct "finalProject/cmd/service/grpc-model/product"
	grpc "finalProject/cmd/service/product"
)

type ProductRepoImpl struct {
	GrpcClient grpc.ProductServiceClient
}

func NewProductRepo(grpcClient grpc.ProductServiceClient) repository.ProductRepo {
	return &ProductRepoImpl{GrpcClient: grpcClient}
}

func (i *ProductRepoImpl) Add(ctx context.Context, req *product.AddReq) (res *product.AddRes, err error) {
	addReq := modelProduct.AddReq{
		Name: req.Name,
		Sku: req.Sku,
		Price: req.Price,
		Quantity: req.Quantity,
	}
	addRes, err := i.GrpcClient.Add(ctx, &addReq)
	if err != nil {
		return nil, errorResponse.AddFail
	}
	return addRes, nil
}

func (i *ProductRepoImpl) Get(ctx context.Context, req *product.GetReq) (res *product.GetRes, err error) {
	findReq := modelProduct.GetReq{
		Id: req.Id,
	}
	findRes, err := i.GrpcClient.Get(context.TODO(), &findReq)
	if err != nil {
		return nil, errorResponse.GetFail
	}
	return findRes, nil
}