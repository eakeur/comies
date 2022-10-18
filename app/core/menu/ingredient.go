package menu

import (
	"comies/app/core/id"
	"comies/app/core/types"
)

type Ingredient struct {
	ID           id.ID          `json:"id"`
	ProductID    id.ID          `json:"product_id"`
	IngredientID id.ID          `json:"ingredient_id"`
	Quantity     types.Quantity `json:"quantity"`
	Optional     bool           `json:"optional"`
}
type IngredientSpecification map[id.ID]Specification

type Specification struct {
	ChangeType Type  `json:"change_type"`
	ReplaceBy  id.ID `json:"replace_by"`
}

func ValidateIngredient(i Ingredient) error {
	if err := id.ValidateID(i.IngredientID); err != nil {
		return ErrInvalidComponentID
	}

	if err := id.ValidateID(i.ProductID); err != nil {
		return ErrInvalidCompositeID
	}

	if i.Quantity <= 0 {
		return ErrInvalidQuantity
	}

	return nil
}

func ModifyIngredientsList(list []Ingredient, changers IngredientSpecification, wantQuantity types.Quantity) (ingredients []Ingredient) {
	for _, i := range list {
		spec := changers[i.ID]

		if spec.ChangeType == IgnoreIngredientChangeType {
			continue
		}

		if spec.ChangeType == ReplaceIngredientChangeType {
			i.IngredientID = spec.ReplaceBy
		}

		i.Quantity *= wantQuantity

		ingredients = append(ingredients, i)
	}

	return
}
