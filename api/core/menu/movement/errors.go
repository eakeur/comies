package movement

import "errors"

// Movements errors
var (
	ErrInvalidType        = errors.New("the movement type is invalid")
	ErrInvalidPeriod      = errors.New("the date period informed is invalid")
	ErrInvalidPrice       = errors.New("this movement has an invalid price")
	ErrInvalidProductType = errors.New("an output movement can not be assigned to an input or composite product")
)
