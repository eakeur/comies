package product

import (
	"context"
	"fmt"
	"gomies/app/core/entities/product"
	"gomies/app/core/types/id"
)

func (w workflow) Create(ctx context.Context, input CreateInput) (CreateOutput, error) {
	const operation = "Workflows.Product.Create"
	w.logger.Debug(ctx, operation, "starting process")
	ctx = w.transactions.Begin(ctx)
	defer w.transactions.Rollback(ctx)

	// If there is a category external ID assigned to the input, retrieves its internal ID
	var categoryID id.ID
	if input.CategoryID != id.Nil {
		w.logger.Trace(ctx, operation, "product category informed. searching for internal id")
		c, err := w.categories.Get(ctx, input.CategoryID)
		if err != nil {
			w.logger.Warn(ctx, operation, err.Error())
			return CreateOutput{}, err
		}
		w.logger.Trace(ctx, operation, fmt.Sprintf("internal_id: %v", c.ID))
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
		w.logger.Warn(ctx, operation, err.Error())
		return CreateOutput{}, err
	}

	prd, err := w.products.Create(ctx, prd)
	if err != nil {
		w.logger.Warn(ctx, operation, err.Error())
		return CreateOutput{}, err
	}

	w.transactions.Commit(ctx)
	w.logger.Debug(ctx, operation, "finished operation")
	return CreateOutput{ID: prd.ExternalID, CreatedAt: prd.CreatedAt}, nil
}
