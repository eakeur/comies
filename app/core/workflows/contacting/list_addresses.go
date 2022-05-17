package contacting

import (
	"context"
	"gomies/app/core/entities/address"
	"gomies/app/sdk/fault"
	"gomies/app/sdk/types"
)

func (w workflow) ListAddresses(ctx context.Context, targetID types.ID) ([]address.Address, error) {

	if targetID.Empty() {
		return nil, fault.Wrap(fault.ErrMissingID)
	}

	addresses, err := w.addresses.List(ctx, targetID)
	if err != nil {
		return nil, fault.Wrap(err)
	}

	return addresses, nil

}
