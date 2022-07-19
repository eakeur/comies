package menu

import (
	"comies/app/gateway/api/gen/menu"
	"comies/app/sdk/throw"
	"comies/app/sdk/types"
	"context"
)

func (s service) RemoveIngredient(ctx context.Context, in *menu.RemoveIngredientRequest) (*menu.Empty, error) {
	err := s.menu.RemoveIngredient(ctx, types.ID(in.Id))
	if err != nil {
		return nil, failures.HandleError(throw.Error(err))
	}

	return &menu.Empty{}, nil
}
