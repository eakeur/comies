package order

import "errors"

var (
	ErrAlreadyOrdered       = errors.New("this order has been already ordered")
	ErrInvalidNumberOfItems = errors.New("there should be at least one item for this order")
)
