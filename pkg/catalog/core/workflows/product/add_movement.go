package product

import (
	"context"
	"gomies/pkg/catalog/core/entities/product"
	"gomies/pkg/sdk/fault"
	"gomies/pkg/sdk/session"
	"gomies/pkg/stocking/core/entities/stock"
	"time"
)

func (w workflow) AddToStock(ctx context.Context, productKey product.Key, movement stock.Movement) (product.StockAdditionResult, error) {
	const operation = "Workflows.Product.AddToStock"

	_, err := session.DelegateSessionProps(ctx, operation, &productKey.Store, &movement.History)
	if err != nil {
		return product.StockAdditionResult{}, fault.Wrap(err, operation)
	}

	ctx = w.transactions.Begin(ctx)
	defer w.transactions.End(ctx)

	movement.TargetID = productKey.ID
	if err := movement.Validate(); err != nil {
		return product.StockAdditionResult{}, fault.Wrap(err, operation)
	}

	prd, err := w.products.GetProductStockInfo(ctx, productKey)
	if err != nil {
		return product.StockAdditionResult{}, fault.Wrap(err, operation)
	}

	computation, err := w.stocks.ComputeStock(ctx, productKey.ID, stock.Filter{FinalDate: time.Now()})
	if err != nil {
		return product.StockAdditionResult{}, fault.Wrap(err, operation)
	}

	newValue := computation.Actual + movement.Value()
	if newValue > prd.MaximumQuantity {
		return product.StockAdditionResult{}, product.ErrStockAlreadyFull
	}

	if newValue < prd.MinimumQuantity {
		return product.StockAdditionResult{}, product.ErrStockNegative
	}

	movement, err = w.stocks.AddToStock(ctx, productKey.ID, movement)
	if err != nil {
		return product.StockAdditionResult{}, fault.Wrap(err, operation)
	}
	return product.StockAdditionResult{
		Movement:       movement,
		RemainingStock: newValue,
	}, nil
}
