package category

import (
	"context"
	"gomies/app/core/entities/category"
	"gomies/app/core/managers/session"
	"gomies/app/core/types/fault"
)

func (w workflow) Update(ctx context.Context, ct category.Category) error {
	const operation = "Workflows.Category.Update"
	ctx = w.transactions.Begin(ctx)
	defer w.transactions.End(ctx)

	_, err := session.DelegateSessionProps(ctx, operation, &ct.Entity)
	if err != nil {
		return fault.Wrap(err, operation)
	}

	if err := ct.Validate(); err != nil {
		return fault.Wrap(err, operation)
	}

	err = w.categories.Update(ctx, ct)
	if err != nil {
		return fault.Wrap(err, operation)
	}

	return nil
}
