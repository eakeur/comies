package product

import (
	"comies/app/core/entities/product"
	"comies/app/gateway/persistence/postgres/tests"
	"comies/app/sdk/types"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_actions_Update(t *testing.T) {
	t.Parallel()

	type args struct {
		product product.Product
	}
	cases := []struct {
		name    string
		args    args
		wantErr error
		before  tests.Callback
		after   tests.Callback
	}{
		{
			name: "should update product successfully",
			args: args{
				product: product.Product{
					ID:   1,
					Code: "PRDXTA",
					Name: "Product XTA",
					Type: product.InputType,
					Sale: product.Sale{
						CostPrice:   20,
						SalePrice:   30,
						SaleUnit:    types.Kilogram,
						MinimumSale: 2,
					},
				},
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
					from products where id = $1
				`
				d.CheckValue(ctx, script, true, 1, "PRDXTA", "Product XTA", product.InputType, 20, 30, types.Kilogram, 2)
			},
		},
		{
			name: "should fail for repeated product code",
			args: args{
				product: product.Product{
					ID:   1,
					Code: "PRDXTA",
					Name: "Product XTA",
					Type: product.InputType,
					Sale: product.Sale{
						CostPrice:   20,
						SalePrice:   30,
						SaleUnit:    types.Kilogram,
						MinimumSale: 2,
					},
				},
			},
			wantErr: product.ErrCodeAlreadyExists,
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
				}, product.Product{
					ID:   2,
					Code: "PRDXTA",
					Name: "Product XTA",
					Type: product.OutputType,
					Sale: product.Sale{
						CostPrice:   10,
						SalePrice:   20,
						SaleUnit:    types.Unit,
						MinimumSale: 1,
					}})
				if err != nil {
					t.Error(err)
				}
			},
		},
		{
			name: "should fail for nonexistent movement",
			args: args{
				product: product.Product{
					ID:   1,
					Code: "PRDXTA",
					Name: "Product XTA",
					Type: product.InputType,
					Sale: product.Sale{
						CostPrice:   20,
						SalePrice:   30,
						SaleUnit:    types.Kilogram,
						MinimumSale: 2,
					},
				},
			},
			wantErr: product.ErrNotFound,
		},
	}
	for _, tt := range cases {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctx, db := tests.FetchTestTX(t, tt.before)
			defer db.Drop(tt.after)

			a := actions{}
			err := a.Update(ctx, tt.args.product)
			assert.ErrorIs(t, err, tt.wantErr)
		})
	}
}
