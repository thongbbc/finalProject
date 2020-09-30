package models

// Model definition same as gorm.Model, but including column and json tags
type Model struct {
	Id       int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

type ErrorResponse struct {
	Code string `json:code`
	Message string `json: message`
}