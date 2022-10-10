package menu

import (
	"comies/app/core/id"
	"comies/app/core/types"
	"comies/app/data/ingredients"
	"context"
)

func RemoveIngredient(ctx context.Context, id id.ID) error {

	if id.Empty() {
		return types.ErrMissingID
	}

	err := ingredients.Remove(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
