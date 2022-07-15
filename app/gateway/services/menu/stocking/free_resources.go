package stocking

import (
	"comies/app/sdk/fault"
	"comies/app/sdk/types"
	"context"
)

func (s service) FreeResources(ctx context.Context, reservationID types.ID) error {
	err := s.stocks.FreeResources(ctx, reservationID)
	if err != nil {
		return fault.Wrap(err)
	}

	return nil
}
