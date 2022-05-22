package crew

import (
	"context"
	"gomies/app/core/entities/member"
	"gomies/app/sdk/fault"
)

func (w workflow) Create(ctx context.Context, input member.Member) (member.Member, error) {

	if err := input.Validate(); err != nil {
		return member.Member{}, fault.Wrap(err)
	}

	input, err := w.crew.Create(ctx, input)
	if err != nil {
		return member.Member{}, fault.Wrap(err)
	}

	return input, nil
}
