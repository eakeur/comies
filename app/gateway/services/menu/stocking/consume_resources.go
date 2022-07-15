package stocking

import (
	"comies/app/sdk/throw"
	"comies/app/sdk/types"
	"context"
)

func (s service) ConsumeResources(ctx context.Context, reservationID types.ID) error {
	err := s.stocks.ConsumeResources(ctx, reservationID)
	if err != nil {
		return throw.Error(err)
	}

	return nil
}
