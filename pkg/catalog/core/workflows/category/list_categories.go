package category

import (
	"context"
	"gomies/pkg/catalog/core/entities/category"
	"gomies/pkg/sdk/fault"
)

func (w workflow) ListCategories(ctx context.Context, filter category.Filter) ([]category.Category, error) {
	const operation = "Workflows.Product.ListCategories"

	ct, err := w.categories.List(ctx, filter)
	if err != nil {
		return []category.Category{}, fault.Wrap(err, operation)
	}

	return ct, nil
}
