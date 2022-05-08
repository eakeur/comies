package ingredient

import "errors"

var (
	ErrInvalidIngredientID = errors.New("the product being assigned as ingredient does not exist or is invalid")
	ErrInvalidProductID    = errors.New("the product referred by the id does not exist or is invalid")
	ErrInvalidQuantity     = errors.New("the quantity for this ingredient is invalid. please check if it is greater than 0")
)
