package stock

import (
	"comies/app/core/entities/stock"
	"comies/app/gateway/persistence/postgres/tests"
	"comies/app/sdk/fault"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_actions_Update(t *testing.T) {
	t.Parallel()

	type args struct {
		stock stock.Stock
	}
	cases := []struct {
		name    string
		args    args
		wantErr error
		before  tests.Callback
		after   tests.Callback
	}{
		{
			name: "should update stock successfully",
			args: args{
				stock: stock.Stock{
					ID:              1,
					TargetID:        1,
					MinimumQuantity: 100,
					MaximumQuantity: 10000,
					Location:        "Under the blue table",
				},
			},
			before: func(ctx context.Context, db *tests.Database, t *testing.T) {
				_, err := db.InsertStocks(ctx, stock.Stock{
					ID:              1,
					TargetID:        1,
					MinimumQuantity: 10,
					MaximumQuantity: 1000,
					Location:        "Under the table",
				})
				if err != nil {
					t.Error(err)
				}
			},
			after: func(ctx context.Context, db *tests.Database, t *testing.T) {
				const script = `
					select
						minimum_quantity = $1 and 
						maximum_quantity = $2 and 
						location = $3
					as equal
					from stocks where id = $4
				`
				db.CheckValue(ctx, script, true, 100, 10000, "Under the blue table", 1)
			},
		},
		{
			name: "should fail for nonexistent stock",
			args: args{
				stock: stock.Stock{
					ID:              1,
					TargetID:        1,
					MinimumQuantity: 100,
					MaximumQuantity: 10000,
					Location:        "Under the blue table",
				},
			},
			wantErr: fault.ErrNotFound,
		},
	}
	for _, tt := range cases {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctx, db := tests.FetchTestTX(t, tt.before)
			defer db.Drop(tt.after)

			a := actions{}
			err := a.Update(ctx, tt.args.stock)
			assert.ErrorIs(t, err, tt.wantErr)
		})
	}
}
