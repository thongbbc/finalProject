package main

import (
	"encoding/json"
	"finalProject/cmd/service/gateway/router"
	"fmt"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"net/http"
	"os"
)


// @title Thong API
// @version 1.0
// @description This is a server build by golang.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email nguyenanhthong1995@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	r := router.InitGateway()

	r.Use(static.Serve("/docs", static.LocalFile("./docs", false)))

	url := ginSwagger.URL("http://localhost:8080/info") // The url pointing to API definition
	//Config Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	r.GET("/info", func(c *gin.Context) {
		dir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
			fmt.Println(err)
		}
		fmt.Println(dir)
		file, err := os.Open(fmt.Sprintf("%s/cmd/service/gateway/%s", dir, "docs/swagger.json"))
		if err != nil {
			fmt.Println("====================================")
			fmt.Println(err)
		}
		decoder := json.NewDecoder(file)
		var data interface{}
		if err := decoder.Decode(&data); err != nil {
		}
		c.JSON(http.StatusOK, data)

	})
	fmt.Println("START gateway service at 8080")
	fmt.Println("====================================")
	r.Run(":8080")
}
