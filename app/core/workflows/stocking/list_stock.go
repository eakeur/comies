package stocking

import (
	"context"
	"gomies/app/core/entities/stock"
	"gomies/app/sdk/fault"
)

func (w workflow) ListStock(ctx context.Context) ([]stock.Stock, int, error) {

	list, count, err := w.stocks.ListStocks(ctx)
	if err != nil {
		return []stock.Stock{}, 0, fault.Wrap(err)
	}

	return list, count, nil
}
