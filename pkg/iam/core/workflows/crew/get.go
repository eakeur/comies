package crew

import (
	"context"
	"gomies/pkg/iam/core/entities/crew"
	"gomies/pkg/sdk/fault"
)

func (w workflow) Get(ctx context.Context, key crew.Key) (crew.Member, error) {
	const operation = "Workflows.Crew.Get"

	ct, err := w.crew.Get(ctx, key)
	if err != nil {
		return crew.Member{}, fault.Wrap(err, operation)
	}

	return ct, nil
}
