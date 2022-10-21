package menu

import (
	"comies/app/core/types"
	"comies/app/data/products"
	"context"
)

func RemoveProduct(ctx context.Context, productID types.ID) error {
	if err := types.ValidateID(productID); err != nil {
		return err
	}

	return products.Remove(ctx, productID)
}
