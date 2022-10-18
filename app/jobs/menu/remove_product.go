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

	err := products.Remove(ctx, productID)
	if err != nil {
		return err
	}

	return err

}
