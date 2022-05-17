package contacting

import (
	"context"
	"gomies/app/sdk/fault"
	"gomies/app/sdk/types"
)

func (w workflow) RemoveTargetsPhones(ctx context.Context, targetID types.ID) error {

	if targetID.Empty() {
		return fault.Wrap(fault.ErrMissingID)
	}

	err := w.phones.RemoveAllByTarget(ctx, targetID)
	if err != nil {
		return fault.Wrap(err)
	}

	return nil
}
