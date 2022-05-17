package menu

import (
	"context"
	"gomies/app/core/entities/category"
	"gomies/app/sdk/fault"
)

func (w workflow) UpdateCategory(ctx context.Context, c category.Category) error {

	if err := c.Validate(); err != nil {
		return fault.Wrap(err)
	}

	err := w.categories.UpdateCategory(ctx, c)
	if err != nil {
		return fault.Wrap(err)
	}

	return nil
}
