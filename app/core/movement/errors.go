package movement

import "errors"

// Movements errors
var (
	ErrInvalidType        = errors.New("the movement type is invalid")
	ErrInvalidPeriod      = errors.New("the date period informed is invalid")
	ErrInvalidProductType = errors.New("an output movement can not be assigned to a input product or to composite products")
	ErrStockAlreadyFull   = errors.New("the stock is already full")
	ErrStockNegative      = errors.New("the stock, after this output, would be negative. that cannot happen")
)
