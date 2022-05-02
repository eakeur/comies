package category

import (
	"context"
	"gomies/app/core/entities/catalog/category"
	"gomies/app/sdk/fault"
)

func (w workflow) UpdateCategory(ctx context.Context, c category.Category) error {
	const operation = "Workflows.Product.UpdateCategory"

	if err := c.Validate(); err != nil {
		return fault.Wrap(err, operation)
	}

	err := w.categories.UpdateCategory(ctx, c)
	if err != nil {
		return fault.Wrap(err, operation)
	}

	return nil
}
