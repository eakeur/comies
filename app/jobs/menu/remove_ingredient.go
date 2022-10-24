package menu

import (
	"comies/app/core/types"
	"context"
)

func (w jobs) RemoveIngredient(ctx context.Context, id types.ID) error {

	if err := id.Validate(); err != nil {
		return err
	}

	return w.ingredients.Remove(ctx, id)
}
