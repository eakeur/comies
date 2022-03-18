package product

import (
	"context"
	"gomies/pkg/catalog/core/entities/category"
	"gomies/pkg/catalog/core/entities/product"
	"gomies/pkg/sdk/fault"
	"gomies/pkg/sdk/session"
	"gomies/pkg/sdk/types"
)

func (w workflow) SaveProduct(ctx context.Context, input product.Product, flag ...types.WritingFlag) (product.Product, error) {
	const operation = "Workflows.Product.SaveProduct"

	_, err := session.DelegateSessionProps(ctx, operation, &input.Store, &input.History)
	if err != nil {
		return product.Product{}, fault.Wrap(err, operation)
	}

	ctx = w.transactions.Begin(ctx)
	defer w.transactions.End(ctx)

	if err := input.Validate(); err != nil {
		return product.Product{}, fault.Wrap(err, operation)
	}

	// If there is a category external ID assigned to the input, retrieves its internal ID
	if !input.CategoryExternalID.Empty() {
		c, err := w.categories.Get(ctx, category.Key{ID: input.CategoryExternalID})
		if err != nil {
			return product.Product{}, fault.Wrap(err, operation)
		}
		input.CategoryID = c.Entity.ID
	} else {
		input.CategoryID = 0
	}

	input, err = w.products.Save(ctx, input, flag...)
	if err != nil {
		return product.Product{}, fault.Wrap(err, operation)
	}

	return input, nil
}
