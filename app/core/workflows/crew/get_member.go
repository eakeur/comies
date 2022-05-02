package crew

import (
	"context"
	"gomies/app/core/entities/iam/crew"
	"gomies/app/sdk/fault"
)

func (w workflow) GetMember(ctx context.Context, key crew.Key) (crew.Member, error) {
	const operation = "Workflows.Crew.GetProducts"

	ct, err := w.crew.GetMember(ctx, key)
	if err != nil {
		return crew.Member{}, fault.Wrap(err, operation)
	}

	return ct, nil
}
