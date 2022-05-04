package stock

import (
	"context"
	"gomies/app/core/entities/stock"
	"gomies/app/sdk/fault"
	"gomies/app/sdk/listing"
)

func (w workflow) ListStock(ctx context.Context, filter listing.Filter) ([]stock.Stock, int, error) {
	const operation = "Workflows.Stock.ListStock"

	list, count, err := w.stocks.ListStocks(ctx, filter)
	if err != nil {
		return []stock.Stock{}, 0, fault.Wrap(err, operation)
	}

	return list, count, nil
}
