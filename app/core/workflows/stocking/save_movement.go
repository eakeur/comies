package stocking

import (
	"comies/app/core/entities/movement"
	"comies/app/core/entities/stock"
	"comies/app/sdk/fault"
	"comies/app/sdk/types"
	"context"
)

func (w workflow) SaveMovement(ctx context.Context, resourceID types.ID, mv movement.Movement) (AdditionResult, error) {

	actual, err := w.movements.GetBalanceByResourceID(ctx, resourceID, movement.Filter{})
	if err != nil {
		return AdditionResult{}, fault.Wrap(err)
	}

	stk, err := w.stocks.GetByID(ctx, resourceID)
	if err != nil {
		return AdditionResult{}, fault.Wrap(err)
	}

	mv.StockID = stk.ID
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

	_, err = w.movements.Create(ctx, mv)
	if err != nil {
		return AdditionResult{}, fault.Wrap(err)
	}

	return AdditionResult{
		Count: actual,
	}, nil
}
