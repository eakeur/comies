package product

import (
	"context"
	"gomies/app/core/entities/catalog/product"
	"gomies/pkg/sdk/fault"
	"gomies/pkg/sdk/types"
)

func (w workflow) CreateMovement(ctx context.Context, productID types.UID, mov Movement) (types.Quantity, error) {
	const operation = "Workflows.Product.CreateMovement"

	stk, err := w.products.GetProductStockInfo(ctx, product.Key{ID: productID})
	if err != nil {
		return 0, fault.Wrap(err, operation)
	}

	mov.ProductID = productID
	remaining, err := w.stocks.CreateMovement(ctx, stk, productID, mov)
	if err != nil {
		return 0, fault.Wrap(err, operation)
	}

	return remaining, nil
}
