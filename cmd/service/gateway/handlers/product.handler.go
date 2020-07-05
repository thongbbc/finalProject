package productHandler

import (
	"finalProject/cmd/service/gateway/repository"
	modelProduct "finalProject/cmd/service/grpc-model/product"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ProductHandler struct {
	ProductRepo repository.ProductRepo
}

func (h ProductHandler) AddProduct(c *gin.Context)  {
	p := &modelProduct.AddReq{}
	c.BindJSON(p)
	addRes, err := h.ProductRepo.Add(c, p)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	fmt.Println("Response of add method is:", addRes)
	c.JSON(http.StatusOK, addRes)
}
func (h ProductHandler) GetProduct(c *gin.Context)  {
	id, paramError := strconv.Atoi(c.Param("id"))
	if paramError != nil {
		c.JSON(http.StatusBadRequest, paramError)
		return
	}
	getReq := &modelProduct.GetReq{
		Id: int32(id),
	}

	getRes, err := h.ProductRepo.Get(c, getReq)
	if err != nil {
		c.JSON(http.StatusNotFound, err)
		return
	}
	fmt.Println("Response of Get method is:", getRes)
	c.JSON(http.StatusOK, getRes)
}