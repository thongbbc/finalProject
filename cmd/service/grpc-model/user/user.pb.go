// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/v1/model/user/user.proto

/*
Package user is a generated protocol buffer package.

It is generated from these files:
	proto/v1/model/user/user.proto

It has these top-level messages:
	DeleteUserReq
	DeleteUserRes
	CreateUserReq
	CreateUserRes
	GetUserReq
	GetUserRes
	LoginReq
	User
*/
package user

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import proto_v1_model_error "finalProject/cmd/service/grpc-model/error"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type DeleteUserReq struct {
	Id int32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *DeleteUserReq) Reset()                    { *m = DeleteUserReq{} }
func (m *DeleteUserReq) String() string            { return proto.CompactTextString(m) }
func (*DeleteUserReq) ProtoMessage()               {}
func (*DeleteUserReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *DeleteUserReq) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type DeleteUserRes struct {
	Id int32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *DeleteUserRes) Reset()                    { *m = DeleteUserRes{} }
func (m *DeleteUserRes) String() string            { return proto.CompactTextString(m) }
func (*DeleteUserRes) ProtoMessage()               {}
func (*DeleteUserRes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *DeleteUserRes) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type CreateUserReq struct {
	Email    string `protobuf:"bytes,2,opt,name=email" json:"email,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=password" json:"password,omitempty"`
	Name     string `protobuf:"bytes,4,opt,name=name" json:"name,omitempty"`
}

func (m *CreateUserReq) Reset()                    { *m = CreateUserReq{} }
func (m *CreateUserReq) String() string            { return proto.CompactTextString(m) }
func (*CreateUserReq) ProtoMessage()               {}
func (*CreateUserReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CreateUserReq) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *CreateUserReq) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *CreateUserReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type CreateUserRes struct {
	User  *User                       `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
	Error *proto_v1_model_error.Error `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
}

func (m *CreateUserRes) Reset()                    { *m = CreateUserRes{} }
func (m *CreateUserRes) String() string            { return proto.CompactTextString(m) }
func (*CreateUserRes) ProtoMessage()               {}
func (*CreateUserRes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *CreateUserRes) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *CreateUserRes) GetError() *proto_v1_model_error.Error {
	if m != nil {
		return m.Error
	}
	return nil
}

type GetUserReq struct {
	Id int32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *GetUserReq) Reset()                    { *m = GetUserReq{} }
func (m *GetUserReq) String() string            { return proto.CompactTextString(m) }
func (*GetUserReq) ProtoMessage()               {}
func (*GetUserReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *GetUserReq) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetUserRes struct {
	User  *User                       `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
	Error *proto_v1_model_error.Error `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
}

func (m *GetUserRes) Reset()                    { *m = GetUserRes{} }
func (m *GetUserRes) String() string            { return proto.CompactTextString(m) }
func (*GetUserRes) ProtoMessage()               {}
func (*GetUserRes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *GetUserRes) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *GetUserRes) GetError() *proto_v1_model_error.Error {
	if m != nil {
		return m.Error
	}
	return nil
}

type LoginReq struct {
	Email    string `protobuf:"bytes,1,opt,name=email" json:"email,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
}

func (m *LoginReq) Reset()                    { *m = LoginReq{} }
func (m *LoginReq) String() string            { return proto.CompactTextString(m) }
func (*LoginReq) ProtoMessage()               {}
func (*LoginReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *LoginReq) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *LoginReq) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type User struct {
	Id       int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Email    string `protobuf:"bytes,2,opt,name=email" json:"email,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=password" json:"password,omitempty"`
	Name     string `protobuf:"bytes,4,opt,name=name" json:"name,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *User) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*DeleteUserReq)(nil), "proto.v1.model.user.DeleteUserReq")
	proto.RegisterType((*DeleteUserRes)(nil), "proto.v1.model.user.DeleteUserRes")
	proto.RegisterType((*CreateUserReq)(nil), "proto.v1.model.user.CreateUserReq")
	proto.RegisterType((*CreateUserRes)(nil), "proto.v1.model.user.CreateUserRes")
	proto.RegisterType((*GetUserReq)(nil), "proto.v1.model.user.GetUserReq")
	proto.RegisterType((*GetUserRes)(nil), "proto.v1.model.user.GetUserRes")
	proto.RegisterType((*LoginReq)(nil), "proto.v1.model.user.LoginReq")
	proto.RegisterType((*User)(nil), "proto.v1.model.user.User")
}

func init() { proto.RegisterFile("proto/v1/model/user/user.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 295 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x52, 0xcd, 0x4f, 0xbb, 0x40,
	0x10, 0x0d, 0xfc, 0xe8, 0x2f, 0x75, 0x9a, 0x7a, 0x58, 0x3d, 0x60, 0x35, 0xda, 0x70, 0x22, 0x26,
	0xdd, 0x0d, 0xf5, 0xea, 0xc9, 0x8f, 0x78, 0xf1, 0x60, 0x48, 0x7a, 0x31, 0x5e, 0x10, 0xc6, 0x66,
	0x0d, 0xb0, 0x74, 0x16, 0xf1, 0xdf, 0x37, 0x0c, 0x89, 0x15, 0xb4, 0x37, 0xe3, 0x65, 0x60, 0xf6,
	0xbd, 0xf9, 0x78, 0x6f, 0x17, 0x4e, 0x2b, 0x32, 0xb5, 0x51, 0x4d, 0xa4, 0x0a, 0x93, 0x61, 0xae,
	0xde, 0x2c, 0x12, 0x07, 0xc9, 0x80, 0x38, 0xe0, 0x8f, 0x6c, 0x22, 0xc9, 0xb8, 0x6c, 0xa1, 0xd9,
	0x7c, 0x50, 0x84, 0x44, 0x86, 0xba, 0xd8, 0x95, 0x05, 0x67, 0x30, 0xbd, 0xc1, 0x1c, 0x6b, 0x5c,
	0x59, 0xa4, 0x18, 0x37, 0x62, 0x1f, 0x5c, 0x9d, 0xf9, 0xce, 0xdc, 0x09, 0x47, 0xb1, 0xab, 0xb3,
	0x21, 0xc1, 0x7e, 0x23, 0xac, 0x60, 0x7a, 0x4d, 0x98, 0x6c, 0x3b, 0x1c, 0xc2, 0x08, 0x8b, 0x44,
	0xe7, 0xbe, 0x3b, 0x77, 0xc2, 0xbd, 0xb8, 0x4b, 0xc4, 0x0c, 0xc6, 0x55, 0x62, 0xed, 0xbb, 0xa1,
	0xcc, 0xff, 0xc7, 0xc0, 0x67, 0x2e, 0x04, 0x78, 0x65, 0x52, 0xa0, 0xef, 0xf1, 0x39, 0xff, 0x07,
	0x9b, 0x7e, 0x5b, 0x2b, 0x16, 0xe0, 0xb5, 0x9a, 0x78, 0xf2, 0x64, 0x79, 0x24, 0x7f, 0xd0, 0x2b,
	0x99, 0xcb, 0x34, 0x11, 0xc1, 0x88, 0x75, 0xf2, 0x16, 0x93, 0xe5, 0xf1, 0x90, 0xdf, 0x99, 0x70,
	0xdb, 0xc6, 0xb8, 0x63, 0x06, 0x27, 0x00, 0x77, 0x58, 0xef, 0x32, 0xa2, 0xfc, 0x82, 0xfe, 0xc5,
	0x36, 0x97, 0x30, 0xbe, 0x37, 0x6b, 0x5d, 0xf6, 0x2c, 0x75, 0x76, 0x59, 0xea, 0xf6, 0x2d, 0x0d,
	0x9e, 0xc0, 0x6b, 0xc7, 0x0f, 0x55, 0xfc, 0xce, 0xe5, 0x5c, 0x9d, 0x3f, 0x86, 0x2f, 0xba, 0x4c,
	0xf2, 0x07, 0x32, 0xaf, 0x98, 0xd6, 0x2a, 0x2d, 0x32, 0x65, 0x91, 0x1a, 0x9d, 0xa2, 0x5a, 0x53,
	0x95, 0x2e, 0xb6, 0x6f, 0xf4, 0xf9, 0x3f, 0x4b, 0xbd, 0xf8, 0x08, 0x00, 0x00, 0xff, 0xff, 0x14,
	0xea, 0x65, 0x2e, 0xc1, 0x02, 0x00, 0x00,
}
