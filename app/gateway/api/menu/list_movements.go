package menu

import (
	"comies/app/core/entities/movement"
	"comies/app/gateway/api/gen/menu"
	"comies/app/sdk/throw"
	"comies/app/sdk/types"
	"context"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s service) ListMovements(ctx context.Context, in *menu.ListMovementsRequest) (*menu.ListMovementsResponse, error) {
	list, err := s.menu.ListMovements(ctx, movement.Filter{
		ProductID:   types.ID(in.ProductID),
		InitialDate: in.Start.AsTime(),
		FinalDate:   in.End.AsTime(),
	})
	if err != nil {
		return nil, throw.Error(err)
	}

	var movements []*menu.Movement
	for _, p := range list {
		movements = append(movements, &menu.Movement{
			Id:        int64(p.ID),
			ProductID: int64(p.ProductID),
			Type:      menu.MovementType(p.Type),
			Date:      timestamppb.New(p.Date),
			Quantity:  int64(p.Quantity),
			Value:     int64(p.PaidValue),
		})
	}

	return &menu.ListMovementsResponse{
		Movements: movements,
	}, nil
}
