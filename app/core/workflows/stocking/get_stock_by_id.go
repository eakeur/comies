package stocking

import (
	"comies/app/core/entities/stock"
	"comies/app/sdk/throw"
	"comies/app/sdk/types"
	"context"
)

func (w workflow) GetStockByID(ctx context.Context, id types.ID) (stock.Stock, error) {

	s, err := w.stocks.GetByID(ctx, id)
	if err != nil {
		return stock.Stock{}, throw.Error(err)
	}

	return s, nil
}
