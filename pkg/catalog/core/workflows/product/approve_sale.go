package product

import (
	"context"
	"gomies/pkg/catalog/core/entities/product"
	"gomies/pkg/sdk/fault"
	"gomies/pkg/sdk/session"
	"gomies/pkg/stocking/core/entities/stock"
	"time"
)

func (w workflow) ApproveSale(ctx context.Context, req product.ApproveSaleRequest) error {
	const operation = "Workflows.Product.ApproveSale"

	_, err := session.DelegateSessionProps(ctx, operation, &req.Store, nil)
	if err != nil {
		return fault.Wrap(err, operation)
	}

	saleProps, err := w.products.GetProductSaleInfo(ctx, req.Key)
	if err != nil {
		return fault.Wrap(err, operation)
	}

	if saleProps.SalePrice != req.Price {
		return fault.Wrap(product.ErrInvalidSalePrice, operation)
	}

	if saleProps.MinimumSale > req.Quantity {
		return fault.Wrap(product.ErrInvalidSaleQuantity, operation)
	}

	if saleProps.HasIngredients {
		ingredients, err := w.products.ListIngredients(ctx, req.Key)
		if err != nil {
			return fault.Wrap(err, operation)
		}

		for _, ingredient := range ingredients {
			stk, err := w.stocks.ComputeStock(ctx, ingredient.IngredientExternalID, stock.Filter{FinalDate: time.Now()})
			if err != nil {
				return fault.Wrap(err, operation)
			}
			if stk.Actual < ingredient.Quantity * req.Quantity {
				return fault.Wrap(product.ErrNotEnoughStocked, operation)
			}
		}
	} else {
		stk, err := w.stocks.ComputeStock(ctx, req.ID, stock.Filter{FinalDate: time.Now()})
		if err != nil {
			return fault.Wrap(err, operation)
		}
		if stk.Actual < req.Quantity {
			return fault.Wrap(product.ErrNotEnoughStocked, operation)
		}
	}

	return nil
}
