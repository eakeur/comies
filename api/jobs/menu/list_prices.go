package menu

import (
	"comies/core/menu/price"
	"comies/core/types"
	"context"
)

func (w jobs) ListPrices(ctx context.Context, productID types.ID) ([]price.Price, error) {
	if err := productID.Validate(); err != nil {
		return nil, err
	}

	return w.prices.List(ctx, productID)
}
