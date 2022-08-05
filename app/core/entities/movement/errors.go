package movement

import "errors"

var (
	ErrInvalidType        = errors.New("the movement type is invalid")
	ErrInvalidPeriod      = errors.New("the date period informed is invalid")
	ErrInvalidProductType = errors.New("an output movement can not be assigned to a input product or to composite products")
)
