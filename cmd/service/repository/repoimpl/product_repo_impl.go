package repoimpl

import (
	"context"
	"encoding/json"
	"finalProject/cmd/service/model"
	product "finalProject/cmd/service/proto"
	"finalProject/cmd/service/repository"
	"fmt"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

type ProductRepoImpl struct {
	DB      *gorm.DB
	RedisDb *redis.Client
}

func NewProductRepo(db *gorm.DB, redisDb *redis.Client) repository.ProductRepo {
	return &ProductRepoImpl{
		DB:      db,
		RedisDb: redisDb,
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
	return res, nil
}

func (i *ProductRepoImpl) Update(ctx context.Context, req *product.UpdateReq) (res *product.UpdateRes, err error) {
	p := model.Product{}
	if err = i.DB.First(&p, req.Product.GetId()).Error; err != nil {
		return nil, err
	}
	p.Name = req.Product.GetName()
	p.Price = req.Product.GetPrice()
	p.Quantity = p.Quantity * 5
	p.Sku = req.Product.GetSku()
	i.DB.Save(p)
	productRet := &product.Product{}
	p.Fill(productRet)
	res = &product.UpdateRes{}
	res.Product = productRet
	return res, nil
}

func (i *ProductRepoImpl) Delete(ctx context.Context, req *product.DeleteReq) (res *product.DeleteRes, err error) {
	id := req.GetId()
	p := model.Product{}
	if err = i.DB.First(&p, req.GetId()).Error; err != nil {
		fmt.Println("errr Delete")
		return nil, err
	}
	i.DB.Where("id = ?", id).Delete(p)
	res = nil
	return res, nil
}

func (i *ProductRepoImpl) Get(ctx context.Context, req *product.GetReq) (res *product.GetRes, err error) {
	val, err := i.RedisDb.Get(string(req.Id)).Result()
	productRet := &product.Product{}
	res = &product.GetRes{}
	if err != nil {
		p := model.Product{}
		i.DB.First(&p, req.Id)

		p.Fill(productRet)
		json, err := json.Marshal(product.Product{
			Id:       productRet.Id,
			Sku:      productRet.Sku,
			Name:     productRet.Name,
			Price:    productRet.Price,
			Quantity: productRet.Quantity,
		})
		err = i.RedisDb.Set(string(req.Id), json, 0).Err()
		if err != nil {
			panic(err)
		}
		res.Product = productRet
		return res, nil
	}
	err = json.Unmarshal([]byte(val), &res.Product)
	return res, nil
}
