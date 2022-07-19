package menu

import (
	"comies/app/gateway/api/gen/menu"
	"comies/app/sdk/throw"
	"comies/app/sdk/types"
	"context"
)

func (s service) UpdateProduct(ctx context.Context, in *menu.UpdateProductRequest) (*menu.Empty, error) {
	prd := InternalProduct(in.Product)
	prd.ID = types.ID(in.Product.Id)
	err := s.menu.UpdateProduct(ctx, prd)
	if err != nil {
		return nil, failures.HandleError(throw.Error(err))
	}

	return &menu.Empty{}, nil
}
