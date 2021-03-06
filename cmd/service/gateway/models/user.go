package models


// User Model
type CreateUserReq struct {
	Email    string `protobuf:"bytes,2,opt,name=email" json:"email,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=password" json:"password,omitempty"`
	Name     string `protobuf:"bytes,4,opt,name=name" json:"name,omitempty"`
}

type LoginReq struct {
	Email string `protobuf:"bytes,2,opt,name=email" json:"email,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=password" json:"password,omitempty"`
}

type AuthenticationRes struct {
	Email string `protobuf:"bytes,2,opt,name=email" json:"email,omitempty"`
	JWT string `protobuf:"bytes,3,opt,name=jwt" json:"jwt,omitempty"`
}

type LoginRes struct {
	AuthenticationRes
	error error
}