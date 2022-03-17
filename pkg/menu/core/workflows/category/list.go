package category

import (
	"context"
	"gomies/pkg/menu/core/entities/category"
	"gomies/pkg/sdk/fault"
	"gomies/pkg/sdk/session"
)

func (w workflow) List(ctx context.Context, filter category.Filter) ([]category.Category, error) {
	const operation = "Workflows.Category.List"

	_, err := session.FromContext(ctx, operation)
	if err != nil {
		return []category.Category{}, fault.Wrap(err, operation)
	}

	ct, err := w.categories.List(ctx, filter)
	if err != nil {
		return []category.Category{}, fault.Wrap(err, operation)
	}

	return ct, nil
}
