package category

import (
	"context"
	"gomies/app/core/entities/catalog/category"
	"gomies/app/sdk/fault"
)

func (w workflow) ListCategories(ctx context.Context, filter category.Filter) ([]category.Category, int, error) {
	const operation = "Workflows.Product.ListCategories"

	ct, count, err := w.categories.ListCategories(ctx, filter)
	if err != nil {
		return []category.Category{}, 0, fault.Wrap(err, operation)
	}

	return ct, count, nil
}
