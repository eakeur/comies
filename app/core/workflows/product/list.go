package product

import (
	"context"
	"gomies/app/core/entities/product"
)

func (w workflow) List(ctx context.Context, filter product.Filter) ([]product.Product, error) {
	const operation = "Workflows.Product.List"
	w.logger.Debug(ctx, operation, "starting process")
	list, err := w.products.List(ctx, filter)
	if err != nil {
		w.logger.Warn(ctx, operation, err.Error())
		return []product.Product{}, err
	}
	w.logger.Debug(ctx, operation, "finished process")
	return list, err
}
