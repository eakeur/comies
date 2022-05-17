package crew

import (
	"context"
	"gomies/app/core/entities/crew"
	"gomies/app/sdk/fault"
)

func (w workflow) CreateMember(ctx context.Context, input crew.Member) (crew.Member, error) {

	if err := input.Validate(); err != nil {
		return crew.Member{}, fault.Wrap(err)
	}

	input, err := w.crew.CreateMember(ctx, input)
	if err != nil {
		return crew.Member{}, fault.Wrap(err)
	}

	return input, nil
}
