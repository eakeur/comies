package product

import (
	"context"
	"gomies/app/core/entities/product"
	"gomies/app/core/managers/session"
	"gomies/app/core/types/fault"
	"gomies/app/core/types/id"
)

func (w workflow) Update(ctx context.Context, prd product.Product) error {
	const operation = "Workflows.Product.Update"
	w.transactions.Begin(ctx)
	defer w.transactions.End(ctx)

	_, err := session.DelegateSessionProps(ctx, operation, &prd.Entity)
	if err != nil {
		return fault.Wrap(err, operation)
	}

	if err := prd.Validate(); err != nil {
		return fault.Wrap(err, operation)
	}

	// If there is a category external ID assigned to the input, retrieves its internal ID
	if prd.CategoryExternalID != id.Nil && (prd.CategoryExternalID != id.External{}) {
		c, err := w.categories.Get(ctx, prd.CategoryExternalID)
		if err != nil {
			return fault.Wrap(err, operation)
		}
		prd.CategoryID = c.ID
	}

	err = w.products.Update(ctx, prd)
	if err != nil {
		return fault.Wrap(err, operation)
	}

	return nil
}
