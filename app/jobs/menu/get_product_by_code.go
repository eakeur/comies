package menu

import (
	"comies/app/core/product"
	"comies/app/data/products"
	"context"
)

func GetProductByCode(ctx context.Context, code string) (product.Product, error) {
	return products.GetByCode(ctx, code)
}
