package repository

import (
	"finalProject/cmd/service/user/service"
)

type UserRepo interface {
	service.UserServiceServer
}