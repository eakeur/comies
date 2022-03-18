package stock

import (
	"context"
	"gomies/pkg/sdk/fault"
	"gomies/pkg/stocking/core/entities/stock"
)

func (w workflow) ListMovements(ctx context.Context, filter stock.Filter) ([]stock.Movement, error) {
	const operation = "Workflows.Stock.ListMovements"

	if err := filter.Validate(); err != nil {
		return []stock.Movement{}, fault.Wrap(err, operation)
	}

	movements, err := w.stocks.ListMovements(ctx, filter)
	if err != nil {
		return []stock.Movement{}, fault.Wrap(err, operation)
	}

	return movements, nil
}
