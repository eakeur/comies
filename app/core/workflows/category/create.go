package category

import (
	"context"
	"gomies/app/core/entities/category"
	"gomies/app/core/managers/session"
	"gomies/app/core/types/fault"
)

func (w workflow) Create(ctx context.Context, ct category.Category) (category.Category, error) {
	const operation = "Workflows.Category.Create"
	ctx = w.transactions.Begin(ctx)
	defer w.transactions.End(ctx)

	_, err := session.DelegateSessionProps(ctx, operation, &ct.Entity)
	if err != nil {
		return category.Category{}, fault.Wrap(err, operation)
	}

	if err := ct.Validate(); err != nil {
		return category.Category{}, fault.Wrap(err, operation)
	}

	ct, err = w.categories.Create(ctx, ct)
	if err != nil {
		return category.Category{}, fault.Wrap(err, operation)
	}

	return ct, nil
}
