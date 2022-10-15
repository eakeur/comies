package menu

import (
	"comies/app/core/id"
	"comies/app/core/menu"
	"comies/app/data/ids"
	"comies/app/data/products"
	"context"
)

func SaveProduct(ctx context.Context, p menu.Product) (menu.Product, error) {
	if err := menu.ValidateProduct(p); err != nil {
		return menu.Product{}, err
	}

	if err := id.ValidateID(p.ID); err == nil {
		return p, products.Update(ctx, p)
	}

	p.ID = ids.Create()
	return p, products.Create(ctx, p)
}
