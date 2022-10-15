package ordering

import "errors"

// Item errors
var (
	ErrInvalidQuantity = errors.New("this content has invalid quantity")
)

// Order errors
var (
	ErrAlreadyOrdered       = errors.New("this order has been already ordered")
	ErrAlreadyPreparing     = errors.New("this order is already being prepared")
	ErrInvalidNumberOfItems = errors.New("there should be at least one item for this order")

	ErrInvalidDeliveryType = errors.New("the delivery type is invalid")
	ErrInvalidPlacementDate = errors.New("the placement date is invalid")
	ErrInvalidCustomerName = errors.New("the customer name is invalid")
	ErrInvalidCustomerPhone = errors.New("the customer phone is invalid")
	ErrInvalidCustomerAddress = errors.New("the customer address is invalod")

)

// General errors
var (
	ErrInvalidStatus        = errors.New("the status is invalid")
)