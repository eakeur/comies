package product

import (
	"comies/app/core/entities/product"
	"comies/app/gateway/persistence/postgres/tests"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_actions_List(t *testing.T) {
	t.Parallel()

	type args struct {
		filter product.Filter
	}
	cases := []struct {
		name    string
		args    args
		before  tests.Callback
		want    []product.Product
		wantErr error
	}{
		{
			name: "should return empty list of orders",
			args: args{},
			want: []product.Product{},
		},
		{
			name: "should return list with code filtered products",

			args: args{
				filter: product.Filter{
					Code: "PRD",
				},
			},
			want: []product.Product{
				{
					ID:   1,
					Code: "PRDA",
					Name: "Product A",
					Type: product.OutputType,
					Sale: product.Sale{},
				}, {
					ID:   2,
					Code: "PRDB",
					Name: "Product B",
					Type: product.OutputType,
					Sale: product.Sale{},
				}, {
					ID:   3,
					Code: "PRDC",
					Name: "Product C",
					Type: product.OutputType,
					Sale: product.Sale{},
				},
			},
			before: func(ctx context.Context, d *tests.Database, t *testing.T) {
				_, err := d.InsertProducts(ctx, product.Product{
					ID:   1,
					Code: "PRDA",
					Name: "Product A",
					Type: product.OutputType,
					Sale: product.Sale{},
				}, product.Product{
					ID:   2,
					Code: "PRDB",
					Name: "Product B",
					Type: product.OutputType,
					Sale: product.Sale{},
				}, product.Product{
					ID:   3,
					Code: "PRDC",
					Name: "Product C",
					Type: product.OutputType,
					Sale: product.Sale{},
				}, product.Product{
					ID:   4,
					Code: "PRAD",
					Name: "Product D",
					Type: product.OutputType,
					Sale: product.Sale{},
				})
				if err != nil {
					t.Error(err)
				}
			},
		},
		{
			name: "should return list with name filtered products",

			args: args{
				filter: product.Filter{
					Name: "Meat",
				},
			},
			want: []product.Product{
				{
					ID:   1,
					Code: "PRDA",
					Name: "Baked Meat",
					Type: product.OutputType,
					Sale: product.Sale{},
				}, {
					ID:   2,
					Code: "PRDB",
					Name: "Fried Meat",
					Type: product.OutputType,
					Sale: product.Sale{},
				}, {
					ID:   3,
					Code: "PRDC",
					Name: "Raw Meat",
					Type: product.OutputType,
					Sale: product.Sale{},
				},
			},
			before: func(ctx context.Context, d *tests.Database, t *testing.T) {
				_, err := d.InsertProducts(ctx, product.Product{
					ID:   1,
					Code: "PRDA",
					Name: "Baked Meat",
					Type: product.OutputType,
					Sale: product.Sale{},
				}, product.Product{
					ID:   2,
					Code: "PRDB",
					Name: "Fried Meat",
					Type: product.OutputType,
					Sale: product.Sale{},
				}, product.Product{
					ID:   3,
					Code: "PRDC",
					Name: "Raw Meat",
					Type: product.OutputType,
					Sale: product.Sale{},
				}, product.Product{
					ID:   4,
					Code: "PRAD",
					Name: "Orange Juice",
					Type: product.OutputType,
					Sale: product.Sale{},
				})
				if err != nil {
					t.Error(err)
				}
			},
		},
		{
			name: "should return list with type filtered products",

			args: args{
				filter: product.Filter{
					Type: product.InputType,
				},
			},
			want: []product.Product{
				{
					ID:   4,
					Code: "PRAD",
					Name: "Orange Juice",
					Type: product.InputType,
					Sale: product.Sale{},
				},
			},
			before: func(ctx context.Context, d *tests.Database, t *testing.T) {
				_, err := d.InsertProducts(ctx, product.Product{
					ID:   1,
					Code: "PRDA",
					Name: "Baked Meat",
					Type: product.OutputType,
					Sale: product.Sale{},
				}, product.Product{
					ID:   2,
					Code: "PRDB",
					Name: "Fried Meat",
					Type: product.OutputType,
					Sale: product.Sale{},
				}, product.Product{
					ID:   3,
					Code: "PRDC",
					Name: "Raw Meat",
					Type: product.OutputType,
					Sale: product.Sale{},
				}, product.Product{
					ID:   4,
					Code: "PRAD",
					Name: "Orange Juice",
					Type: product.InputType,
					Sale: product.Sale{},
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

			got, err := actions{db: db.Pool}.List(ctx, tt.args.filter)
			assert.ErrorIs(t, err, tt.wantErr)
			assert.Equal(t, tt.want, got)
		})
	}
}
