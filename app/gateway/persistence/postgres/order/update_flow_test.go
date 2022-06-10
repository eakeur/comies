package order

import (
	"context"
	"gomies/app/core/entities/order"
	"gomies/app/gateway/persistence/postgres/tests"
	"gomies/app/sdk/fault"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_actions_UpdateFlow(t *testing.T) {
	t.Parallel()

	var occurredAt = time.Now().UTC()

	type args struct {
		flow order.FlowUpdate
	}
	cases := []struct {
		name    string
		args    args
		want    order.FlowUpdate
		wantErr error
		before  tests.Callback
		after   tests.Callback
	}{
		{
			name: "should return created order",
			args: args{
				flow: order.FlowUpdate{
					ID:         1,
					OrderID:    1,
					Status:     order.DeliveringStatus,
					OccurredAt: occurredAt,
				},
			},
			want: order.FlowUpdate{
				ID:         1,
				OrderID:    1,
				Status:     order.DeliveringStatus,
				OccurredAt: occurredAt,
			},
			before: func(ctx context.Context, db *tests.Database, t *testing.T) {
				if _, err := db.InsertOrders(ctx, order.Order{ID: 1}); err != nil {
					t.Error(err)
				}
			},
			after: func(ctx context.Context, db *tests.Database, t *testing.T) {
				db.CheckValue(ctx, "select count(id) from orders_flow", int64(1))
			},
		},
		{
			name: "should return ErrAlreadyExists error for existing status",
			args: args{
				flow: order.FlowUpdate{
					ID:         5,
					OrderID:    1,
					Status:     order.DeliveringStatus,
					OccurredAt: occurredAt,
				},
			},
			wantErr: fault.ErrAlreadyExists,
			before: func(ctx context.Context, db *tests.Database, t *testing.T) {
				if _, err := db.InsertOrders(ctx, order.Order{ID: 1}); err != nil {
					t.Error(err)
				}

				if _, err := db.InsertOrdersFlow(ctx, order.FlowUpdate{ID: 4, OrderID: 1, Status: order.DeliveringStatus, OccurredAt: occurredAt}); err != nil {
					t.Error(err)
				}
			},
		},
		{
			name: "should return ErrAlreadyExists error for existing id",
			args: args{
				flow: order.FlowUpdate{
					ID:         4,
					OrderID:    1,
					Status:     order.FinishedStatus,
					OccurredAt: occurredAt,
				},
			},
			wantErr: fault.ErrAlreadyExists,
			before: func(ctx context.Context, db *tests.Database, t *testing.T) {
				if _, err := db.InsertOrders(ctx, order.Order{ID: 1}); err != nil {
					t.Error(err)
				}

				if _, err := db.InsertOrdersFlow(ctx, order.FlowUpdate{ID: 4, OrderID: 1, Status: order.DeliveringStatus, OccurredAt: occurredAt}); err != nil {
					t.Error(err)
				}
			},
		},
		{
			name: "should return ErrAlreadyExists error for existing id",
			args: args{
				flow: order.FlowUpdate{
					ID:         4,
					OrderID:    1,
					Status:     order.FinishedStatus,
					OccurredAt: occurredAt,
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
			got, err := a.UpdateFlow(ctx, tt.args.flow)
			assert.ErrorIs(t, err, tt.wantErr)
			assert.Equalf(t, tt.want, got, "UpdateFlow(%v)", tt.args.flow)
		})
	}
}
