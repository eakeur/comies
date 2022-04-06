package product

import (
	"context"
	"gomies/app/core/entities/catalog/category"
	"gomies/app/core/entities/catalog/product"
	"gomies/pkg/sdk/fault"
)

func (w workflow) UpdateProduct(ctx context.Context, prd product.Product) error {
	const operation = "Workflows.Product.UpdateProduct"

	if err := prd.Validate(); err != nil {
		return fault.Wrap(err, operation)
	}

	// If there is a category external ID assigned to the prd, retrieves its internal ID
	if !prd.CategoryExternalID.Empty() {
		c, err := w.categories.GetCategory(ctx, category.Key{ID: prd.CategoryExternalID})
		if err != nil {
			return fault.Wrap(err, operation)
		}
		prd.CategoryID = c.Entity.ID
	} else {
		prd.CategoryID = 0
	}

	err := w.products.UpdateProduct(ctx, prd)
	if err != nil {
		return fault.Wrap(err, operation)
	}

	return nil
}
