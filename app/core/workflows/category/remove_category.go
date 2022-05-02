package category

import (
	"context"
	"gomies/app/core/entities/catalog/category"
	"gomies/app/sdk/fault"
)

func (w workflow) RemoveCategory(ctx context.Context, ext category.Key) error {
	const operation = "Workflows.Category.RemoveCategory"

	err := w.categories.RemoveCategory(ctx, ext)
	if err != nil {
		return fault.Wrap(err, operation)
	}

	return nil
}
