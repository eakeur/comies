package menu

import (
	"comies/app/core/types"
	"context"
)

type SaleInfo struct {
	Unit    types.UnitType
	Minimum types.Quantity
	Price   types.Currency
}

func (w jobs) GetProductSaleInfoByID(ctx context.Context, id types.ID) (SaleInfo, error) {
	p, err := w.products.GetByID(ctx, id)
	if err != nil {
		return SaleInfo{}, err
	}

	price, err := w.prices.GetLatestValue(ctx, id)
	if err != nil {
		return SaleInfo{}, err
	}

	return SaleInfo{
		Unit:    p.SaleUnit,
		Minimum: p.MinimumSale,
		Price:   price,
	}, nil
}
