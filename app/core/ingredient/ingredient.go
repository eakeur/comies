package ingredient

import "comies/app/core/types"

type Ingredient struct {
	ID           types.ID       `json:"id"`
	ProductID    types.ID       `json:"product_id"`
	IngredientID types.ID       `json:"ingredient_id"`
	Quantity     types.Quantity `json:"quantity"`
	Optional     bool           `json:"optional"`
}
