package contacting

import (
	"context"
	"gomies/app/core/entities/phone"
	"gomies/app/sdk/fault"
	"gomies/app/sdk/types"
)

func (w workflow) GetPhoneByID(ctx context.Context, phoneID types.ID) (phone.Phone, error) {

	if phoneID.Empty() {
		return phone.Phone{}, fault.Wrap(fault.ErrMissingID)
	}

	p, err := w.phones.GetByID(ctx, phoneID)
	if err != nil {
		return phone.Phone{}, fault.Wrap(err)
	}

	return p, nil
}
