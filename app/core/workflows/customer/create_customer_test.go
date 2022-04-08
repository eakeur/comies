package customer

import (
	"context"
	"github.com/stretchr/testify/assert"
	"gomies/app/core/entities/customer"
	"gomies/pkg/sdk/fault"
	"testing"
)

func TestWorkflow_CreateCustomer(t *testing.T) {
	t.Parallel()

	type (
		args struct {
			customer customer.Customer
		}

		fields struct {
			customers *customer.ActionsMock
		}

		test struct {
			name    string
			args    args
			want    customer.Customer
			wantErr error
			fields  fields
		}
	)

	for _, c := range []test{
		{
			name: "should return customer created",
			args: args{
				customer: customer.Customer{
					Name:              "Aubergine",
					PhoneDigest:       "991912222;43430987",
					AddressCodeDigest: "03911200;",
				},
			},
			want: customer.Customer{
				Name:              "Aubergine",
				PhoneDigest:       "991912222;43430987",
				AddressCodeDigest: "03911200;",
			},
			fields: fields{
				customers: &customer.ActionsMock{
					CreateCustomerFunc: func(ctx context.Context, c customer.Customer) (customer.Customer, error) {
						return c, nil
					},
				},
			},
		},
		{
			name: "should fail with ErrNotFound",
			args: args{
				customer: customer.Customer{
					Name:              "Aubergine",
					PhoneDigest:       "991912222;43430987",
					AddressCodeDigest: "03911200;",
				},
			},
			wantErr: fault.ErrAlreadyExists,
			fields: fields{
				customers: &customer.ActionsMock{
					CreateCustomerFunc: func(ctx context.Context, c customer.Customer) (customer.Customer, error) {
						return customer.Customer{}, fault.ErrAlreadyExists
					},
				},
			},
		},
	} {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			wf := workflow{customers: c.fields.customers}
			got, gotErr := wf.CreateCustomer(context.Background(), c.args.customer)
			assert.ErrorIs(t, gotErr, c.wantErr)
			assert.Equal(t, c.want, got)
		})
	}

}
