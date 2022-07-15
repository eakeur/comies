package stocking

import (
	"comies/app/core/entities/stock"
	"comies/app/sdk/fault"
	"context"
)

func (w workflow) ListStock(ctx context.Context) ([]stock.Stock, error) {

	list, err := w.stocks.ListStocks(ctx)
	if err != nil {
		return []stock.Stock{}, fault.Wrap(err)
	}

	return list, nil
}
