package stock

import (
	"context"
	"gomies/app/core/entities/stock"
	"gomies/pkg/sdk/fault"
)

func (w workflow) ListMovements(ctx context.Context, filter stock.Filter) ([]stock.Movement, int, error) {
	const operation = "Workflows.Stock.ListMovements"

	if err := filter.Validate(); err != nil {
		return []stock.Movement{}, 0, fault.Wrap(err, operation)
	}

	movements, count, err := w.stocks.ListMovements(ctx, filter)
	if err != nil {
		return []stock.Movement{}, 0, fault.Wrap(err, operation)
	}

	return movements, count, nil
}
