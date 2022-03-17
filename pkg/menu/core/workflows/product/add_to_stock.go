package product

import (
	"context"
	"gomies/pkg/menu/core/entities/product"
	"gomies/pkg/sdk/fault"
	"gomies/pkg/sdk/session"
	"gomies/pkg/sdk/types"
	"gomies/pkg/stocking/core/entities/stock"
	"time"
)

func (w workflow) AddToStock(ctx context.Context, productID types.External, movement stock.Movement) (stock.Movement, error) {
	const operation = "Workflows.Product.AddToStock"
	ctx = w.transactions.Begin(ctx)
	defer w.transactions.End(ctx)

	_, err := session.DelegateSessionProps(ctx, operation, nil, &movement.History)
	if err != nil {
		return stock.Movement{}, fault.Wrap(err, operation)
	}

	movement.TargetID = productID
	if err := movement.Validate(); err != nil {
		return stock.Movement{}, fault.Wrap(err, operation)
	}

	prd, err := w.products.Get(ctx, product.Key{ID: movement.TargetID})
	if err != nil {
		return stock.Movement{}, fault.Wrap(err, operation)
	}

	computation, err := w.stocks.ComputeStock(ctx, productID, stock.Filter{FinalDate: time.Now()})
	if err != nil {
		return stock.Movement{}, fault.Wrap(err, operation)
	}

	if computation.Actual+movement.Quantity > prd.MaximumQuantity {
		return stock.Movement{}, stock.ErrStockAlreadyFull
	}

	movement, err = w.stocks.AddToStock(ctx, productID, movement)
	if err != nil {
		return stock.Movement{}, fault.Wrap(err, operation)
	}
	return movement, nil
}
