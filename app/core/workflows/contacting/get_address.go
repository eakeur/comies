package contacting

import (
	"context"
	"gomies/app/core/entities/contacting"
	"gomies/pkg/sdk/fault"
	"gomies/pkg/sdk/types"
)

func (w workflow) GetAddress(ctx context.Context, targetID types.UID, id types.UID) (contacting.Address, error) {
	const operation = "Workflows.Contacting.GetAddress"

	if targetID.Empty() || id.Empty() {
		return contacting.Address{}, fault.Wrap(contacting.ErrMissingResourceID, operation)
	}

	address, err := w.contacts.GetAddress(ctx, targetID, id)
	if err != nil {
		return contacting.Address{}, fault.Wrap(err, operation)
	}

	return address, nil
}
