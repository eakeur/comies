package crew

import (
	"context"
	"gomies/app/core/entities/member"
	"gomies/app/sdk/fault"
)

func (w workflow) Remove(ctx context.Context, key member.Key) error {

	err := w.crew.Remove(ctx, key)
	if err != nil {
		return fault.Wrap(err)
	}

	return nil
}
