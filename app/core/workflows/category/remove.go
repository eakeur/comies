package category

import (
	"context"
	"gomies/app/core/managers/session"
	"gomies/app/core/types/fault"
	"gomies/app/core/types/id"
)

func (w workflow) Remove(ctx context.Context, ext id.External) error {
	const operation = "Workflows.Category.Remove"
	ctx = w.transactions.Begin(ctx)
	defer w.transactions.End(ctx)

	_, err := session.FromContext(ctx, operation)
	if err != nil {
		return fault.Wrap(err, operation)
	}

	err = w.categories.Remove(ctx, ext)
	if err != nil {
		return fault.Wrap(err, operation)
	}

	return nil
}
