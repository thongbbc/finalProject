package model
import (
	"finalProject/cmd/service/grpc-model/user"
)
type User struct  {
	Id       int32	`gorm:"primary_key"`
	Name     string
	Email    string
	Password string
}

func (p *User) Set(in *user.CreateUserReq) {
	p.Name = in.Name
	p.Email = in.Email
	p.Password = in.Password
}


func (p *User) Fill(in *user.User) {
	in.Id = p.Id
	in.Name = p.Name
	in.Email = p.Email
	in.Password = p.Password
}