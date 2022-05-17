package stocking

import (
	"context"
	"gomies/app/core/entities/movement"
	"gomies/app/core/entities/stock"
	"gomies/app/sdk/fault"
	"gomies/app/sdk/types"
)

func (w workflow) SaveMovement(ctx context.Context, resourceID types.ID, mv movement.Movement) (AdditionResult, error) {

	actual, err := w.movements.GetBalance(ctx, movement.Filter{ResourceID: resourceID})
	if err != nil {
		return AdditionResult{}, fault.Wrap(err)
	}

	stk, err := w.stocks.GetStockByID(ctx, resourceID)
	if err != nil {
		return AdditionResult{}, fault.Wrap(err)
	}

	mv.ResourceID = resourceID
	if err := mv.Validate(); err != nil {
		return AdditionResult{}, fault.Wrap(err)
	}

	actual += mv.Value()
	if actual > stk.MaximumQuantity {
		return AdditionResult{}, fault.Wrap(stock.ErrStockFull)
	}

	if actual < stk.MinimumQuantity {
		return AdditionResult{}, fault.Wrap(stock.ErrStockEmpty)
	}

	_, err = w.movements.Save(ctx, mv)
	if err != nil {
		return AdditionResult{}, fault.Wrap(err)
	}

	return AdditionResult{
		Count: actual,
	}, nil
}
