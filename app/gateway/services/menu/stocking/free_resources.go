package stocking

import (
	"comies/app/sdk/throw"
	"comies/app/sdk/types"
	"context"
)

func (s service) FreeResources(ctx context.Context, reservationID types.ID) error {
	err := s.stocks.FreeResources(ctx, reservationID)
	if err != nil {
		return throw.Error(err)
	}

	return nil
}
