package bill

import "errors"

var (
	ErrInvalidName   = errors.New("bill has an invalid name")
	ErrMustHaveItems = errors.New("bill must have at least one item")
)
