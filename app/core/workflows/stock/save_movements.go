package stock

import (
	"context"
	"gomies/app/core/entities/stock"
	"gomies/app/sdk/fault"
	"gomies/app/sdk/types"
)

func (w workflow) SaveMovements(ctx context.Context, config stock.Config, resourceID types.ID, movements ...stock.Movement) (stock.AdditionResult, error) {
	const operation = "Workflows.Stock.SaveMovements"

	actual, err := w.stocks.Compute(ctx, stock.Filter{ResourceID: resourceID})
	if err != nil {
		return stock.AdditionResult{}, fault.Wrap(err, operation)
	}

	for _, movement := range movements {
		movement.TargetID = resourceID
		if err := movement.Validate(); err != nil {
			return stock.AdditionResult{}, fault.Wrap(err, operation)
		}

		actual += movement.Value()
		if actual > config.MaximumQuantity {
			return stock.AdditionResult{}, fault.Wrap(stock.ErrStockFull, operation)
		}

		if actual < config.MinimumQuantity {
			return stock.AdditionResult{}, fault.Wrap(stock.ErrStockEmpty, operation)
		}
	}

	_, err = w.stocks.SaveMovements(ctx, movements...)
	if err != nil {
		return stock.AdditionResult{}, fault.Wrap(err, operation)
	}

	return stock.AdditionResult{
		Count: actual,
	}, nil
}
