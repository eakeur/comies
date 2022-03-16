package product

import (
	"context"
	"gomies/app/core/entities/product"
	"gomies/app/core/entities/stock"
	"gomies/app/core/managers/session"
	"gomies/app/core/types/fault"
	"time"
)

func (w workflow) AddToStock(ctx context.Context, movement stock.Movement) (stock.Movement, error) {
	const operation = "Workflows.Product.AddToStock"
	ctx = w.transactions.Begin(ctx)
	defer w.transactions.End(ctx)

	_, err := session.DelegateSessionProps(ctx, operation, &movement.Entity)
	if err != nil {
		return stock.Movement{}, fault.Wrap(err, operation)
	}

	if err := movement.Validate(); err != nil {
		return stock.Movement{}, fault.Wrap(err, operation)
	}

	prd, err := w.products.Get(ctx, movement.TargetID, product.Stock)
	if err != nil {
		return stock.Movement{}, fault.Wrap(err, operation)
	}

	computation, err := w.stocks.ComputeStock(ctx, stock.Filter{TargetID: movement.TargetID, FinalDate: time.Now()})
	if err != nil {
		return stock.Movement{}, fault.Wrap(err, operation)
	}

	if computation.Actual+movement.Quantity > prd.Stock.MaximumQuantity {
		return stock.Movement{}, stock.ErrStockAlreadyFull
	}

	movement, err = w.stocks.AddToStock(ctx, movement)
	if err != nil {
		return stock.Movement{}, fault.Wrap(err, operation)
	}
	return movement, nil
}
