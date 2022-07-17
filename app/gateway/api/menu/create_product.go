package menu

import (
	"comies/app/gateway/api/gen/menu"
	"comies/app/sdk/throw"
	"context"
)

func (s service) CreateProduct(ctx context.Context, in *menu.CreateProductRequest) (*menu.CreateProductResponse, error) {

	prd, err := s.menu.CreateProduct(ctx, InternalProduct(in))
	if err != nil {
		return nil, failures.HandleError(throw.Error(err))
	}

	return &menu.CreateProductResponse{
		Id: int64(prd.ID),
	}, nil
}
