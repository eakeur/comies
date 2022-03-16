package product

import "errors"

var (
	ErrInvalidCode = errors.New("the code is not valid. please provide one longer than 2 (two) and shorter than 12 (twelve) characters")

	ErrInvalidName = errors.New("the name is not valid. please provide one longer than 2 (two) and shorter than 60 (sixty) characters")

	ErrInvalidSalePrice = errors.New("the price is not valid. please provide an amount higher than 0")

	ErrMinimumSaleQuantity = errors.New("the minimum sale quantity provided is not valid. please provide a value higher than 0")

	ErrNotFound = errors.New("the product searched does not exist or could not be found")

	ErrAlreadyExists = errors.New("the product you are trying to create already exists. please refer to it or assign it another code")
)
