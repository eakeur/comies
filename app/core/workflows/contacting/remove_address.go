package contacting

import (
	"context"
	"gomies/app/core/entities/contacting"
	"gomies/app/sdk/fault"
	"gomies/app/sdk/types"
)

func (w workflow) RemoveAddress(ctx context.Context, targetID types.ID, id types.ID) error {
	const operation = "Workflows.Contacting.RemoveAddress"

	if targetID.Empty() {
		return fault.Wrap(contacting.ErrMissingResourceID, operation)
	}

	err := w.contacts.RemoveAddresses(ctx, targetID, id)
	if err != nil {
		return fault.Wrap(err, operation)
	}

	return nil
}
