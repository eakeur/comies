package product

import (
	"context"
	"gomies/app/core/entities/stock"
	"gomies/app/core/types/fault"
)

func (w workflow) ListStock(ctx context.Context, filter stock.Filter) ([]stock.Movement, error) {
	const operation = "Workflows.Product.ListStock"
	list, err := w.stocks.ListMovements(ctx, filter)
	if err != nil {
		return []stock.Movement{}, fault.Wrap(err, operation)
	}
	return list, err
}
