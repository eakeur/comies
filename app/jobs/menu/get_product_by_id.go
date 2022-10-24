package menu

import (
	"comies/app/core/menu/product"
	"comies/app/core/types"
	"context"
)

func (w jobs) GetProductByID(ctx context.Context, id types.ID) (product.Product, error) {

	if err := id.Validate(); err != nil {
		return product.Product{}, err
	}

	return w.products.GetByID(ctx, id)
}
