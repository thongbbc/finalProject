

package handlers

import (
	"finalProject/cmd/service/gateway/repository"
	modelUser "finalProject/cmd/service/grpc-model/user"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"net/http"
	"strconv"
)

type UserHandler struct {
	UserRepo repository.UserRepo
}


// @Summary Register user
// @Description register user
// @Accept  json
// @Produce  json
// @Param Register body models.CreateUserReq true "Register User"
// @Success 201 {object} models.CreateUserReq
// @Failure 400 {object} models.ErrorResponse
// @Router /auth/register [post]
// curl -XPOST -H "Content-Type: application/json" --data '{"name": "test", "email": "test@gmail.com"}' http://localhost:3000/v1/auth/register
func (h UserHandler) RegisterUser(c *gin.Context)  {
	p := &modelUser.CreateUserReq{}
	c.BindJSON(p)
	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(p.Password), bcrypt.DefaultCost)
	p.Password = string(hashedPassword)
	addRes, err := h.UserRepo.RegisterUser(c, p)
	if err != nil || addRes.Error != nil {
		messageError := "Register failed!"
		if addRes.Error !=nil && addRes.Error.Message != "" {
			messageError = addRes.Error.Message
		}
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": messageError})
		return
	}
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

// @Summary Login user
// @Description Login user
// @Accept  json
// @Produce  json
// @Param Login body models.LoginReq true "Login Account"
// @Success 201 {object} models.AuthenticationRes
// @Failure 400 {object} models.ErrorResponse
// @Router /auth/login [post]
// curl -XPOST -H "Content-Type: application/json" --data '{"password": "test", "email": "test@gmail.com"}' http://localhost:3000/v1/auth/login
func (h UserHandler) Login(c *gin.Context)  {
	p := &modelUser.LoginReq{}
	errBindJson := c.BindJSON(p)
	if errBindJson != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "Bad request"})
		return
	}

	getReq := &modelUser.LoginReq{
		Email: p.Email,
		Password: p.Password,
	}

	getRes, err := h.UserRepo.Login(c, getReq)
	if err != nil && err.Message != "" && err.Code == int32(codes.Unauthenticated) {
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": err.Message})
		return
	} else if err != nil && err.Message != "" && err.Code == int32(codes.NotFound)  {
		c.JSON(http.StatusNotFound, gin.H{"code": 404, "message": err.Message})
		return
	} else if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 404, "message": "Login Failed"})
		return
	}
	c.JSON(http.StatusOK, getRes)
}