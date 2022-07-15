package stocking

import (
	"comies/app/core/entities/stock"
	"comies/app/sdk/fault"
	"context"
)

func (w workflow) CreateStock(ctx context.Context, s stock.Stock) (stock.Stock, error) {

	s, err := w.stocks.Create(ctx, s)
	if err != nil {
		return stock.Stock{}, fault.Wrap(err)
	}

	return s, nil
}
