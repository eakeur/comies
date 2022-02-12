package product

import (
	"context"
	"fmt"
	"gomies/app/core/entities/product"
	"gomies/app/core/entities/stock"
	"time"
)

func (w workflow) AddToStock(ctx context.Context, movement stock.Movement) (stock.Movement, error) {
	const operation = "Workflows.Product.AddToStock"
	ctx = w.transactions.Begin(ctx)
	defer w.transactions.Rollback(ctx)

	prd, err := w.products.Get(ctx, movement.ExternalID, product.Stock)
	if err != nil {
		return stock.Movement{}, err
	}

	computation, err := w.stocks.ComputeStock(ctx, stock.Filter{TargetID: movement.TargetID, FinalDate: time.Now()})
	if err != nil {
		return stock.Movement{}, err
	}

	if computation.Actual+movement.Quantity > prd.Stock.MaximumQuantity {
		w.logger.Warn(ctx, operation, fmt.Sprintf("tried to add %v to stock with max of %v units and %v used", movement.Quantity, prd.Stock.MaximumQuantity, computation.Actual))
		return stock.Movement{}, nil
	}

	movement, err = w.stocks.AddToStock(ctx, movement)
	if err != nil {
		return stock.Movement{}, err
	}
	w.transactions.Commit(ctx)
	return movement, nil
}
