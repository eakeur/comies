package stock

import "errors"

var (
	ErrStockAlreadyFull = errors.New("the stock is already full")

	ErrMustHaveTargetID = errors.New("the computation filter must specify a target id")

	ErrNotFound = errors.New("the movement searched does not exist or could not be found")
)
