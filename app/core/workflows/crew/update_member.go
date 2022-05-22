package crew

import (
	"context"
	"gomies/app/core/entities/member"
	"gomies/app/sdk/fault"
)

func (w workflow) Update(ctx context.Context, m member.Member) error {

	if err := m.Validate(); err != nil {
		return fault.Wrap(err)
	}

	err := w.crew.Update(ctx, m)
	if err != nil {
		return fault.Wrap(err)
	}

	return nil
}
