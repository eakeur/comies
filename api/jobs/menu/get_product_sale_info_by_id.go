package menu

import (
	"comies/core/types"
	"context"
)

func (w jobs) GetProductLatestPriceByID(ctx context.Context, id types.ID) (types.Currency, error) {
	return w.prices.GetLatestValue(ctx, id)
}
