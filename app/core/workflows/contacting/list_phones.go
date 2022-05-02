package contacting

import (
	"context"
	"gomies/app/core/entities/contacting"
	"gomies/app/sdk/fault"
	"gomies/app/sdk/types"
)

func (w workflow) ListPhones(ctx context.Context, targetID types.ID) ([]contacting.Phone, error) {
	const operation = "Workflows.Contacting.ListPhones"

	if targetID.Empty() {
		return []contacting.Phone{}, fault.Wrap(contacting.ErrMissingResourceID, operation)
	}

	addresses, err := w.contacts.ListPhones(ctx, targetID)
	if err != nil {
		return nil, fault.Wrap(err, operation)
	}

	return addresses, nil

}
