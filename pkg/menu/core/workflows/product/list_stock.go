package product

import (
	"context"
	"gomies/pkg/sdk/fault"
	"gomies/pkg/sdk/session"
	"gomies/pkg/stocking/core/entities/stock"
)

func (w workflow) ListStock(ctx context.Context, filter stock.Filter) ([]stock.Movement, error) {
	const operation = "Workflows.Product.ListStock"

	_, err := session.FromContext(ctx, operation)
	if err != nil {
		return []stock.Movement{}, fault.Wrap(err, operation)
	}

	list, err := w.stocks.ListMovements(ctx, filter)
	if err != nil {
		return []stock.Movement{}, fault.Wrap(err, operation)
	}
	return list, err
}
