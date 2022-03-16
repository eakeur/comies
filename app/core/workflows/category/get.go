package category

import (
	"context"
	"gomies/app/core/entities/category"
	"gomies/app/core/managers/session"
	"gomies/app/core/types/fault"
	"gomies/app/core/types/id"
)

func (w workflow) Get(ctx context.Context, ext id.External) (category.Category, error) {
	const operation = "Workflows.Category.Get"

	_, err := session.FromContext(ctx, operation)
	if err != nil {
		return category.Category{}, fault.Wrap(err, operation)
	}

	ct, err := w.categories.Get(ctx, ext)
	if err != nil {
		return category.Category{}, fault.Wrap(err, operation)
	}

	return ct, nil
}
