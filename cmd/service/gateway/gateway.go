package main
import (
	"github.com/gin-gonic/gin"
	"strconv"
	handlers  "finalProject/cmd/service/gateway/handlers"
)


func main() {
	r := gin.Default()
	productHandler := handlers.ProductHandler{}
	// curl -XPOST -H "Content-Type: application/json" --data '{"sku": "P123", "price": 1000}' http://localhost:3000/product
	r.POST("product", productHandler.AddProduct)
	// route: product/1
	r.GET("/product/:id", productHandler.GetProduct)
	r.Run(":"+strconv.Itoa(3000))
}