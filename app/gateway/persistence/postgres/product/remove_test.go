package product

import (
	"comies/app/core/entities/product"
	"comies/app/gateway/persistence/postgres/tests"
	"comies/app/sdk/throw"
	"comies/app/sdk/types"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_actions_Remove(t *testing.T) {
	t.Parallel()

	type args struct {
		id types.ID
	}
	cases := []struct {
		name    string
		args    args
		wantErr error
		before  tests.Callback
		after   tests.Callback
	}{
		{
			name: "should delete product successfully",
			args: args{
				id: 1,
			},
			before: func(ctx context.Context, d *tests.Database, t *testing.T) {
				_, err := d.InsertProducts(ctx, product.Product{
					ID:   1,
					Code: "PRDXT",
					Name: "Product X",
					Type: product.OutputType,
					Sale: product.Sale{
						CostPrice:   10,
						SalePrice:   20,
						SaleUnit:    types.Unit,
						MinimumSale: 1,
					},
				})
				if err != nil {
					t.Error(err)
				}
			},
			after: func(ctx context.Context, d *tests.Database, t *testing.T) {
				d.CheckValue(ctx, "select count(id) from products", 0)
			},
		},
		{
			name: "should fail for nonexistent product",
			args: args{
				id: 1,
			},
			wantErr: throw.ErrNotFound,
		},
	}
	for _, tt := range cases {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctx, db := tests.FetchTestTX(t, tt.before)
			defer db.Drop(tt.after)

			a := actions{}
			err := a.Remove(ctx, tt.args.id)
			assert.ErrorIs(t, err, tt.wantErr)
		})
	}
}
