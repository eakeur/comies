package crew

import (
	"context"
	"gomies/app/core/entities/crew"
	"gomies/app/sdk/fault"
)

func (w workflow) RemoveMember(ctx context.Context, key crew.Key) error {

	err := w.crew.RemoveMember(ctx, key)
	if err != nil {
		return fault.Wrap(err)
	}

	return nil
}
