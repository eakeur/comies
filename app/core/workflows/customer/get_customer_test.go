package customer

import (
	"context"
	"github.com/stretchr/testify/assert"
	"gomies/app/core/entities/customer"
	"gomies/pkg/sdk/fault"
	"gomies/pkg/sdk/types"
	"testing"
)

func TestWorkflow_GetCustomer(t *testing.T) {
	t.Parallel()

	fakeID := types.NewUID()

	type (
		fields struct {
			customers *customer.ActionsMock
		}

		args struct {
			uid types.UID
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
			name: "should return customer found with id",
			args: args{
				uid: fakeID,
			},
			want: customer.Customer{
				Name:              "Granger",
				PhoneDigest:       "99558874;",
				AddressCodeDigest: "03322550;",
			},
			fields: fields{
				customers: &customer.ActionsMock{
					GetCustomerFunc: func(ctx context.Context, uid types.UID) (customer.Customer, error) {
						return customer.Customer{
							Name:              "Granger",
							PhoneDigest:       "99558874;",
							AddressCodeDigest: "03322550;",
						}, nil
					},
				},
			},
		},
		{
			name: "should fail with ErrNotFound",
			args: args{
				uid: fakeID,
			},
			wantErr: fault.ErrNotFound,
			fields: fields{
				customers: &customer.ActionsMock{
					GetCustomerFunc: func(ctx context.Context, uid types.UID) (customer.Customer, error) {
						return customer.Customer{}, fault.ErrNotFound
					},
				},
			},
		},
	} {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			wf := workflow{customers: c.fields.customers}
			got, gotErr := wf.GetCustomer(context.Background(), c.args.uid)
			assert.ErrorIs(t, gotErr, c.wantErr)
			assert.Equal(t, c.want, got)
		})
	}
}
