package crew

import (
	"context"
	"gomies/pkg/iam/core/entities/crew"
	"gomies/pkg/sdk/fault"
	"gomies/pkg/sdk/types"
)

func (w workflow) Save(ctx context.Context, input crew.Member, flag ...types.WritingFlag) (crew.Member, error) {
	const operation = "Workflows.Crew.Save"

	if err := input.Validate(); err != nil {
		return crew.Member{}, fault.Wrap(err, operation)
	}

	input, err := w.crew.Save(ctx, input, flag...)
	if err != nil {
		return crew.Member{}, fault.Wrap(err, operation)
	}

	return input, nil
}
