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
					item.Item{ID: 1, OrderID: 1, ProductID: 1},
					item.Item{ID: 2, OrderID: 1, ProductID: 2},
					item.Item{ID: 3, OrderID: 1, ProductID: 3},
				)
				if err != nil {
					t.Error(err)
				}
			},
			args: args{
				orderID: 1,
			},
			want: []item.Item{
				{ID: 1, OrderID: 1, ProductID: 1},
				{ID: 2, OrderID: 1, ProductID: 2},
				{ID: 3, OrderID: 1, ProductID: 3},
			},
		},
	}
	for _, tt := range cases {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctx, db := tests.FetchTestDB(t, tt.before)

			got, err := actions{db: db.Pool}.List(ctx, tt.args.orderID)
			assert.ErrorIs(t, err, tt.wantErr)
			assert.Equal(t, tt.want, got)
		})
	}
}
