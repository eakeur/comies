package menu

import (
	"comies/app/sdk/throw"
	"comies/app/sdk/types"
	"context"
)

func (w workflow) RemoveIngredient(ctx context.Context, id types.ID) error {

	if id.Empty() {
		return throw.Error(throw.ErrMissingID)
	}

	err := w.ingredients.Remove(ctx, id)
	if err != nil {
		return throw.Error(err).Params(map[string]interface{}{
			"id": id.String(),
		})
	}

	return nil
}
