package product

import (
	"context"
	"gomies/pkg/catalog/core/entities/product"
	"gomies/pkg/sdk/fault"
	"gomies/pkg/sdk/types"
	"gomies/pkg/stocking/core/entities/stock"
)

func (w workflow) ApproveSale(ctx context.Context, req product.ApproveSaleRequest) error {
	const operation = "Workflows.Product.ApproveSale"

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

	if !saleProps.HasIngredients {
		stk, err := w.stocks.Compute(ctx, stock.Filter{ResourceID: req.ID})
		if err != nil {
			return fault.Wrap(err, operation)
		}
		if stk < req.Quantity {
			return fault.Wrap(product.ErrNotEnoughStocked, operation)
		}

		return nil
	}

	err = w.checkIngredients(ctx, req.Key, req.Quantity)
	if err != nil {
		return fault.Wrap(err, operation)
	}

	return nil
}

func (w workflow) checkIngredients(ctx context.Context, productKey product.Key, quantity types.Quantity) error {
	ingredients, err := w.products.ListIngredients(ctx, productKey)
	if err != nil {
		return err
	}

	ingredientIDs := make([]types.UID, len(ingredients))
	for i, ingredient := range ingredients {
		ingredientIDs[i] = ingredient.IngredientExternalID
	}

	calc, err := w.stocks.ComputeSome(ctx, stock.Filter{}, ingredientIDs...)
	if err != nil {
		return err
	}

	for i, c := range calc {
		if c < ingredients[i].Quantity*quantity {
			return product.ErrNotEnoughStocked
		}
	}

	return nil
}
