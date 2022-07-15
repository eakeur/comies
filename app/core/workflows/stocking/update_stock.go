package stocking

import (
	"comies/app/core/entities/stock"
	"comies/app/sdk/throw"
	"context"
)

func (w workflow) UpdateStock(ctx context.Context, s stock.Stock) error {

	err := w.stocks.Update(ctx, s)
	if err != nil {
		return throw.Error(err)
	}

	return nil
}
