package stocking

import (
	"comies/app/core/entities/stock"
	"comies/app/sdk/fault"
	"context"
)

func (w workflow) UpdateStock(ctx context.Context, s stock.Stock) error {

	err := w.stocks.UpdateStock(ctx, s)
	if err != nil {
		return fault.Wrap(err)
	}

	return nil
}
