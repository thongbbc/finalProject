package repository

import (
	product "finalProject/cmd/service/proto"
)

type ProductRepo interface {
	product.ProductServiceServer
}