// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/v1/model/product/product.proto

/*
Package product is a generated protocol buffer package.

It is generated from these files:
	proto/v1/model/product/product.proto

It has these top-level messages:
	AddReq
	AddRes
	UpdateReq
	UpdateRes
	DeleteReq
	DeleteRes
	GetReq
	GetRes
	Product
*/
package product

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

type AddReq struct {
	Sku      string  `protobuf:"bytes,2,opt,name=sku" json:"sku,omitempty"`
	Name     string  `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Price    float32 `protobuf:"fixed32,4,opt,name=price" json:"price,omitempty"`
	Quantity int32   `protobuf:"varint,5,opt,name=quantity" json:"quantity,omitempty"`
}

func (m *AddReq) Reset()                    { *m = AddReq{} }
func (m *AddReq) String() string            { return proto.CompactTextString(m) }
func (*AddReq) ProtoMessage()               {}
func (*AddReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AddReq) GetSku() string {
	if m != nil {
		return m.Sku
	}
	return ""
}

func (m *AddReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AddReq) GetPrice() float32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *AddReq) GetQuantity() int32 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

type AddRes struct {
	Product *Product                      `protobuf:"bytes,1,opt,name=product" json:"product,omitempty"`
	Error   []*proto_v1_model_error.Error `protobuf:"bytes,2,rep,name=error" json:"error,omitempty"`
}

func (m *AddRes) Reset()                    { *m = AddRes{} }
func (m *AddRes) String() string            { return proto.CompactTextString(m) }
func (*AddRes) ProtoMessage()               {}
func (*AddRes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AddRes) GetProduct() *Product {
	if m != nil {
		return m.Product
	}
	return nil
}

func (m *AddRes) GetError() []*proto_v1_model_error.Error {
	if m != nil {
		return m.Error
	}
	return nil
}

type UpdateReq struct {
	Product *Product `protobuf:"bytes,1,opt,name=product" json:"product,omitempty"`
}

func (m *UpdateReq) Reset()                    { *m = UpdateReq{} }
func (m *UpdateReq) String() string            { return proto.CompactTextString(m) }
func (*UpdateReq) ProtoMessage()               {}
func (*UpdateReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *UpdateReq) GetProduct() *Product {
	if m != nil {
		return m.Product
	}
	return nil
}

type UpdateRes struct {
	Product *Product                    `protobuf:"bytes,1,opt,name=product" json:"product,omitempty"`
	Error   *proto_v1_model_error.Error `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
}

func (m *UpdateRes) Reset()                    { *m = UpdateRes{} }
func (m *UpdateRes) String() string            { return proto.CompactTextString(m) }
func (*UpdateRes) ProtoMessage()               {}
func (*UpdateRes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *UpdateRes) GetProduct() *Product {
	if m != nil {
		return m.Product
	}
	return nil
}

func (m *UpdateRes) GetError() *proto_v1_model_error.Error {
	if m != nil {
		return m.Error
	}
	return nil
}

type DeleteReq struct {
	Id int32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *DeleteReq) Reset()                    { *m = DeleteReq{} }
func (m *DeleteReq) String() string            { return proto.CompactTextString(m) }
func (*DeleteReq) ProtoMessage()               {}
func (*DeleteReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *DeleteReq) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type DeleteRes struct {
	Error *proto_v1_model_error.Error `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
}

func (m *DeleteRes) Reset()                    { *m = DeleteRes{} }
func (m *DeleteRes) String() string            { return proto.CompactTextString(m) }
func (*DeleteRes) ProtoMessage()               {}
func (*DeleteRes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *DeleteRes) GetError() *proto_v1_model_error.Error {
	if m != nil {
		return m.Error
	}
	return nil
}

type GetReq struct {
	Id int32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *GetReq) Reset()                    { *m = GetReq{} }
func (m *GetReq) String() string            { return proto.CompactTextString(m) }
func (*GetReq) ProtoMessage()               {}
func (*GetReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *GetReq) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetRes struct {
	Product *Product                    `protobuf:"bytes,1,opt,name=product" json:"product,omitempty"`
	Error   *proto_v1_model_error.Error `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
}

func (m *GetRes) Reset()                    { *m = GetRes{} }
func (m *GetRes) String() string            { return proto.CompactTextString(m) }
func (*GetRes) ProtoMessage()               {}
func (*GetRes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *GetRes) GetProduct() *Product {
	if m != nil {
		return m.Product
	}
	return nil
}

func (m *GetRes) GetError() *proto_v1_model_error.Error {
	if m != nil {
		return m.Error
	}
	return nil
}

// define Structure
type Product struct {
	Id       int32   `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Sku      string  `protobuf:"bytes,2,opt,name=sku" json:"sku,omitempty"`
	Name     string  `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Price    float32 `protobuf:"fixed32,4,opt,name=price" json:"price,omitempty"`
	Quantity int32   `protobuf:"varint,5,opt,name=quantity" json:"quantity,omitempty"`
}

func (m *Product) Reset()                    { *m = Product{} }
func (m *Product) String() string            { return proto.CompactTextString(m) }
func (*Product) ProtoMessage()               {}
func (*Product) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *Product) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Product) GetSku() string {
	if m != nil {
		return m.Sku
	}
	return ""
}

func (m *Product) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Product) GetPrice() float32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *Product) GetQuantity() int32 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

func init() {
	proto.RegisterType((*AddReq)(nil), "proto.v1.model.product.AddReq")
	proto.RegisterType((*AddRes)(nil), "proto.v1.model.product.AddRes")
	proto.RegisterType((*UpdateReq)(nil), "proto.v1.model.product.UpdateReq")
	proto.RegisterType((*UpdateRes)(nil), "proto.v1.model.product.UpdateRes")
	proto.RegisterType((*DeleteReq)(nil), "proto.v1.model.product.DeleteReq")
	proto.RegisterType((*DeleteRes)(nil), "proto.v1.model.product.DeleteRes")
	proto.RegisterType((*GetReq)(nil), "proto.v1.model.product.GetReq")
	proto.RegisterType((*GetRes)(nil), "proto.v1.model.product.GetRes")
	proto.RegisterType((*Product)(nil), "proto.v1.model.product.Product")
}

func init() { proto.RegisterFile("proto/v1/model/product/product.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 323 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x53, 0xcf, 0x4b, 0xc3, 0x30,
	0x14, 0x26, 0xdd, 0xba, 0xb9, 0x37, 0x10, 0x09, 0x22, 0x61, 0x3b, 0x58, 0x8a, 0x87, 0x82, 0xac,
	0x65, 0xf3, 0xe4, 0x45, 0x50, 0xfc, 0x71, 0x1d, 0x01, 0x2f, 0x9e, 0xac, 0xc9, 0x53, 0xa2, 0xdb,
	0xd2, 0xa5, 0x59, 0x61, 0xff, 0xbd, 0x2c, 0xe9, 0x1c, 0x8e, 0x1d, 0x86, 0xa8, 0x97, 0xf7, 0xa3,
	0xdf, 0x7b, 0xdf, 0xfb, 0x3e, 0x68, 0xe0, 0xac, 0x30, 0xda, 0xea, 0xac, 0x1a, 0x66, 0x53, 0x2d,
	0x71, 0x92, 0x15, 0x46, 0xcb, 0x85, 0xb0, 0xeb, 0x9c, 0x3a, 0x98, 0x9e, 0xb8, 0x94, 0x56, 0xc3,
	0xd4, 0x4d, 0xa5, 0x35, 0xda, 0x8b, 0xb6, 0xb6, 0xd1, 0x18, 0x6d, 0x7c, 0xf4, 0x9b, 0xf1, 0x33,
	0xb4, 0xae, 0xa5, 0xe4, 0x38, 0xa7, 0x47, 0xd0, 0x28, 0x3f, 0x16, 0x2c, 0x88, 0x48, 0xd2, 0xe1,
	0xab, 0x92, 0x52, 0x68, 0xce, 0xf2, 0x29, 0xb2, 0x86, 0xfb, 0xe4, 0x6a, 0x7a, 0x0c, 0x61, 0x61,
	0x94, 0x40, 0xd6, 0x8c, 0x48, 0x12, 0x70, 0xdf, 0xd0, 0x1e, 0x1c, 0xcc, 0x17, 0xf9, 0xcc, 0x2a,
	0xbb, 0x64, 0x61, 0x44, 0x92, 0x90, 0x7f, 0xf5, 0x71, 0x55, 0x5f, 0x28, 0xe9, 0x25, 0xb4, 0x6b,
	0x61, 0x8c, 0x44, 0x24, 0xe9, 0x8e, 0x4e, 0xd3, 0xdd, 0xba, 0xd3, 0xb1, 0xcf, 0x7c, 0x3d, 0x4f,
	0x87, 0x10, 0x3a, 0xd5, 0x2c, 0x88, 0x1a, 0x49, 0x77, 0xd4, 0xdf, 0x5e, 0xf4, 0x96, 0xee, 0x56,
	0x91, 0xfb, 0xc9, 0xf8, 0x1e, 0x3a, 0x8f, 0x85, 0xcc, 0x2d, 0xae, 0xcc, 0xfd, 0xfc, 0x74, 0xbc,
	0xdc, 0xf0, 0xfc, 0x96, 0x05, 0xb2, 0xa7, 0x85, 0x3e, 0x74, 0x6e, 0x71, 0x82, 0xde, 0xc2, 0x21,
	0x04, 0x4a, 0xba, 0xab, 0x21, 0x0f, 0x94, 0x8c, 0xaf, 0x36, 0x60, 0xb9, 0x21, 0x27, 0x7b, 0x93,
	0x33, 0x68, 0x3d, 0xa0, 0xdd, 0xc5, 0x5c, 0xd5, 0xc8, 0x7f, 0xdb, 0x9d, 0x43, 0xbb, 0xa6, 0xd9,
	0x96, 0xf4, 0x57, 0x3f, 0xe7, 0xcd, 0xe0, 0xe9, 0xfc, 0x55, 0xcd, 0xf2, 0xc9, 0xd8, 0xe8, 0x77,
	0x14, 0x36, 0x13, 0x53, 0x99, 0x95, 0x68, 0x2a, 0x25, 0x30, 0x7b, 0x33, 0x85, 0x18, 0x7c, 0x7b,
	0x75, 0x2f, 0x2d, 0x67, 0xe2, 0xe2, 0x33, 0x00, 0x00, 0xff, 0xff, 0x97, 0x2e, 0xe2, 0x53, 0x96,
	0x03, 0x00, 0x00,
}
