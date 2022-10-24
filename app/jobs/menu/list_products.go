package menu

import (
	"comies/app/core/menu/product"
	"context"
)

func (w jobs) ListProducts(ctx context.Context, filter product.Filter) ([]product.Product, error) {

	return w.products.List(ctx, filter)
}
