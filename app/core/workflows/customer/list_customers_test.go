package customer

import (
	"context"
	"github.com/stretchr/testify/assert"
	"gomies/app/core/entities/customer"
	"testing"
)

func TestWorkflow_ListCustomers(t *testing.T) {
	t.Parallel()

	type (
		fields struct {
			customers *customer.ActionsMock
		}

		args struct {
			filter customer.Filter
		}

		test struct {
			name    string
			args    args
			want    []customer.Customer
			wantErr error
			fields  fields
		}
	)

	for _, c := range []test{
		{
			name: "should return customers",
			args: args{},
			want: []customer.Customer{
				{
					Name:              "Granger",
					PhoneDigest:       "99558874;",
					AddressCodeDigest: "03322550;",
				},
			},
			fields: fields{
				customers: &customer.ActionsMock{
					ListCustomersFunc: func(ctx context.Context, f customer.Filter) ([]customer.Customer, int, error) {
						return []customer.Customer{
							{
								Name:              "Granger",
								PhoneDigest:       "99558874;",
								AddressCodeDigest: "03322550;",
							},
						}, 1, nil
					},
				},
			},
		},
	} {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			wf := workflow{customers: c.fields.customers}
			got, _, gotErr := wf.ListCustomers(context.Background(), c.args.filter)
			assert.ErrorIs(t, gotErr, c.wantErr)
			assert.Equal(t, c.want, got)
		})
	}
}
