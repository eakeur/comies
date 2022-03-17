package category

import (
	"context"
	"gomies/pkg/sdk/fault"
	"gomies/pkg/sdk/session"
	"gomies/pkg/sdk/types"
)

func (w workflow) Remove(ctx context.Context, ext types.External) error {
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
