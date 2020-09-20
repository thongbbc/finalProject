package repository

import (
	proto_v1_model_product "finalProject/cmd/service/grpc-model/product"
	"github.com/gin-gonic/gin"
)

type ProductRepo interface {
	Add(*gin.Context, *proto_v1_model_product.AddReq) (*proto_v1_model_product.AddRes, error)
	Get(*gin.Context, *proto_v1_model_product.GetReq) (*proto_v1_model_product.GetRes, error)
}