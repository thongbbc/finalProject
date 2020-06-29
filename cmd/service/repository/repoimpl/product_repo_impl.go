package repoimpl

import (
	"context"
	"finalProject/cmd/service/model"
	product "finalProject/cmd/service/proto"
	"finalProject/cmd/service/repository"
	"github.com/jinzhu/gorm"
)

type ProductRepoImpl struct {
	DB *gorm.DB

}

func NewProductRepo(db *gorm.DB) repository.ProductRepo {
	return &ProductRepoImpl{
		DB: db,
	}
}

func (i *ProductRepoImpl) Add(ctx context.Context, req *product.AddReq) (res *product.AddRes, err error) {
	p := model.Product{}
	p.Set(req)
	i.DB.Create(&p)
	productRet := &product.Product{}
	p.Fill(productRet)
	res = &product.AddRes{}
	res.Product = productRet
	return res, nil}

func (i *ProductRepoImpl) Update(context.Context, *product.UpdateReq) (*product.UpdateRes, error) {
	return nil, nil
}
func (i *ProductRepoImpl) Delete(context.Context, *product.DeleteReq) (*product.DeleteRes, error) {
	return nil, nil
}
func (i *ProductRepoImpl) Get(ctx context.Context, req *product.GetReq) (res *product.GetRes, err error) {
	p := model.Product{}
	i.DB.First(&p, req.Id)

	productRet := &product.Product{}
	p.Fill(productRet)

	res = &product.GetRes{}
	res.Product = productRet
	return res, nil
}
