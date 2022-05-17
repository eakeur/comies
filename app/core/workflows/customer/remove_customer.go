package customer

import (
	"context"
	"gomies/app/sdk/fault"
	"gomies/app/sdk/types"
)

func (w workflow) RemoveCustomer(ctx context.Context, id types.ID) error {

	if id.Empty() {
		return fault.Wrap(fault.ErrMissingID)
	}

	err := w.customers.RemoveCustomer(ctx, id)
	if err != nil {
		return fault.Wrap(err)
	}

	return nil
}
