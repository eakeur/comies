package stock

import (
	"comies/app/core/entities/stock"
	"comies/app/gateway/persistence/postgres/tests"
	"comies/app/sdk/fault"
	"comies/app/sdk/types"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_actions_Remove(t *testing.T) {
	t.Parallel()

	type args struct {
		resourceID types.ID
	}
	cases := []struct {
		name    string
		args    args
		wantErr error
		before  tests.Callback
		after   tests.Callback
	}{
		{
			name: "should delete stock successfully",
			args: args{
				resourceID: 1,
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
			after: func(ctx context.Context, d *tests.Database, t *testing.T) {
				d.CheckValue(ctx, "select count(id) from stocks", 0)
			},
		},
		{
			name: "should fail for nonexistent stock",
			args: args{
				resourceID: 1,
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
			err := a.Remove(ctx, tt.args.resourceID)
			assert.ErrorIs(t, err, tt.wantErr)
		})
	}
}
