package customer

import (
	"context"
	"gomies/app/core/entities/customer"
	"gomies/app/sdk/fault"
)

func (w workflow) ListCustomers(ctx context.Context, filter customer.Filter) ([]customer.Customer, int, error) {

	ct, count, err := w.customers.ListCustomers(ctx, filter)
	if err != nil {
		return []customer.Customer{}, 0, fault.Wrap(err)
	}

	return ct, count, nil
}
