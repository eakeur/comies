package order

import "errors"

var (
	ErrAlreadyOrdered       = errors.New("this order has been already ordered")
	ErrAlreadyPrepared      = errors.New("this order is already prepared")
	ErrInvalidNumberOfItems = errors.New("there should be at least one item for this order")

	ErrInvalidDeliveryType    = errors.New("the delivery type is invalid")
	ErrInvalidPlacementDate   = errors.New("the placement date is invalid")
	ErrInvalidCustomerName    = errors.New("the customer name is invalid")
	ErrInvalidCustomerPhone   = errors.New("the customer phone is invalid")
	ErrInvalidCustomerAddress = errors.New("the customer address is invalod")
)
