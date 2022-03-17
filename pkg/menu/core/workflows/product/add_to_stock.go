package product

import (
	"context"
	"gomies/pkg/menu/core/entities/product"
	"gomies/pkg/sdk/fault"
	"gomies/pkg/sdk/session"
	"gomies/pkg/stocking/core/entities/stock"
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

	prd, err := w.products.Get(ctx, product.Key{ID: movement.TargetID})
	if err != nil {
		return stock.Movement{}, fault.Wrap(err, operation)
	}

	computation, err := w.stocks.ComputeStock(ctx, stock.Filter{TargetID: movement.TargetID, FinalDate: time.Now()})
	if err != nil {
		return stock.Movement{}, fault.Wrap(err, operation)
	}

	if computation.Actual+movement.Quantity > prd.MaximumQuantity {
		return stock.Movement{}, stock.ErrStockAlreadyFull
	}

	movement, err = w.stocks.AddToStock(ctx, movement)
	if err != nil {
		return stock.Movement{}, fault.Wrap(err, operation)
	}
	return movement, nil
}
