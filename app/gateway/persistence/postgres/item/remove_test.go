package item

import (
	"context"
	"github.com/stretchr/testify/assert"
	"gomies/app/core/entities/item"
	"gomies/app/core/entities/order"
	"gomies/app/gateway/persistence/postgres/tests"
	"gomies/app/sdk/fault"
	"gomies/app/sdk/types"
	"testing"
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
					item.Item{ID: 1, OrderID: 1, ProductID: 1, Store: tests.DefaultStore},
					item.Item{ID: 2, OrderID: 1, ProductID: 2, Store: tests.DefaultStore},
					item.Item{ID: 3, OrderID: 1, ProductID: 3, Store: tests.DefaultStore},
				)
				if err != nil {
					t.Error(err)
				}
			},
			after: func(ctx context.Context, d *tests.Database, t *testing.T) {
				var quantityInserted item.Status
				r := d.Pool.QueryRow(ctx, "select count(id) from items")
				if err := r.Scan(&quantityInserted); err != nil || quantityInserted != 2 {
					t.Errorf("the item was not deleted: %v", err)
				}
			},
		},
		{
			name: "should fail for nonexistent item",
			args: args{
				itemID: 1,
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
			err := a.Remove(ctx, tt.args.itemID)
			assert.ErrorIs(t, err, tt.wantErr)
		})
	}
}
