package crew

import (
	"context"
	crew2 "gomies/app/core/entities/crew"
	"gomies/app/sdk/fault"
)

func (w workflow) GetMember(ctx context.Context, key crew2.Key) (crew2.Member, error) {
	const operation = "Workflows.Crew.GetProducts"

	ct, err := w.crew.GetMember(ctx, key)
	if err != nil {
		return crew2.Member{}, fault.Wrap(err, operation)
	}

	return ct, nil
}
