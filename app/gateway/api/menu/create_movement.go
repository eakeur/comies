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
		ProductID: types.ID(in.ProductID),
		Type:      InternalMovementType(in.Type),
		Date:      in.Date.AsTime().UTC(),
		Quantity:  types.Quantity(in.Quantity),
		PaidValue: types.Currency(in.Value),
	})
	if err != nil {
		return nil, failures.HandleError(throw.Error(err))
	}

	return &menu.CreateMovementResponse{
		Id: int64(ing.ID),
	}, nil
}
