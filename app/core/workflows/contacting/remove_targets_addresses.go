package contacting

import (
	"context"
	"gomies/app/sdk/fault"
	"gomies/app/sdk/types"
)

func (w workflow) RemoveTargetsAddresses(ctx context.Context, targetID types.ID) error {

	if targetID.Empty() {
		return fault.Wrap(fault.ErrMissingID)
	}

	err := w.addresses.RemoveAllByTarget(ctx, targetID)
	if err != nil {
		return fault.Wrap(err)
	}

	return nil
}
