package order

import "errors"

var (
	ErrAlreadyOrdered       = errors.New("this order has been already ordered")
	ErrAlreadyPreparing     = errors.New("this order is already being prepared")
	ErrInvalidNumberOfItems = errors.New("there should be at least one item for this order")
)
