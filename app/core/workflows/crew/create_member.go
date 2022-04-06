package crew

import (
	"context"
	"gomies/app/core/entities/iam/crew"
	"gomies/pkg/sdk/fault"
)

func (w workflow) CreateMember(ctx context.Context, input crew.Member) (crew.Member, error) {
	const operation = "Workflows.Crew.CreateMember"

	if err := input.Validate(); err != nil {
		return crew.Member{}, fault.Wrap(err, operation)
	}

	input, err := w.crew.CreateMember(ctx, input)
	if err != nil {
		return crew.Member{}, fault.Wrap(err, operation)
	}

	return input, nil
}
