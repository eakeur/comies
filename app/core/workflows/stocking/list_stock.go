package stocking

import (
	"context"
	"gomies/app/core/entities/stock"
	"gomies/app/sdk/fault"
)

func (w workflow) ListStock(ctx context.Context) ([]stock.Stock, error) {

	list, err := w.stocks.ListStocks(ctx)
	if err != nil {
		return []stock.Stock{}, fault.Wrap(err)
	}

	return list, nil
}
