package price

import "errors"

var (
	ErrInvalidDate  = errors.New("the date field of this type is empty")
	ErrInvalidValue = errors.New("the value of price should be bigger than or equal to 0")
)
