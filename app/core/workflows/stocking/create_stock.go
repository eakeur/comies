package stocking

import (
	"context"
	"gomies/app/core/entities/stock"
	"gomies/app/sdk/fault"
)

func (w workflow) CreateStock(ctx context.Context, s stock.Stock) (stock.Stock, error) {

	s, err := w.stocks.CreateStock(ctx, s)
	if err != nil {
		return stock.Stock{}, fault.Wrap(err)
	}

	return s, nil
}
