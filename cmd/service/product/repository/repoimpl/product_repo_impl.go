package repoimpl

import (
	"context"
	"encoding/json"
	"finalProject/cmd/service/grpc-model/product"
	"finalProject/cmd/service/product/model"
	"finalProject/cmd/service/product/repository"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)



type ProductRepoImpl struct {
	DB *gorm.DB
	RedisDb *redis.Client
}

func NewProductRepo(db *gorm.DB, redisDb *redis.Client) repository.ProductRepo {
	return &ProductRepoImpl{
		DB: db,
		RedisDb: redisDb,
	}
}

func (i *ProductRepoImpl) Add(ctx context.Context, req *product.AddReq) (res *product.AddRes, err error) {
	p := model.Product{}
	p.Set(req)
	errInsert := i.DB.Create(&p).Error
	if errInsert != nil {
		return nil, status.Errorf(codes.AlreadyExists, fmt.Sprintf("Cannot add product SKU=%s Name=%s", req.Sku, req.Name))
	}
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
	val, redisError := i.RedisDb.Get(string(req.Id)).Result()
	productRet := &product.Product{}
	res = &product.GetRes{}
	if redisError != nil {
		p := model.Product{}
		notfound := i.DB.First(&p, req.Id).Error
		if notfound != nil {
			return nil, status.Error(codes.NotFound, fmt.Sprintf("Not found product with id= %d", req.Id))
		}
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
			return nil, err
		}
		fmt.Println("Get from db", res)
		res.Product = productRet
		return res, nil
	}

	fmt.Println("Get from redis", res)
	err = json.Unmarshal([]byte(val), &res.Product)
	if err != nil {
		return nil, err
	}
	return res, nil
}

