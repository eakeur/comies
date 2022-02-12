package product

import (
	"context"
	"gomies/app/core/types/fault"
	"gomies/app/core/types/id"
)

func (w workflow) Remove(ctx context.Context, ext id.External) error {
	const operation = "Workflows.Product.Remove"
	ctx = w.transactions.Begin(ctx)
	defer w.transactions.End(ctx)

	err := w.products.Remove(ctx, ext)
	if err != nil {
		return fault.Wrap(err, operation)
	}

	return err

}
