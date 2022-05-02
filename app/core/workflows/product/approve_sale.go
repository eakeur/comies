package product

import (
	"context"
	"gomies/app/core/entities/catalog/product"
	"gomies/app/sdk/fault"
	"gomies/app/sdk/types"
)

func (w workflow) ApproveSale(ctx context.Context, req ApproveSaleRequest) error {
	const operation = "Workflows.Product.ApproveSale"

	saleProps, err := w.approveSaleProperties(ctx, req)
	if err != nil {
		return fault.Wrap(err, operation)
	}

	err = w.verifyStockAvailability(ctx, req, saleProps.HasIngredients)
	if err != nil {
		return fault.Wrap(err, operation)
	}

	return nil
}

func (w workflow) approveSaleProperties(ctx context.Context, req ApproveSaleRequest) (product.Sale, error) {
	const operation = "Workflows.Product.approveSaleProperties"

	saleProps, err := w.products.GetProductSaleInfo(ctx, product.Key{ID: req.ProductID})
	if err != nil {
		return product.Sale{}, fault.Wrap(err, operation)
	}

	if saleProps.SalePrice != req.Price {
		return product.Sale{}, fault.Wrap(product.ErrInvalidSalePrice, operation)
	}

	if saleProps.MinimumSale > req.Quantity {
		return product.Sale{}, fault.Wrap(product.ErrInvalidSaleQuantity, operation)
	}

	return saleProps, nil
}

func (w workflow) verifyStockAvailability(ctx context.Context, req ApproveSaleRequest, hasIngredients bool) error {
	const operation = "Workflows.Product.verifyStockAvailability"

	computations, err := w.fetchStockComputation(ctx, req.ProductID, hasIngredients, req.IngredientsToReplace, req.IngredientsToIgnore)
	if err != nil {
		return fault.Wrap(err, operation)
	}

	for _, calc := range computations {
		if calc.AvailableQuantity < calc.Quantity*req.Quantity {
			return product.ErrNotEnoughStocked
		}
	}

	return nil
}

func (w workflow) fetchStockComputation(ctx context.Context, productID types.ID, hasIngredients bool, substitutions Substitutions, ignoring []types.ID) ([]IngredientToVerify, error) {
	const operation = "Workflows.Product.fetchStockComputation"

	if !hasIngredients {
		stk, err := w.stocks.Compute(ctx, productID)
		if err != nil {
			return []IngredientToVerify{}, fault.Wrap(err, operation)
		}

		return []IngredientToVerify{
			{
				IngredientID:      productID,
				Quantity:          1,
				AvailableQuantity: stk,
			},
		}, nil
	}

	ingredients, err := w.products.ListIngredients(ctx, product.Key{ID: productID})
	if err != nil {
		return []IngredientToVerify{}, fault.Wrap(err, operation)
	}

	ids := reduceIngredients(ingredients, substitutions, ignoring)

	computations, err := w.stocks.ComputeSome(ctx, ids...)
	if err != nil {
		return []IngredientToVerify{}, fault.Wrap(err, operation)
	}

	computationMap := make([]IngredientToVerify, len(computations))
	for i, calc := range computations {
		ingredient := ingredients[i]

		computationMap[i] = IngredientToVerify{
			IngredientID:      ingredient.ID,
			Quantity:          ingredient.Quantity,
			AvailableQuantity: calc,
		}
	}

	return []IngredientToVerify{}, nil
}

// reduceIngredients filters the ingredients array, removing the ignored items and the substituted ones
// and reduces it, returning an types.ID array
func reduceIngredients(list []product.Ingredient, substitutions Substitutions, ignored []types.ID) []types.ID {
	ingrendientsIDs := make([]types.ID, len(list))
	for i, ingredient := range list {
		if sub, ok := substitutions[ingredient.IngredientID]; ok {
			ingrendientsIDs[i] = sub
			continue
		}

		var shouldIgnore bool
		for _, ignore := range ignored {
			if ignore == ingredient.ID {
				shouldIgnore = true
				break
			}
		}

		if shouldIgnore {
			ingrendientsIDs[i] = ingredient.IngredientID
		}

	}

	return ingrendientsIDs
}
