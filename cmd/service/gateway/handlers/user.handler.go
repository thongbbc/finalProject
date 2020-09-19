

package handlers

import (
	"finalProject/cmd/service/gateway/repository"
	modelUser "finalProject/cmd/service/grpc-model/user"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserHandler struct {
	UserRepo repository.UserRepo
}
// curl -XPOST -H "Content-Type: application/json" --data '{"name": "test", "email": "test@gmail.com"}' http://localhost:3000/v1/auth/register
func (h UserHandler) RegisterUser(c *gin.Context)  {
	p := &modelUser.CreateUserReq{}
	c.BindJSON(p)
	addRes, err := h.UserRepo.RegisterUser(c, p)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	fmt.Println("Response of add method is:", addRes)
	c.JSON(http.StatusOK, addRes)
}
// route: user/1
func (h UserHandler) GetUser(c *gin.Context)  {
	id, paramError := strconv.Atoi(c.Param("id"))
	if paramError != nil {
		c.JSON(http.StatusBadRequest, paramError)
		return
	}
	getReq := &modelUser.GetUserReq{
		Id: int32(id),
	}

	getRes, err := h.UserRepo.GetUser(c, getReq)
	if err != nil {
		c.JSON(http.StatusNotFound, err)
		return
	}
	fmt.Println("Response of Get method is:", getRes)
	c.JSON(http.StatusOK, getRes)
}