package menu

import (
	"comies/app/core/throw"
	"comies/app/core/types"
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
