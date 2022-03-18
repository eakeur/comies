package product

import (
	"context"
	"gomies/pkg/catalog/core/entities/product"
	"gomies/pkg/sdk/fault"
	"gomies/pkg/sdk/session"
)

func (w workflow) AddIngredient(ctx context.Context, productKey product.Key, ingredient product.Ingredient) (product.Ingredient, error) {
	const operation = "Workflows.Product.AddIngredient"

	ses, err := session.DelegateSessionProps(ctx, operation, &ingredient.Store, &ingredient.History)
	if err != nil {
		return product.Ingredient{}, fault.Wrap(err, operation)
	}
	ses.Delegate(operation, &productKey.Store, nil)

	if err := ingredient.Validate(); err != nil {
		return product.Ingredient{}, fault.Wrap(err, operation)
	}

	ctx = w.transactions.Begin(ctx)
	defer w.transactions.End(ctx)

	ingredient.ProductExternalID = productKey.ID
	_, err = w.products.SaveIngredients(ctx, productKey, ingredient)
	if err != nil {
		return product.Ingredient{}, fault.Wrap(err, operation)
	}

	return ingredient, nil
}
