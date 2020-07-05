package productHandler

import (
	models "finalProject/cmd/service/product/model"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
}

func (h ProductHandler) GetProduct(c *gin.Context) {
	product := models.Product{}
	c.JSON(200, product)
}

func (h ProductHandler) AddProduct(c *gin.Context) {
	p := &models.Product{}
	c.BindJSON(p)
	c.JSON(200, p)
}