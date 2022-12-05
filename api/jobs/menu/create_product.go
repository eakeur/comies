package menu

import (
	"comies/core/menu/price"
	"comies/core/menu/product"
	"comies/core/types"
	"context"
	"time"
)

type ProductCreation struct {
	product.Product
	SalePrice types.Currency
}

func (w jobs) CreateProduct(ctx context.Context, p ProductCreation) (types.ID, error) {
	save, err := p.WithID(w.createID()).Validate()
	if err != nil {
		return 0, err
	}

	err = w.products.Create(ctx, save)
	if err != nil {
		return 0, err
	}

	pr, err := price.Price{TargetID: save.ID}.
		WithID(w.createID()).
		WithDate(time.Now()).
		WithValue(p.SalePrice).
		Validate()
	if err != nil {
		return 0, err
	}

	return 0, w.prices.Create(ctx, pr)
}
