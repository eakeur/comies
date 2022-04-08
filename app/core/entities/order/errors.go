package order

import "errors"

var (
	ErrProductsAndQuantities = errors.New("the array of product IDs should be the same size as the array of quantity IDs in the Item object")
	ErrMissingProductIDs     = errors.New("there must be at least one product id related to this item")
	ErrMissingQuantities     = errors.New("there must be at least one quantity related to this item")
)
