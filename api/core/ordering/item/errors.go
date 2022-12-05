package item

import "errors"

var (
	ErrInvalidQuantity = errors.New("this item has invalid quantity")
	ErrInvalidStatus   = errors.New("the item status is invalid")
)
