package stock

import (
	"context"
	"gomies/pkg/sdk/fault"
	"gomies/pkg/stocking/core/entities/stock"
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
