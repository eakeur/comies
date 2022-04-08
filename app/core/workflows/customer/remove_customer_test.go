package customer

import (
	"context"
	"gomies/app/core/entities/customer"
	"gomies/pkg/sdk/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWorkflow_Remove(t *testing.T) {
	t.Parallel()

	type (
		args struct {
			uid types.UID
		}

		fields struct {
			customers *customer.ActionsMock
		}

		test struct {
			name    string
			args    args
			fields  fields
			wantErr error
		}
	)

	cases := []test{
		{
			name: "should return member found",
			args: args{
				uid: types.NewUID(),
			},
			fields: fields{
				customers: &customer.ActionsMock{
					RemoveCustomerFunc: func(ctx context.Context, uid types.UID) error {
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

			err := workflow{customers: c.fields.customers}.RemoveCustomer(context.Background(), c.args.uid)
			assert.ErrorIs(t, err, c.wantErr)

		})
	}
}
