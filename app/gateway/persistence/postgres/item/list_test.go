package item

import (
	"context"
	"github.com/stretchr/testify/assert"
	"gomies/app/core/entities/item"
	"gomies/app/core/entities/order"
	"gomies/app/gateway/persistence/postgres/tests"
	"gomies/app/sdk/types"
	"testing"
)

func Test_actions_List(t *testing.T) {
	t.Parallel()

	type args struct {
		orderID types.ID
	}
	cases := []struct {
		name    string
		args    args
		before  tests.Callback
		want    []item.Item
		wantErr error
	}{
		{
			name: "should return empty list of items",
			args: args{
				orderID: 1,
			},
			want: []item.Item{},
		},

		{
			name: "should return list with items",
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
			args: args{
				orderID: 1,
			},
			want: []item.Item{
				{ID: 1, OrderID: 1, ProductID: 1, Store: tests.DefaultStore},
				{ID: 2, OrderID: 1, ProductID: 2, Store: tests.DefaultStore},
				{ID: 3, OrderID: 1, ProductID: 3, Store: tests.DefaultStore},
			},
		},
	}
	for _, tt := range cases {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctx, db := tests.FetchTestDB(t, tt.before)
			defer db.Drop()

			got, err := actions{db: db.Pool}.List(ctx, tt.args.orderID)
			assert.ErrorIs(t, err, tt.wantErr)
			assert.Equal(t, tt.want, got)
		})
	}
}
