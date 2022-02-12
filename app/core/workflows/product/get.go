package product

import (
	"context"
	"gomies/app/core/entities/product"
	"gomies/app/core/types/id"
)

func (w workflow) Get(ctx context.Context, ext id.External) (product.Product, error) {
	const operation = "Workflows.Product.Get"
	w.logger.Debug(ctx, operation, "starting process")
	prod, err := w.products.Get(ctx, ext, product.All)
	if err != nil {
		w.logger.Warn(ctx, operation, err.Error())
		return product.Product{}, err
	}
	w.logger.Debug(ctx, operation, "finished process")
	return prod, nil
}
