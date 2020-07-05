package router

import (
	"finalProject/cmd/service/gateway/driver"
	"finalProject/cmd/service/gateway/handlers"
	handlers "finalProject/cmd/service/gateway/handlers"
	"finalProject/cmd/service/gateway/repository/repoimpl"
	"github.com/gin-gonic/gin"
)

type API struct {
	Gin        *gin.Engine
	ProductHandler productHandler.ProductHandler
}

func (api *API) SetupRouter() *gin.Engine {
	api.Gin.POST("/product", api.ProductHandler.AddProduct)
	api.Gin.GET("/product/:id", api.ProductHandler.GetProduct)
	return api.Gin
}


func InitGateway() *gin.Engine {
	grpcClient := driver.ConnectProduct("product-service", "5000")

	productHandler := handlers.ProductHandler{
		ProductRepo: repoimpl.NewProductRepo(grpcClient),
	}
	// curl -XPOST -H "Content-Type: application/json" --data '{"sku": "P123", "price": 1000}' http://localhost:3000/product
	// route: product/1
	api := API{
		Gin:         gin.Default(),
		ProductHandler: productHandler,
	}

	r := api.SetupRouter()
	return r
}