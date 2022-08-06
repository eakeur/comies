package menu

import (
	"comies/app/core/throw"
	"comies/app/core/types"
	"context"
)

func (w workflow) GetProductNameByID(ctx context.Context, id types.ID) (string, error) {
	prod, err := w.products.GetNameByID(ctx, id)
	if err != nil {
		return "", throw.Error(err)
	}
	return prod, nil
}
