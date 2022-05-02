package crew

import (
	"context"
	"gomies/app/core/entities/iam/crew"
	"gomies/app/sdk/fault"
)

func (w workflow) ListMembers(ctx context.Context, filter crew.Filter) ([]crew.Member, int, error) {
	const operation = "Workflows.Crew.ListProducts"
	ct, count, err := w.crew.ListMembers(ctx, filter)
	if err != nil {
		return []crew.Member{}, 0, fault.Wrap(err, operation)
	}

	return ct, count, nil
}
