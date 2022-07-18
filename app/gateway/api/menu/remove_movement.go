package menu

import (
	"comies/app/gateway/api/gen/menu"
	"comies/app/sdk/throw"
	"comies/app/sdk/types"
	"context"
)

func (s service) RemoveMovement(ctx context.Context, in *menu.RemoveMovementRequest) (*menu.Empty, error) {
	err := s.menu.RemoveMovement(ctx, types.ID(in.Id))
	if err != nil {
		return nil, failures.HandleError(throw.Error(err))
	}

	return &menu.Empty{}, nil
}
