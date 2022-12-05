package ingredient

import "comies/core/types"

type Ingredient struct {
	ProductID    types.ID       `json:"product_id"`
	IngredientID types.ID       `json:"ingredient_id"`
	Quantity     types.Quantity `json:"quantity"`
	Optional     bool           `json:"optional"`
}

func (i Ingredient) Validate() (Ingredient, error) {
	if err := i.IngredientID.Validate(); err != nil {
		return i, ErrInvalidComponentID
	}

	if err := i.ProductID.Validate(); err != nil {
		return i, ErrInvalidCompositeID
	}

	if i.Quantity <= 0 {
		return i, ErrInvalidQuantity
	}

	return i, nil
}
