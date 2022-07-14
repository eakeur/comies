package crew

import (
	"context"
	member "gomies/app/core/entities/member"
	"gomies/app/sdk/fault"
)

func (w workflow) List(ctx context.Context, filter member.Filter) ([]member.Member, error) {

	ct, err := w.crew.List(ctx, filter)
	if err != nil {
		return []member.Member{}, fault.Wrap(err)
	}

	return ct, nil
}
