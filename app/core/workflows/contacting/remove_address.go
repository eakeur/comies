package contacting

import (
	"context"
	"gomies/app/sdk/fault"
	"gomies/app/sdk/types"
)

func (w workflow) RemoveAddress(ctx context.Context, id types.ID) error {

	if id.Empty() {
		return fault.Wrap(fault.ErrMissingID)
	}

	err := w.addresses.Remove(ctx, id)
	if err != nil {
		return fault.Wrap(err)
	}

	return nil
}
