package menu

import (
	"context"
	"gomies/app/sdk/fault"
	"gomies/app/sdk/types"
)

func (w workflow) RemoveProductIngredient(ctx context.Context, id types.ID) error {

	if id.Empty() {
		return fault.Wrap(fault.ErrMissingID)
	}

	err := w.ingredients.Remove(ctx, id)
	if err != nil {
		return fault.Wrap(err).Params(map[string]interface{}{
			"id": id.String(),
		})
	}

	return nil
}
