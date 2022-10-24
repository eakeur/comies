package menu

import (
	"comies/app/core/menu/price"
	"comies/app/core/types"
	"context"
	"time"
)

func (w jobs) UpdateProductPrice(ctx context.Context, productID types.ID, val types.Currency) error {
	save, err := price.Price{TargetID: productID}.
		WithID(w.createID()).
		WithDate(time.Now()).
		WithValue(val).
		Validate()

	if err != nil {
		return err
	}

	return w.prices.Create(ctx, save)
}
