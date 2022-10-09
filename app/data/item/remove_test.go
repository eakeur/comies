package item

import (
	"comies/app/core/entities/item"
	"comies/app/core/entities/order"
	"comies/app/core/types"
	"comies/app/gateway/persistence/postgres/tests"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_actions_Remove(t *testing.T) {
	t.Parallel()

	type args struct {
		itemID types.ID
	}
	cases := []struct {
		name    string
		args    args
		wantErr error
		before  tests.Callback
		after   tests.Callback
	}{
		{
			name: "should delete item successfully",
			args: args{
				itemID: 1,
			},
			before: func(ctx context.Context, d *tests.Database, t *testing.T) {
				_, err := d.InsertOrders(ctx, order.Order{ID: 1})
				if err != nil {
					t.Error(err)
				}

				_, err = d.InsertItems(ctx,
					item.Item{ID: 1, OrderID: 1, ProductID: 1},
					item.Item{ID: 2, OrderID: 1, ProductID: 2},
					item.Item{ID: 3, OrderID: 1, ProductID: 3},
				)
				if err != nil {
					t.Error(err)
				}
			},
			after: func(ctx context.Context, d *tests.Database, t *testing.T) {
				d.CheckValue(t, ctx, "select count(id) from items", int64(2))
			},
		},
		{
			name: "should fail for nonexistent item",
			args: args{
				itemID: 1,
			},
			wantErr: types.ErrNotFound,
		},
	}
	for _, tt := range cases {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctx, _ := tests.FetchTestTX(t, tt.before, tt.after)

			a := actions{}
			err := a.Remove(ctx, tt.args.itemID)
			assert.ErrorIs(t, err, tt.wantErr)
		})
	}
}
