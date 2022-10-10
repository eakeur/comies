package ingredient

import "errors"

var (
	ErrInvalidCompositeType  = errors.New("a product can only have a ingredient if it is of composite type")
	ErrInvalidIngredientType = errors.New("a ingredient can not be of output type")
	ErrInvalidIngredientID   = errors.New("the product being assigned as ingredient does not exist or is invalid")
	ErrInvalidProductID      = errors.New("the product referred by the id does not exist or is invalid")
	ErrInvalidQuantity       = errors.New("the quantity for this ingredient is invalid. please check if it is greater than 0")
)
