package order

import "errors"

var (
	ErrMissingProducts = errors.New("there must be at least one product related to this item")
	ErrInvalidQuantity = errors.New("an item must have a quantity bigger than 0")
)
