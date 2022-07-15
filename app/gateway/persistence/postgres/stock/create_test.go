package stock

import (
	"comies/app/core/entities/stock"
	"comies/app/gateway/persistence/postgres/tests"
	"comies/app/sdk/fault"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_actions_Create(t *testing.T) {
	t.Parallel()

	type args struct {
		stock stock.Stock
	}
	cases := []struct {
		name    string
		args    args
		want    stock.Stock
		wantErr error
		before  tests.Callback
		after   tests.Callback
	}{
		{
			name: "should return created stock",
			args: args{
				stock: stock.Stock{
					ID:              1,
					TargetID:        1,
					MinimumQuantity: 10,
					MaximumQuantity: 1000,
					Location:        "Under the table",
				},
			},
			want: stock.Stock{
				ID:              1,
				TargetID:        1,
				MinimumQuantity: 10,
				MaximumQuantity: 1000,
				Location:        "Under the table",
			},
			after: func(ctx context.Context, db *tests.Database, t *testing.T) {
				const script = `
					select 
						id = $1 and 
						target_id = $2 and 
						minimum_quantity = $3 and 
						maximum_quantity = $4 and 
						location = $5
					as equal
					from stocks where id = $1
				`
				db.CheckValue(ctx, script, true, 1, 1, 10, 1000, "Under the table")
			},
		},
		{
			name: "should fail for existent target id",
			args: args{
				stock: stock.Stock{
					ID:              1,
					TargetID:        1,
					MinimumQuantity: 10,
					MaximumQuantity: 1000,
					Location:        "Under the table",
				},
			},
			wantErr: fault.ErrAlreadyExists,
			before: func(ctx context.Context, db *tests.Database, t *testing.T) {
				_, err := db.InsertStocks(ctx, stock.Stock{
					ID:              10,
					TargetID:        1,
					MinimumQuantity: 10,
					MaximumQuantity: 1000,
					Location:        "Under the table",
				})
				if err != nil {
					t.Error(err)
				}
			},
		},
		{
			name: "should fail for existent stock id",
			args: args{
				stock: stock.Stock{
					ID:              1,
					TargetID:        1,
					MinimumQuantity: 10,
					MaximumQuantity: 1000,
					Location:        "Under the table",
				},
			},
			wantErr: fault.ErrAlreadyExists,
			before: func(ctx context.Context, db *tests.Database, t *testing.T) {
				_, err := db.InsertStocks(ctx, stock.Stock{
					ID:              1,
					TargetID:        10,
					MinimumQuantity: 10,
					MaximumQuantity: 1000,
					Location:        "Under the table",
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
			got, err := a.Create(ctx, tt.args.stock)
			assert.ErrorIs(t, err, tt.wantErr)
			assert.Equalf(t, tt.want, got, "Create(%v)", tt.args.stock)
		})
	}
}
