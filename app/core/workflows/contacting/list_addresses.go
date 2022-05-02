package contacting

import (
	"context"
	"gomies/app/core/entities/contacting"
	"gomies/app/sdk/fault"
	"gomies/app/sdk/types"
)

func (w workflow) ListAddresses(ctx context.Context, targetID types.ID) ([]contacting.Address, error) {
	const operation = "Workflows.Contacting.ListAddresses"

	if targetID.Empty() {
		return []contacting.Address{}, fault.Wrap(contacting.ErrMissingResourceID, operation)
	}

	addresses, err := w.contacts.ListAddresses(ctx, targetID)
	if err != nil {
		return nil, fault.Wrap(err, operation)
	}

	return addresses, nil

}
