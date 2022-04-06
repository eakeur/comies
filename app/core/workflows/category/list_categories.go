package category

import (
	"context"
	"gomies/app/core/entities/catalog/category"
	"gomies/pkg/sdk/fault"
)

func (w workflow) ListCategories(ctx context.Context, filter category.Filter) ([]category.Category, error) {
	const operation = "Workflows.Product.ListCategories"

	ct, err := w.categories.ListCategories(ctx, filter)
	if err != nil {
		return []category.Category{}, fault.Wrap(err, operation)
	}

	return ct, nil
}
