package contacting

import (
	"context"
	"gomies/app/core/entities/contacting"
	"gomies/app/sdk/fault"
	"gomies/app/sdk/types"
)

func (w workflow) GetPhone(ctx context.Context, targetID types.ID, id types.ID) (contacting.Phone, error) {
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
