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

func Test_actions_GetByCode(t *testing.T) {
	t.Parallel()

	type args struct {
		code string
	}
	cases := []struct {
		name    string
		args    args
		want    product.Product
		wantErr error
		before  tests.Callback
		after   tests.Callback
	}{
		{
			name: "should return product with code",
			args: args{
				code: "PRDX",
			},
			want: product.Product{
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
				code: "PRDX",
			},
			wantErr: throw.ErrNotFound,
		},
	}
	for _, tt := range cases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctx, db := tests.FetchTestDB(t, tt.before)
			defer db.Drop(tt.after)

			a := actions{db: db.Pool}
			got, err := a.GetByCode(ctx, tt.args.code)
			assert.ErrorIs(t, err, tt.wantErr)
			assert.Equalf(t, tt.want, got, "GetByCode(%v)", tt.args.code)
		})
	}
}
