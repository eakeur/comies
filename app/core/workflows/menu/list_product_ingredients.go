package menu

import (
	"context"
	"gomies/app/core/entities/ingredient"
	"gomies/app/sdk/fault"
	"gomies/app/sdk/types"
)

func (w workflow) ListProductIngredients(ctx context.Context, productID types.ID) ([]ingredient.Ingredient, error) {
	const operation = "Workflows.Product.AddProductIngredient"

	if productID.Empty() {
		return []ingredient.Ingredient{}, fault.Wrap(fault.ErrMissingUID, operation)
	}

	list, err := w.ingredients.ListIngredients(ctx, productID)
	if err != nil {
		return []ingredient.Ingredient{}, fault.Wrap(err, operation, fault.AdditionalData{
			"product_id": productID.String(),
		})
	}

	return list, nil
}
