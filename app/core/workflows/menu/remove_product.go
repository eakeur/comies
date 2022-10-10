package menu

import (
	"comies/app/core/id"
	"comies/app/data/products"
	"context"
)

func RemoveProduct(ctx context.Context, id id.ID) error {

	err := products.Remove(ctx, id)
	if err != nil {
		return err
	}

	return err

}
