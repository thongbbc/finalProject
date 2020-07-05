package repository

import (
	"finalProject/cmd/service/product"
)

type ProductRepo interface {
	product.ProductServiceServer
}