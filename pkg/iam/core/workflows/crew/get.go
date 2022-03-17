package crew

import (
	"context"
	"gomies/pkg/iam/core/entities/crew"
	"gomies/pkg/sdk/fault"
	"gomies/pkg/sdk/session"
)

func (w workflow) Get(ctx context.Context, key crew.Key) (crew.Member, error) {
	const operation = "Workflows.Crew.Get"

	_, err := session.FromContext(ctx, operation)
	if err != nil {
		return crew.Member{}, fault.Wrap(err, operation)
	}

	ct, err := w.crew.Get(ctx, key)
	if err != nil {
		return crew.Member{}, fault.Wrap(err, operation)
	}

	return ct, nil
}
