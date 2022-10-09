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

func Test_actions_SetStatus(t *testing.T) {
	t.Parallel()

	type args struct {
		itemID types.ID
		status item.Status
	}
	cases := []struct {
		name    string
		args    args
		wantErr error
		before  tests.Callback
		after   tests.Callback
	}{
		{
			name: "should update successfully",
			args: args{
				itemID: 1,
				status: item.DoneStatus,
			},
			before: func(ctx context.Context, d *tests.Database, t *testing.T) {
				_, err := d.InsertOrders(ctx, order.Order{ID: 1})
				if err != nil {
					t.Error(err)
				}

				_, err = d.InsertItems(ctx,
					item.Item{ID: 1, OrderID: 1, ProductID: 2},
				)
				if err != nil {
					t.Error(err)
				}
			},
			after: func(ctx context.Context, d *tests.Database, t *testing.T) {
				var status item.Status
				r := d.Pool.QueryRow(ctx, "select status from items limit 1")
				if err := r.Scan(&status); err != nil || status != item.DoneStatus {
					t.Errorf("the item was not updated: %v", err)
				}
			},
		},
		{
			name: "should fail for nonexistent item",
			args: args{
				itemID: 1,
				status: item.DoneStatus,
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
			err := a.SetStatus(ctx, tt.args.itemID, tt.args.status)
			assert.ErrorIs(t, err, tt.wantErr)
		})
	}
}
