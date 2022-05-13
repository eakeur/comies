package menu

import (
	"context"
	"gomies/app/sdk/fault"
	"gomies/app/sdk/types"
)

func (w workflow) RemoveProductIngredient(ctx context.Context, id types.ID) error {
	const operation = "Workflows.Product.RemoveProductIngredient"

	if id.Empty() {
		return fault.Wrap(fault.ErrMissingUID, operation)
	}

	err := w.ingredients.RemoveIngredient(ctx, id)
	if err != nil {
		return fault.Wrap(err, operation, fault.AdditionalData{
			"id": id.String(),
		})
	}

	return nil
}
