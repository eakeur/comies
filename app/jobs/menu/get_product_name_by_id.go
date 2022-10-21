package menu

import (
	"comies/app/core/types"
	"comies/app/data/products"
	"context"
)

func GetProductNameByID(ctx context.Context, id types.ID) (string, error) {
	if err := types.ValidateID(id); err != nil {
		return "", err
	}

	return products.GetNameByID(ctx, id)
}
