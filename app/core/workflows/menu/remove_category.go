package menu

import (
	"context"
	"gomies/app/core/entities/category"
	"gomies/app/sdk/fault"
)

func (w workflow) RemoveCategory(ctx context.Context, ext category.Key) error {

	err := w.categories.RemoveCategory(ctx, ext)
	if err != nil {
		return fault.Wrap(err)
	}

	return nil
}
