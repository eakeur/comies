package category

import (
	"context"
	"gomies/pkg/catalog/core/entities/category"
	"gomies/pkg/sdk/fault"
	"gomies/pkg/sdk/session"
)

func (w workflow) ListCategories(ctx context.Context, filter category.Filter) ([]category.Category, error) {
	const operation = "Workflows.Product.ListCategories"

	_, err := session.DelegateSessionProps(ctx, operation, &filter.Store, nil)
	if err != nil {
		return []category.Category{}, fault.Wrap(err, operation)
	}

	ct, err := w.categories.List(ctx, filter)
	if err != nil {
		return []category.Category{}, fault.Wrap(err, operation)
	}

	return ct, nil
}
