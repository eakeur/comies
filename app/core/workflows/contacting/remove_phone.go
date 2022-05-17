package contacting

import (
	"context"
	"gomies/app/sdk/fault"
	"gomies/app/sdk/types"
)

func (w workflow) RemovePhone(ctx context.Context, id types.ID) error {

	if id.Empty() {
		return fault.Wrap(fault.ErrMissingID)
	}

	err := w.phones.Remove(ctx, id)
	if err != nil {
		return fault.Wrap(err)
	}

	return nil
}
