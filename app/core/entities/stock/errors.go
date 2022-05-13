package stock

import "errors"

var (
	ErrStockFull = errors.New("the stock has already achieved the maximum value")

	ErrStockEmpty = errors.New("the stock has already achieved the minimum value")
)
