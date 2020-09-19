package repository

import (
	"finalProject/cmd/service/product/service"
)

type ProductRepo interface {
	service.ProductServiceServer
}