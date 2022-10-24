package menu

import (
	"comies/app/core/menu/product"
	"context"
)

func (w jobs) ListProductsRunningOut(ctx context.Context) ([]product.Product, error) {
	return w.products.ListRunningOut(ctx)
}
