package item

import "errors"

var (
	ErrInvalidQuantity = errors.New("this content has invalid quantity")
	ErrInvalidStatus   = errors.New("the status is invalid")
)
