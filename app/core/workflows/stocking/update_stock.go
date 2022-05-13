package stocking

import (
	"context"
	"gomies/app/core/entities/stock"
	"gomies/app/sdk/fault"
)

func (w workflow) UpdateStock(ctx context.Context, s stock.Stock) error {
	const operation = "Workflows.Stock.UpdateStock"

	err := w.stocks.UpdateStock(ctx, s)
	if err != nil {
		return fault.Wrap(err, operation)
	}

	return nil
}
