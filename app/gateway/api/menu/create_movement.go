package menu

import (
	"comies/app/core/entities/movement"
	"comies/app/gateway/api/gen/menu"
	"comies/app/sdk/throw"
	"comies/app/sdk/types"
	"context"
)

func (s service) CreateMovement(ctx context.Context, in *menu.CreateMovementRequest) (*menu.CreateMovementResponse, error) {
	ing, err := s.menu.CreateMovement(ctx, movement.Movement{
		ID:        0,
		ProductID: types.ID(in.Movement.ProductID),
		Type:      movement.Type(in.Movement.Type),
		Date:      in.Movement.Date.AsTime().UTC(),
		Quantity:  types.Quantity(in.Movement.Quantity),
		PaidValue: types.Currency(in.Movement.Value),
	})
	if err != nil {
		return nil, failures.HandleError(throw.Error(err))
	}

	return &menu.CreateMovementResponse{
		Id: int64(ing.ID),
	}, nil
}
