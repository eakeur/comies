package stock

import (
	"context"
	"gomies/app/core/entities/stock"
	"gomies/app/sdk/fault"
)

func (w workflow) CreateStock(ctx context.Context, s stock.Stock) (stock.Stock, error) {
	const operation = "Workflows.Stock.CreateStock"

	s, err := w.stocks.CreateStock(ctx, s)
	if err != nil {
		return stock.Stock{}, fault.Wrap(err, operation)
	}

	return s, nil
}
