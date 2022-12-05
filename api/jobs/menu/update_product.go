package menu

import (
	"comies/core/menu/product"
	"context"
)

func (w jobs) UpdateProduct(ctx context.Context, p product.Product) error {
	if err := p.ID.Validate(); err != nil {
		return err
	}

	if _, err := p.Validate(); err != nil {
		return err
	}

	return w.products.Update(ctx, p)
}
