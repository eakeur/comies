package category

import (
	"context"
	"gomies/pkg/menu/core/entities/category"
	"gomies/pkg/sdk/fault"
	"gomies/pkg/sdk/session"
	"gomies/pkg/sdk/types"
)

func (w workflow) Save(ctx context.Context, ct category.Category, flag ...types.WritingFlag) (category.Category, error) {
	const operation = "Workflows.Category.Save"
	ctx = w.transactions.Begin(ctx)
	defer w.transactions.End(ctx)

	_, err := session.DelegateSessionProps(ctx, operation, &ct.Store, &ct.History)
	if err != nil {
		return category.Category{}, fault.Wrap(err, operation)
	}

	if err := ct.Validate(); err != nil {
		return category.Category{}, fault.Wrap(err, operation)
	}

	ct, err = w.categories.Save(ctx, ct, flag...)
	if err != nil {
		return category.Category{}, fault.Wrap(err, operation)
	}

	return ct, nil
}
