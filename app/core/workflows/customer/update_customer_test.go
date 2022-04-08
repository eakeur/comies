package customer

import (
	"context"
	"gomies/app/core/entities/customer"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWorkflow_UpdateMember(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	type (
		args struct {
			customer customer.Customer
		}

		opts struct {
			customers *customer.ActionsMock
		}

		test struct {
			name    string
			args    args
			opts    opts
			wantErr error
		}
	)

	cases := []test{
		{
			name: "should return nil for successful update",
			args: args{
				customer: customer.Customer{},
			},
			opts: opts{
				customers: &customer.ActionsMock{
					UpdateCustomerFunc: func(ctx context.Context, c customer.Customer) error {
						return nil
					},
				},
			},
		},
	}

	for _, c := range cases {
		c := c

		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			wf := workflow{
				customers: c.opts.customers,
			}
			err := wf.UpdateCustomer(ctx, c.args.customer)
			assert.ErrorIs(t, err, c.wantErr)
		})
	}

}
