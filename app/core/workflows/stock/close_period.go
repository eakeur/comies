package stock

import (
	"context"
	"gomies/app/core/entities/stock"
	"gomies/app/sdk/fault"
)

func (w workflow) ClosePeriod(ctx context.Context, filter stock.Filter) error {
	const operation = "Workflows.Stock.ClosePeriod"

	if err := filter.Validate(); err != nil {
		return fault.Wrap(err, operation)
	}

	err := w.stocks.ArchiveMovements(ctx, filter)
	if err != nil {
		return fault.Wrap(err, operation)
	}

	return nil
}
