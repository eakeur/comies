package contacting

import (
	"context"
	"gomies/pkg/contacting/core/entities/contacting"
	"gomies/pkg/sdk/fault"
	"gomies/pkg/sdk/types"
)

func (w workflow) RemoveAddress(ctx context.Context, targetID types.UID, id types.UID) error {
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
