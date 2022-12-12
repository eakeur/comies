package movement

import "errors"

// Movements errors
var (
	ErrInvalidType   = errors.New("the movement type is invalid")
	ErrInvalidPeriod = errors.New("the date period informed is invalid")
	ErrInvalidPrice  = errors.New("this movement has an invalid price")
)
