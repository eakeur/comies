package product

import (
	"context"
	"gomies/app/core/entities/product"
	"gomies/app/core/types/fault"
	"gomies/app/core/types/id"
)

func (w workflow) Create(ctx context.Context, input CreateInput) (product.Product, error) {
	const operation = "Workflows.Product.Create"
	ctx = w.transactions.Begin(ctx)
	defer w.transactions.End(ctx)

	// If there is a category external ID assigned to the input, retrieves its internal ID
	var categoryID id.ID
	if input.CategoryID != id.Nil {
		c, err := w.categories.Get(ctx, input.CategoryID)
		if err != nil {
			return product.Product{}, fault.Wrap(err, operation)
		}
		categoryID = c.ID
	}

	// Mounts the effective product entity
	prd := product.Product{
		Code:       input.Code,
		Name:       input.Name,
		CategoryID: categoryID, // the internal ID fetched above, or zero
		Stock: product.StockInformation{
			CostPrice:       input.CostPrice,
			MaximumQuantity: input.MaximumStockQuantity,
			MinimumQuantity: input.MinimumStockQuantity,
			Location:        input.Location,
		},
		Sale: product.SaleInformation{
			Display:         input.Display,
			Description:     input.Description,
			Price:           input.Price,
			Unit:            input.Unit,
			MinimumSale:     input.MinimumSale,
			MaximumDiscount: input.MaximumDiscount,
		},
	}

	// Validates all product input
	if err := prd.Validate(); err != nil {
		return product.Product{}, fault.Wrap(err, operation)
	}

	prd, err := w.products.Create(ctx, prd)
	if err != nil {
		return product.Product{}, fault.Wrap(err, operation)
	}

	return prd, nil
}
