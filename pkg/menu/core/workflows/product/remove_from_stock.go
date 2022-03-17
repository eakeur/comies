package product

import (
	"context"
	"gomies/pkg/sdk/fault"
	"gomies/pkg/sdk/session"
	"gomies/pkg/sdk/types"
)

func (w workflow) RemoveFromStock(ctx context.Context, productID types.External, ext types.External) error {
	const operation = "Workflows.Product.RemoveFromStock"
	ctx = w.transactions.Begin(ctx)
	defer w.transactions.End(ctx)

	_, err := session.FromContext(ctx, operation)
	if err != nil {
		return fault.Wrap(err, operation)
	}

	err = w.stocks.RemoveFromStock(ctx, productID, ext)
	if err != nil {
		return fault.Wrap(err, operation)
	}
	return nil
}
