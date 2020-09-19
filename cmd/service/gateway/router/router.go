package router

import (
	"finalProject/cmd/service/gateway/driver"

	handlers "finalProject/cmd/service/gateway/handlers"
	"finalProject/cmd/service/gateway/repository/repoimpl"
	productService "finalProject/cmd/service/product/service"
	userService "finalProject/cmd/service/user/service"

	"github.com/gin-gonic/gin"
)

type API struct {
	Gin        *gin.Engine
	ProductHandler handlers.ProductHandler
	UserHandler handlers.UserHandler
}

func (api *API) SetupRouter() *gin.Engine {
	v1 := api.Gin.Group("/api/v1")
	//auth
	auth := v1.Group("/auth")
	auth.POST("/register", api.UserHandler.RegisterUser)
	auth.POST("/login", api.UserHandler.GetUser)

	api.Gin.GET("/user/:id", api.UserHandler.GetUser)
	api.Gin.POST("/product", api.ProductHandler.AddProduct)
	api.Gin.GET("/product/:id", api.ProductHandler.GetProduct)
	return api.Gin
}


func InitGateway(host ...string) *gin.Engine {
	var productClient productService.ProductServiceClient
	var userClient userService.UserServiceClient

	if host == nil {
		userClient = driver.ConnectUserService("user-service", "5001")
		productClient = driver.ConnectProductService("product-service", "5000")
	} else {
		userClient = driver.ConnectUserService(host[0], "5001")
		productClient = driver.ConnectProductService(host[1], "5000")
	}
	userHandler := handlers.UserHandler{
		UserRepo: repoimpl.NewUserRepo(userClient),
	}
	productHandler := handlers.ProductHandler{
		ProductRepo: repoimpl.NewProductRepo(productClient),
	}
	// curl -XPOST -H "Content-Type: application/json" --data '{"sku": "P123", "price": 1000}' http://localhost:3000/product
	// route: product/1
	api := API{
		Gin:         gin.Default(),
		ProductHandler: productHandler,
		UserHandler: userHandler,
	}

	r := api.SetupRouter()
	return r
}