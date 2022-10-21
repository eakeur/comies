package ingredient

import "errors"

// Ingredients errors
var (
	ErrInvalidComponentID = errors.New("the product being assigned as ingredient does not exist or is invalid. a ingredient can not be of output type")
	ErrInvalidCompositeID = errors.New("the product referred by the id does not exist or is invalid. a product can only have a ingredient if it is of composite type")
	ErrInvalidQuantity    = errors.New("the quantity for this ingredient is invalid. please check if it is greater than 0")
)
