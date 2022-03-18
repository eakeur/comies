package crew

import (
	"context"
	"gomies/pkg/iam/core/entities/crew"
	"gomies/pkg/sdk/fault"
)

func (w workflow) Remove(ctx context.Context, key crew.Key) error {
	const operation = "Workflows.Crew.Remove"

	err := w.crew.Remove(ctx, key)
	if err != nil {
		return fault.Wrap(err, operation)
	}

	return nil
}
