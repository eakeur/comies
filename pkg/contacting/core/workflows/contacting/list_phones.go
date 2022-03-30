package contacting

import (
	"context"
	"gomies/pkg/contacting/core/entities/contacting"
	"gomies/pkg/sdk/fault"
	"gomies/pkg/sdk/types"
)

func (w workflow) ListPhones(ctx context.Context, targetID types.UID) ([]contacting.Phone, error) {
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
