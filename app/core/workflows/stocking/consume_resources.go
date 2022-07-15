package stocking

import (
	"comies/app/sdk/throw"
	"comies/app/sdk/types"
	"context"
)

func (w workflow) ConsumeResources(ctx context.Context, reservationID types.ID) error {

	err := w.movements.SetOutputType(ctx, reservationID)
	if err != nil {
		return throw.Error(err)
	}

	return nil
}
