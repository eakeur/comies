package order

import (
	"context"
	"gomies/app/core/entities/order"
	"gomies/app/gateway/persistence/postgres/tests"
	"gomies/app/sdk/fault"
	"gomies/app/sdk/types"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_actions_GetStatus(t *testing.T) {
	t.Parallel()

	type args struct {
		orderID types.ID
	}
	cases := []struct {
		name    string
		args    args
		want    order.Status
		wantErr error
		before  tests.Callback
		after   tests.Callback
	}{
		{
			name: "should return order status",
			args: args{
				orderID: 1,
			},
			want: order.DeliveringStatus,
			before: func(ctx context.Context, db *tests.Database, t *testing.T) {
				if _, err := db.InsertOrders(ctx, order.Order{ID: 1}); err != nil {
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
			},
		},
		{
			name: "should return ErrNotFound error for nonexistent order",
			args: args{
				orderID: 1,
			},
			wantErr: fault.ErrNotFound,
		},
	}
	for _, tt := range cases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctx, db := tests.FetchTestDB(t, tt.before)
			defer db.Drop(tt.after)

			a := actions{db: db.Pool}
			got, err := a.GetStatus(ctx, tt.args.orderID)
			assert.ErrorIs(t, err, tt.wantErr)
			assert.Equalf(t, tt.want, got, "GetStatus(%v)", tt.args.orderID)
		})
	}
}
