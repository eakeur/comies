package crew

import (
	"context"
	"gomies/app/core/entities/crew"
	"gomies/app/sdk/fault"
)

func (w workflow) UpdateMember(ctx context.Context, m crew.Member) error {

	if err := m.Validate(); err != nil {
		return fault.Wrap(err)
	}

	err := w.crew.UpdateMember(ctx, m)
	if err != nil {
		return fault.Wrap(err)
	}

	return nil
}
