package stock

import "errors"

var (
	ErrMissingResourceID = errors.New("the resource id informed is empty or invalid")

	ErrInvalidPeriod = errors.New("the date period informed is invalid")

	ErrStockFull = errors.New("the stock has already achieved the maximum value")

	ErrStockEmpty = errors.New("the stock has already achieved the minimum value")
)
