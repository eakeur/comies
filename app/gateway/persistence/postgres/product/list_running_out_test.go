package product

import (
	"comies/app/core/entities/movement"
	"comies/app/core/entities/product"
	"comies/app/gateway/persistence/postgres/tests"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_actions_ListRunningOut(t *testing.T) {
	t.Parallel()

	type args struct {
	}
	cases := []struct {
		name    string
		args    args
		before  tests.Callback
		want    []product.Product
		wantErr error
	}{
		{
			name: "should return empty list of products",
			args: args{},
			want: []product.Product{},
		},
		{
			name: "should return list with products with equal or less than 25% of the maximum",
			args: args{},
			want: []product.Product{
				{
					ID:      3,
					Code:    "PRDC",
					Name:    "Product C",
					Type:    product.OutputType,
					Balance: 110,
					Stock: product.Stock{
						MaximumQuantity: 4300,
						MinimumQuantity: 100,
					},
				},
				{
					ID:      1,
					Code:    "PRDA",
					Name:    "Product A",
					Type:    product.OutputType,
					Sale:    product.Sale{},
					Balance: 20,
					Stock: product.Stock{
						MaximumQuantity: 100,
						MinimumQuantity: 0,
					},
				},
			},
			before: func(ctx context.Context, d *tests.Database, t *testing.T) {
				_, err := d.InsertProducts(ctx, product.Product{
					ID:   1,
					Code: "PRDA",
					Name: "Product A",
					Type: product.OutputType,
					Sale: product.Sale{},
					Stock: product.Stock{
						MaximumQuantity: 100,
						MinimumQuantity: 0,
					},
				}, product.Product{
					ID:   2,
					Code: "PRDB",
					Name: "Product B",
					Type: product.OutputType,
					Stock: product.Stock{
						MaximumQuantity: 350,
						MinimumQuantity: 10,
					},
				}, product.Product{
					ID:   3,
					Code: "PRDC",
					Name: "Product C",
					Type: product.OutputType,
					Stock: product.Stock{
						MaximumQuantity: 4300,
						MinimumQuantity: 100,
					},
				})
				if err != nil {
					t.Error(err)
				}

				_, err = d.InsertMovements(ctx, movement.Movement{
					ID:        1,
					ProductID: 1,
					Type:      movement.InputType,
					Quantity:  20,
					PaidValue: 0,
					AgentID:   0,
				}, movement.Movement{
					ID:        2,
					ProductID: 2,
					Type:      movement.InputType,
					Quantity:  200,
					PaidValue: 0,
					AgentID:   0,
				}, movement.Movement{
					ID:        3,
					ProductID: 3,
					Type:      movement.InputType,
					Quantity:  800,
					PaidValue: 0,
					AgentID:   0,
				}, movement.Movement{
					ID:        4,
					ProductID: 3,
					Type:      movement.OutputType,
					Quantity:  690,
					PaidValue: 0,
					AgentID:   0,
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

			ctx, db := tests.FetchTestDB(t, tt.before)
			defer db.Drop()

			got, err := actions{db: db.Pool}.ListRunningOut(ctx)
			assert.ErrorIs(t, err, tt.wantErr)
			assert.Equal(t, tt.want, got)
		})
	}
}
