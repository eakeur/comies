package customer

import (
	"context"
	"gomies/pkg/sdk/fault"
	"gomies/pkg/sdk/types"
)

func (w workflow) RemoveCustomer(ctx context.Context, uid types.UID) error {
	const operation = "Workflows.Customer.RemoveCustomer"

	if uid.Empty() {
		return fault.Wrap(fault.ErrMissingUID, operation)
	}

	err := w.customers.RemoveCustomer(ctx, uid)
	if err != nil {
		return fault.Wrap(err, operation)
	}

	return nil
}
