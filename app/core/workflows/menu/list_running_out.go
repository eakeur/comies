package menu

import (
	"comies/app/core/entities/product"
	"comies/app/sdk/throw"
	"context"
)

func (w workflow) ListProductsRunningOut(ctx context.Context) ([]product.Product, error) {
	list, err := w.products.ListRunningOut(ctx)
	if err != nil {
		return []product.Product{}, throw.Error(err)
	}

	return list, err
}
