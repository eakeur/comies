package stock

import (
	"comies/app/core/entities/stock"
	"comies/app/gateway/persistence/postgres/tests"
	"comies/app/sdk/throw"
	"comies/app/sdk/types"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_actions_GetByID(t *testing.T) {
	t.Parallel()

	type args struct {
		resourceID types.ID
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
			name: "should return stock with id",
			args: args{
				resourceID: 1,
			},
			want: stock.Stock{
				ID:              1,
				TargetID:        1,
				MinimumQuantity: 10,
				MaximumQuantity: 1000,
				Location:        "Under the table",
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
		},
		{
			name: "should return ErrNotFound error for nonexistent stock",
			args: args{
				resourceID: 1,
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
			got, err := a.GetByID(ctx, tt.args.resourceID)
			assert.ErrorIs(t, err, tt.wantErr)
			assert.Equalf(t, tt.want, got, "GetByID(%v)", tt.args.resourceID)
		})
	}
}
