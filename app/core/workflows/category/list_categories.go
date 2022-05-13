package category

import (
	"context"
	category2 "gomies/app/core/entities/category"
	"gomies/app/sdk/fault"
)

func (w workflow) ListCategories(ctx context.Context, filter category2.Filter) ([]category2.Category, int, error) {
	const operation = "Workflows.Product.ListCategories"

	ct, count, err := w.categories.ListCategories(ctx, filter)
	if err != nil {
		return []category2.Category{}, 0, fault.Wrap(err, operation)
	}

	return ct, count, nil
}
