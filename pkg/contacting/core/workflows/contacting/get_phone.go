package contacting

import (
	"context"
	"gomies/pkg/contacting/core/entities/contacting"
	"gomies/pkg/sdk/fault"
	"gomies/pkg/sdk/types"
)

func (w workflow) GetPhone(ctx context.Context, targetID types.UID, id types.UID) (contacting.Phone, error) {
	const operation = "Workflows.Contacting.GetPhone"

	if targetID.Empty() || id.Empty() {
		return contacting.Phone{}, fault.Wrap(contacting.ErrMissingResourceID, operation)
	}

	phone, err := w.contacts.GetPhone(ctx, targetID, id)
	if err != nil {
		return contacting.Phone{}, fault.Wrap(err, operation)
	}

	return phone, nil
}
