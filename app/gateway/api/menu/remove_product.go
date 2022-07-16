package menu

import (
	"comies/app/gateway/api/gen/menu"
	"comies/app/sdk/throw"
	"comies/app/sdk/types"
	"context"
)

func (s service) RemoveProduct(ctx context.Context, in *menu.RemoveProductRequest) (*menu.Empty, error) {
	err := s.menu.RemoveProduct(ctx, types.ID(in.Id))
	if err != nil {
		return nil, throw.Error(err)
	}

	return &menu.Empty{}, nil
}
