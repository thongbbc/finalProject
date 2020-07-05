package main
import (
	"finalProject/cmd/service/gateway/driver"
	handlers "finalProject/cmd/service/gateway/handlers"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)


func main() {
	r := gin.Default()
	productHandler := handlers.ProductHandler{}
	// curl -XPOST -H "Content-Type: application/json" --data '{"sku": "P123", "price": 1000}' http://localhost:3000/product
	r.POST("product", productHandler.AddProduct)
	// route: product/1
	r.GET("/product/:id", productHandler.GetProduct)
	fmt.Println("====================================")
	fmt.Println("START gateway service at 3000")
	fmt.Println("====================================")
	driver.ConnectProduct("product-service", "5000")
	r.Run(":"+strconv.Itoa(3000))
}