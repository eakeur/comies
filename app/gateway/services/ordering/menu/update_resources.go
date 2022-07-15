package menu

import (
	"comies/app/sdk/fault"
	"comies/app/sdk/types"
	"context"
)

func (s service) UpdateResources(ctx context.Context, reservationID types.ID, consume bool) error {
	err := s.menu.UpdateReservation(ctx, reservationID, consume)
	if err != nil {
		return fault.Wrap(err)
	}

	return nil
}
