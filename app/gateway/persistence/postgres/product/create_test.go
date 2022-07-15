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

func Test_actions_Create(t *testing.T) {
	t.Parallel()

	type args struct {
		product product.Product
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
			name: "should return created product",
			args: args{
				product: product.Product{
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
			after: func(ctx context.Context, db *tests.Database, t *testing.T) {
				const script = `
					select 
						id = $1 and 
						code = $2 and 
						name = $3 and 
						type = $4 and 
						cost_price = $5 and 
						sale_price = $6 and 
						sale_unit = $7 and
						minimum_sale = $8
					as equal
					from products where id = $1
				`
				db.CheckValue(ctx, script, true, 1, "PRDX", "Product X", product.OutputType, 10, 20, types.Unit, 1)
			},
		},
		{
			name: "should fail for existent product id",
			args: args{
				product: product.Product{
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
			},
			wantErr: fault.ErrAlreadyExists,
			before: func(ctx context.Context, db *tests.Database, t *testing.T) {
				_, err := db.InsertProducts(ctx, product.Product{
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
		},
		{
			name: "should fail for existent product code",
			args: args{
				product: product.Product{
					ID:   2,
					Code: "PRDXT",
					Name: "Product X",
					Type: product.OutputType,
					Sale: product.Sale{
						CostPrice:   10,
						SalePrice:   20,
						SaleUnit:    types.Unit,
						MinimumSale: 1,
					},
				},
			},
			wantErr: fault.ErrAlreadyExists,
			before: func(ctx context.Context, db *tests.Database, t *testing.T) {
				_, err := db.InsertProducts(ctx, product.Product{
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
		},
	}
	for _, tt := range cases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctx, db := tests.FetchTestTX(t, tt.before)
			defer db.Drop(tt.after)

			a := actions{}
			got, err := a.Create(ctx, tt.args.product)
			assert.ErrorIs(t, err, tt.wantErr)
			assert.Equalf(t, tt.want, got, "Create(%v)", tt.args.product)
		})
	}
}
