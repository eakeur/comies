package menu

import (
	"comies/app/sdk/throw"
	"comies/app/sdk/types"
	"context"
)

func (s service) UpdateResources(ctx context.Context, reservationID types.ID, consume bool) error {
	err := s.menu.UpdateReservation(ctx, reservationID, consume)
	if err != nil {
		return throw.Error(err)
	}

	return nil
}
