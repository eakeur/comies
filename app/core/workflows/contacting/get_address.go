package contacting

import (
	"context"
	"gomies/app/core/entities/address"
	"gomies/app/sdk/fault"
	"gomies/app/sdk/types"
)

func (w workflow) GetAddressByID(ctx context.Context, addressID types.ID) (address.Address, error) {

	if addressID.Empty() {
		return address.Address{}, fault.Wrap(fault.ErrMissingID)
	}

	addr, err := w.addresses.GetByID(ctx, addressID)
	if err != nil {
		return address.Address{}, fault.Wrap(err)
	}

	return addr, nil
}
