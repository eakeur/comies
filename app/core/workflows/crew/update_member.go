package crew

import (
	"context"
	"gomies/app/core/entities/crew"
	"gomies/app/sdk/fault"
)

func (w workflow) UpdateMember(ctx context.Context, m crew.Member) error {
	const operation = "Workflows.Product.UpdateMember"

	if err := m.Validate(); err != nil {
		return fault.Wrap(err, operation)
	}

	err := w.crew.UpdateMember(ctx, m)
	if err != nil {
		return fault.Wrap(err, operation)
	}

	return nil
}
