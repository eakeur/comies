package product

import "errors"

var (
	ErrNotFound = errors.New("this product seems to not exist")

	ErrCodeAlreadyExists = errors.New("the product code is assigned to another product")

	ErrInvalidCode = errors.New("the code is not valid. please provide one longer than 2 (two) and shorter than 12 (twelve) characters")

	ErrInvalidName = errors.New("the name is not valid. please provide one longer than 2 (two) and shorter than 60 (sixty) characters")

	ErrInvalidType = errors.New("the product type is invalid")

	ErrInvalidPrice = errors.New("the price is not valid. please provide an amount higher than 0")

	ErrMinimumSaleQuantity = errors.New("the minimum sale quantity provided is not valid. please provide a value higher than 0")

	ErrInvalidSaleQuantity = errors.New("the sale quantity requested for this product is invalid. please check the correct one")

	ErrStockAlreadyFull = errors.New("the stock is already full")

	ErrStockNegative = errors.New("the stock, after this output, would be negative. that cannot happen")
)
