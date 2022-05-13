package ingredient

import (
	"gomies/app/sdk/types"
)

type (
	Ingredient struct {
		ID           types.ID
		ProductID    types.ID
		IngredientID types.ID
		Quantity     types.Quantity
		Optional     bool
		Active       bool
		History      types.History
		types.Store
	}

	IgnoredList []types.ID

	ReplacedList map[types.ID]types.ID
)

func (i Ingredient) Validate() error {
	if i.IngredientID.Empty() {
		return ErrInvalidIngredientID
	}

	if i.ProductID.Empty() {
		return ErrInvalidProductID
	}

	if i.Quantity <= 0 {
		return ErrInvalidQuantity
	}

	return nil
}

func IgnoreAndReplace(list []Ingredient, ignored IgnoredList, replaced ReplacedList, modifier func(i Ingredient) Ingredient) (ingredients []Ingredient) {
	for _, ingredient := range list {
		var ignore bool

		i := Ingredient{
			IngredientID: ingredient.IngredientID,
			Quantity:     ingredient.Quantity,
		}

		if replacing, ok := replaced[ingredient.IngredientID]; ok {
			i.IngredientID = replacing

			if modifier != nil {
				i = modifier(i)
			}

			ingredients = append(ingredients, i)
			continue
		}

		for _, ignoring := range ignored {
			if ignore = ignoring == ingredient.IngredientID; ignore {
				break
			}
		}

		if !ignore {
			if modifier != nil {
				i = modifier(i)
			}

			ingredients = append(ingredients, i)
		}
	}

	return ingredients
}
