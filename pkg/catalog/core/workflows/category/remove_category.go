package category

import (
	"context"
	"gomies/pkg/catalog/core/entities/category"
	"gomies/pkg/sdk/fault"
)

func (w workflow) RemoveCategory(ctx context.Context, ext category.Key) error {
	const operation = "Workflows.Category.RemoveCategory"

	err := w.categories.Remove(ctx, ext)
	if err != nil {
		return fault.Wrap(err, operation)
	}

	return nil
}
