package menu

import (
	"comies/app/core/types"
	"context"
)

func (w jobs) GetProductNameByID(ctx context.Context, id types.ID) (string, error) {
	if err := id.Validate(); err != nil {
		return "", err
	}

	return w.products.GetNameByID(ctx, id)
}
