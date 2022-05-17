package stocking

import (
	"context"
	"gomies/app/core/entities/stock"
	"gomies/app/sdk/fault"
	"gomies/app/sdk/listing"
)

func (w workflow) ListStock(ctx context.Context, filter listing.Filter) ([]stock.Stock, int, error) {

	list, count, err := w.stocks.ListStocks(ctx, filter)
	if err != nil {
		return []stock.Stock{}, 0, fault.Wrap(err)
	}

	return list, count, nil
}
