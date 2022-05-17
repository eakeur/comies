package crew

import (
	"context"
	crew2 "gomies/app/core/entities/crew"
	"gomies/app/sdk/fault"
)

func (w workflow) ListMembers(ctx context.Context, filter crew2.Filter) ([]crew2.Member, int, error) {

	ct, count, err := w.crew.ListMembers(ctx, filter)
	if err != nil {
		return []crew2.Member{}, 0, fault.Wrap(err)
	}

	return ct, count, nil
}
