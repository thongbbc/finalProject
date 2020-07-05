package errorResponse

import (
	"errors"
)

var (
	AddFail = errors.New("Khong the tao user!")
	GetFail   = errors.New("Khong tim thay!")
)
