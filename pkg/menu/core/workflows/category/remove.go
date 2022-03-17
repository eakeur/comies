package category

import (
	"context"
	"gomies/pkg/menu/core/entities/category"
	"gomies/pkg/sdk/fault"
	"gomies/pkg/sdk/session"
)

func (w workflow) Remove(ctx context.Context, ext category.Key) error {
	const operation = "Workflows.Category.Remove"
	ctx = w.transactions.Begin(ctx)
	defer w.transactions.End(ctx)

	_, err := session.DelegateSessionProps(ctx, operation, &ext.Store, nil)
	if err != nil {
		return fault.Wrap(err, operation)
	}

	err = w.categories.Remove(ctx, ext)
	if err != nil {
		return fault.Wrap(err, operation)
	}

	return nil
}
