package stock

import (
	"comies/app/sdk/fault"
	"comies/app/sdk/types"
	"context"
)

func (s service) ConsumeResources(ctx context.Context, reservationID types.ID) error {
	err := s.stocks.ConsumeResources(ctx, reservationID)
	if err != nil {
		return fault.Wrap(err)
	}

	return nil
}
