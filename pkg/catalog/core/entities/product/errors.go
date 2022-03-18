package product

import "errors"

var (
	ErrInvalidCode = errors.New("the code is not valid. please provide one longer than 2 (two) and shorter than 12 (twelve) characters")

	ErrInvalidName = errors.New("the name is not valid. please provide one longer than 2 (two) and shorter than 60 (sixty) characters")

	ErrInvalidPrice = errors.New("the price is not valid. please provide an amount higher than 0")

	ErrMinimumSaleQuantity = errors.New("the minimum sale quantity provided is not valid. please provide a value higher than 0")

	ErrInvalidSalePrice = errors.New("the sale price requested for this product is invalid. please check the correct one")

	ErrInvalidSaleQuantity = errors.New("the sale quantity requested for this product is invalid. please check the correct one")

	ErrNotEnoughStocked = errors.New("the stock for this product cannot afford this sale")

	ErrInvalidIngredient = errors.New("the ingredient for this product is invalid. please check if you assigned it an id")

	ErrInvalidIngredientQuantity = errors.New("the ingredient quantity for this product is invalid. please check if it is greater than 0")

	ErrStockAlreadyFull = errors.New("the stock is already full")

	ErrStockNegative = errors.New("the stock, after this output, would be negative. that cannot happen")
)
