package customer

import (
	"context"
	"gomies/app/sdk/fault"
	"gomies/app/sdk/types"
)

func (w workflow) RemoveCustomer(ctx context.Context, id types.ID) error {
	const operation = "Workflows.Customer.RemoveCustomer"

	if id.Empty() {
		return fault.Wrap(fault.ErrMissingUID, operation)
	}

	err := w.customers.RemoveCustomer(ctx, id)
	if err != nil {
		return fault.Wrap(err, operation)
	}

	return nil
}
