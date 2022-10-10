package menu

import (
	"comies/app/core/id"
	"comies/app/data/products"
	"context"
)

func GetProductNameByID(ctx context.Context, id id.ID) (string, error) {
	prod, err := products.GetNameByID(ctx, id)
	if err != nil {
		return "", err
	}
	return prod, nil
}
