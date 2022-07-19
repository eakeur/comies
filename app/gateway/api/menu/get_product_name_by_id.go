package menu

import (
	"comies/app/gateway/api/gen/menu"
	"comies/app/sdk/throw"
	"comies/app/sdk/types"
	"context"
)

func (s service) GetProductNameById(ctx context.Context, in *menu.GetProductNameByIdRequest) (*menu.GetProductNameByIdResponse, error) {
	prd, err := s.menu.GetProductNameByID(ctx, types.ID(in.Id))
	if err != nil {
		return nil, failures.HandleError(throw.Error(err))
	}

	return &menu.GetProductNameByIdResponse{
		Name: prd,
	}, nil
}
