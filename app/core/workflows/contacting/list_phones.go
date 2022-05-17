package contacting

import (
	"context"
	"gomies/app/core/entities/phone"
	"gomies/app/sdk/fault"
	"gomies/app/sdk/types"
)

func (w workflow) ListPhones(ctx context.Context, targetID types.ID) ([]phone.Phone, error) {

	if targetID.Empty() {
		return nil, fault.Wrap(fault.ErrMissingID)
	}

	phones, err := w.phones.List(ctx, targetID)
	if err != nil {
		return nil, fault.Wrap(err)
	}

	return phones, nil

}
