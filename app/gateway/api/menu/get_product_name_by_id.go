package menu

import (
	"comies/app/gateway/api/gen/menu"
	"comies/app/sdk/throw"
	"comies/app/sdk/types"
	"context"
)

func (s service) GetProductNameByID(ctx context.Context, in *menu.GetProductNameByIDRequest) (*menu.GetProductNameByIDResponse, error) {
	prd, err := s.menu.GetProductNameByID(ctx, types.ID(in.Id))
	if err != nil {
		return nil, throw.Error(err)
	}

	return &menu.GetProductNameByIDResponse{
		Name: prd,
	}, nil
}
