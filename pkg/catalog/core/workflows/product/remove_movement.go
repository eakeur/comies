package product

import (
	"context"
	"gomies/pkg/catalog/core/entities/product"
	"gomies/pkg/sdk/fault"
	"gomies/pkg/sdk/session"
	"gomies/pkg/sdk/types"
)

func (w workflow) RemoveFromStock(ctx context.Context, productKey product.Key, ext types.External) error {
	const operation = "Workflows.Product.RemoveFromStock"

	_, err := session.DelegateSessionProps(ctx, operation, &productKey.Store, nil)
	if err != nil {
		return fault.Wrap(err, operation)
	}

	ctx = w.transactions.Begin(ctx)
	defer w.transactions.End(ctx)

	_, err = w.products.Get(ctx, productKey)
	if err != nil {
		return fault.Wrap(err, operation)
	}

	err = w.stocks.RemoveFromStock(ctx, productKey.ID, ext)
	if err != nil {
		return fault.Wrap(err, operation)
	}
	return nil
}
