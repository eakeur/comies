package crew

import (
	"context"
	"gomies/app/core/entities/iam/crew"
	"gomies/pkg/sdk/fault"
)

func (w workflow) List(ctx context.Context, filter crew.Filter) ([]crew.Member, error) {
	const operation = "Workflows.Crew.List"
	ct, err := w.crew.List(ctx, filter)
	if err != nil {
		return []crew.Member{}, fault.Wrap(err, operation)
	}

	return ct, nil
}
