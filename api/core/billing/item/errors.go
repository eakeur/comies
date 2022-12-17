package item

import "errors"

var (
	ErrInvalidUnitPrice = errors.New("item has invalid credits/debts")
	ErrInvalidQuantity  = errors.New("item has an invalid quantity")
)
