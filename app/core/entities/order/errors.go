package order

import "errors"

var (
	ErrMissingProducts      = errors.New("there must be at least one product related to this item")
	ErrInvalidQuantity      = errors.New("an item must have a quantity bigger than 0")
	ErrOrderAlreadyOrdered  = errors.New("this order has been already ordered")
	ErrInvalidNumberOfItems = errors.New("there should be at least one item for this order")
)
