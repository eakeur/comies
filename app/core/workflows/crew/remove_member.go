package crew

import (
	"context"
	"gomies/app/core/entities/iam/crew"
	"gomies/app/sdk/fault"
)

func (w workflow) RemoveMember(ctx context.Context, key crew.Key) error {
	const operation = "Workflows.Crew.RemoveProduct"

	err := w.crew.RemoveMember(ctx, key)
	if err != nil {
		return fault.Wrap(err, operation)
	}

	return nil
}
