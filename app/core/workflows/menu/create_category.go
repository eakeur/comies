package menu

import (
	"context"
	"gomies/app/core/entities/category"
	"gomies/app/sdk/fault"
)

func (w workflow) CreateCategory(ctx context.Context, ct category.Category) (category.Category, error) {
	const operation = "Workflows.Category.CreateCategory"

	if err := ct.Validate(); err != nil {
		return category.Category{}, fault.Wrap(err, operation)
	}

	ct, err := w.categories.CreateCategory(ctx, ct)
	if err != nil {
		return category.Category{}, fault.Wrap(err, operation)
	}

	return ct, nil
}
