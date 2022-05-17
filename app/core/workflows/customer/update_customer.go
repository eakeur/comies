package customer

import (
	"context"
	"gomies/app/core/entities/customer"
	"gomies/app/sdk/fault"
)

func (w workflow) UpdateCustomer(ctx context.Context, op customer.Customer) error {

	if err := op.Validate(); err != nil {
		return fault.Wrap(err)
	}

	err := w.customers.UpdateCustomer(ctx, op)
	if err != nil {
		return fault.Wrap(err)
	}

	return nil
}
