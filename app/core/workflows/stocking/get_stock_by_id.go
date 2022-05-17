package stocking

import (
	"context"
	"gomies/app/core/entities/stock"
	"gomies/app/sdk/fault"
	"gomies/app/sdk/types"
)

func (w workflow) GetStockByID(ctx context.Context, id types.ID) (stock.Stock, error) {

	s, err := w.stocks.GetStockByID(ctx, id)
	if err != nil {
		return stock.Stock{}, fault.Wrap(err)
	}

	return s, nil
}
