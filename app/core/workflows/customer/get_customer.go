package customer

import (
	"context"
	"gomies/app/core/entities/customer"
	"gomies/app/sdk/fault"
	"gomies/app/sdk/types"
)

func (w workflow) GetCustomer(ctx context.Context, id types.ID) (customer.Customer, error) {

	if id.Empty() {
		return customer.Customer{}, fault.Wrap(fault.ErrMissingID)
	}

	c, err := w.customers.GetCustomer(ctx, id)
	if err != nil {
		return customer.Customer{}, fault.Wrap(err)
	}

	return c, nil
}
