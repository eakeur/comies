package menu

import (
	"context"
	"gomies/app/core/entities/category"
	"gomies/app/sdk/fault"
)

func (w workflow) ListCategories(ctx context.Context, filter category.Filter) ([]category.Category, int, error) {

	ct, count, err := w.categories.ListCategories(ctx, filter)
	if err != nil {
		return []category.Category{}, 0, fault.Wrap(err)
	}

	return ct, count, nil
}
