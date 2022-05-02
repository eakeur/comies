package contacting

import (
	"context"
	"gomies/app/core/entities/contacting"
	"gomies/app/sdk/fault"
	"gomies/app/sdk/types"
)

func (w workflow) GetAddress(ctx context.Context, targetID types.ID, id types.ID) (contacting.Address, error) {
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
