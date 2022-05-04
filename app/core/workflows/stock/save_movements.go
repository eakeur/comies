package stock

import (
	"context"
	"gomies/app/core/entities/stock"
	"gomies/app/sdk/fault"
	"gomies/app/sdk/types"
)

func (w workflow) SaveMovements(ctx context.Context, resourceID types.ID, movement stock.Movement) (stock.AdditionResult, error) {
	const operation = "Workflows.Stock.SaveMovements"

	actual, err := w.stocks.ComputeStock(ctx, stock.Filter{ResourceID: resourceID})
	if err != nil {
		return stock.AdditionResult{}, fault.Wrap(err, operation)
	}

	stk, err := w.stocks.GetStockByID(ctx, resourceID)
	if err != nil {
		return stock.AdditionResult{}, fault.Wrap(err, operation)
	}

	movement.ResourceID = resourceID
	if err := movement.Validate(); err != nil {
		return stock.AdditionResult{}, fault.Wrap(err, operation)
	}

	actual += movement.Value()
	if actual > stk.MaximumQuantity {
		return stock.AdditionResult{}, fault.Wrap(stock.ErrStockFull, operation)
	}

	if actual < stk.MinimumQuantity {
		return stock.AdditionResult{}, fault.Wrap(stock.ErrStockEmpty, operation)
	}

	_, err = w.stocks.SaveMovements(ctx, movement)
	if err != nil {
		return stock.AdditionResult{}, fault.Wrap(err, operation)
	}

	return stock.AdditionResult{
		Count: actual,
	}, nil
}
