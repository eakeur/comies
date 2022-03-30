package contacting

import (
	"context"
	"gomies/pkg/contacting/core/entities/contacting"
	"gomies/pkg/sdk/fault"
	"gomies/pkg/sdk/types"
)

func (w workflow) ListAddresses(ctx context.Context, targetID types.UID) ([]contacting.Address, error) {
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
