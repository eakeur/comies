package customer

import (
	"context"
	"gomies/app/core/entities/customer"
	"gomies/pkg/sdk/fault"
	"gomies/pkg/sdk/types"
)

func (w workflow) GetCustomer(ctx context.Context, uid types.UID) (customer.Customer, error) {
	const operation = "Workflows.Customer.GetCustomer"

	if uid.Empty() {
		return customer.Customer{}, fault.Wrap(fault.ErrMissingUID, operation)
	}

	c, err := w.customers.GetCustomer(ctx, uid)
	if err != nil {
		return customer.Customer{}, fault.Wrap(err, operation)
	}

	return c, nil
}
