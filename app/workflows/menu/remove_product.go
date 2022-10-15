package menu

import (
	"comies/app/core/id"
	"comies/app/data/products"
	"context"
)

func RemoveProduct(ctx context.Context, productID id.ID) error {
	if err := id.ValidateID(productID); err != nil {
		return err
	}
	
	err := products.Remove(ctx, productID)
	if err != nil {
		return err
	}

	return err

}
