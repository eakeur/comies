package stocking

import (
	"comies/app/core/entities/stock"
	"comies/app/sdk/fault"
	"comies/app/sdk/types"
	"context"
)

func (w workflow) GetStockByID(ctx context.Context, id types.ID) (stock.Stock, error) {

	s, err := w.stocks.GetStockByID(ctx, id)
	if err != nil {
		return stock.Stock{}, fault.Wrap(err)
	}

	return s, nil
}
