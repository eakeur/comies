package order

import (
	"comies/app/core/entities/item"
	"comies/app/core/entities/order"
	"comies/app/core/throw"
	"comies/app/core/types"
	"comies/app/gateway/persistence/postgres/tests"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_actions_GetByID(t *testing.T) {
	t.Parallel()

	var placed = time.Date(2001, time.September, 30, 22, 45, 00, 0, time.UTC)

	type args struct {
		id types.ID
	}
	cases := []struct {
		name    string
		args    args
		want    order.Order
		wantErr error
		before  tests.Callback
		after   tests.Callback
	}{
		{
			name: "should return order with id",
			args: args{
				id: 1,
			},
			want: order.Order{
				ID:         1,
				PlacedAt:   placed,
				Status:     order.DeliveringStatus,
				FinalPrice: 300,
			},
			before: func(ctx context.Context, db *tests.Database, t *testing.T) {
				if _, err := db.InsertOrders(ctx, order.Order{ID: 1, PlacedAt: placed}); err != nil {
					t.Error(err)
				}

				if _, err := db.InsertOrdersFlow(ctx,
					order.FlowUpdate{ID: 1, OrderID: 1, Status: order.InTheCartStatus, OccurredAt: time.Now().UTC()},
					order.FlowUpdate{ID: 2, OrderID: 1, Status: order.PendingStatus, OccurredAt: time.Now().UTC()},
					order.FlowUpdate{ID: 3, OrderID: 1, Status: order.PreparingStatus, OccurredAt: time.Now().UTC()},
					order.FlowUpdate{ID: 4, OrderID: 1, Status: order.DeliveringStatus, OccurredAt: time.Now().UTC()},
				); err != nil {
					t.Error(err)
				}

				if _, err := db.InsertItems(ctx,
					item.Item{ID: 1, OrderID: 1, ProductID: 1, Price: 100},
					item.Item{ID: 2, OrderID: 1, ProductID: 2, Price: 100},
					item.Item{ID: 3, OrderID: 1, ProductID: 3, Price: 100},
				); err != nil {
					t.Error(err)
				}
			},
		},
		{
			name: "should return ErrNotFound error for nonexistent order",
			args: args{
				id: 1,
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
			got, err := a.GetByID(ctx, tt.args.id)
			assert.ErrorIs(t, err, tt.wantErr)
			assert.Equalf(t, tt.want, got, "GetByID(%v)", tt.args.id)
		})
	}
}
