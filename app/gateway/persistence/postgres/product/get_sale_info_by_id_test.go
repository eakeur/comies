package product

import (
	"context"
	"github.com/stretchr/testify/assert"
	"gomies/app/core/entities/product"
	"gomies/app/gateway/persistence/postgres/tests"
	"gomies/app/sdk/fault"
	"gomies/app/sdk/types"
	"testing"
)

func Test_actions_GetSaleInfoByID(t *testing.T) {
	t.Parallel()

	type args struct {
		productID types.ID
	}
	cases := []struct {
		name    string
		args    args
		want    product.Sale
		wantErr error
		before  tests.Callback
		after   tests.Callback
	}{
		{
			name: "should return sale data for product id",
			args: args{
				productID: 1,
			},
			want: product.Sale{
				CostPrice:   10,
				SalePrice:   20,
				SaleUnit:    types.Unit,
				MinimumSale: 1,
			},
			before: func(ctx context.Context, db *tests.Database, t *testing.T) {
				if _, err := db.InsertProducts(ctx, product.Product{
					ID:   1,
					Code: "PRDX",
					Name: "Product X",
					Type: product.OutputType,
					Sale: product.Sale{
						CostPrice:   10,
						SalePrice:   20,
						SaleUnit:    types.Unit,
						MinimumSale: 1,
					},
				}); err != nil {
					t.Error(err)
				}
			},
		},
		{
			name: "should return ErrNotFound error for nonexistent product",
			args: args{
				productID: 1,
			},
			wantErr: fault.ErrNotFound,
		},
	}
	for _, tt := range cases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctx, db := tests.FetchTestDB(t, tt.before)
			defer db.Drop(tt.after)

			a := actions{db: db.Pool}
			got, err := a.GetSaleInfoByID(ctx, tt.args.productID)
			assert.ErrorIs(t, err, tt.wantErr)
			assert.Equalf(t, tt.want, got, "GetByID(%v)", tt.args.productID)
		})
	}
}
