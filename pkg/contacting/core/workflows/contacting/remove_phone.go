package contacting

import (
	"context"
	"gomies/pkg/contacting/core/entities/contacting"
	"gomies/pkg/sdk/fault"
	"gomies/pkg/sdk/types"
)

func (w workflow) RemovePhone(ctx context.Context, targetID types.UID, id types.UID) error {
	const operation = "Workflows.Contacting.RemovePhones"

	if targetID.Empty() {
		return fault.Wrap(contacting.ErrMissingResourceID, operation)
	}

	err := w.contacts.RemovePhones(ctx, targetID, id)
	if err != nil {
		return fault.Wrap(err, operation)
	}

	return nil
}
