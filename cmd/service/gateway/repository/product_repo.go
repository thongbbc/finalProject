package repository

import (
	"context"
	proto_v1_model_product "finalProject/cmd/service/grpc-model/product"
)

type ProductRepo interface {
	Add(context.Context, *proto_v1_model_product.AddReq) (*proto_v1_model_product.AddRes, error)
	Get(context.Context, *proto_v1_model_product.GetReq) (*proto_v1_model_product.GetRes, error)
}