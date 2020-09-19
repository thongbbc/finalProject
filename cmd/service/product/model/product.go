package model
import (
	"finalProject/cmd/service/grpc-model/product"
)
type Product struct  {
	Id       int32	`gorm:"primary_key"`
	Sku      string
	Name     string
	Price    float32
	Quantity int32
}

func (p *Product) Set(in *product.AddReq) {
	p.Sku = in.Sku
	p.Name = in.Name
	p.Price = in.Price
	p.Quantity = in.Quantity
}


func (p *Product) Fill(in *product.Product) {
	in.Id = p.Id
	in.Sku = p.Sku
	in.Name = p.Name
	in.Price = p.Price
	in.Quantity = p.Quantity
}