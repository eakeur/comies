package order

import (
	"comies/app/core/entities/order"
	"comies/app/core/types"
	"comies/app/gateway/persistence/postgres/tests"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_actions_ListFlow(t *testing.T) {
	t.Parallel()

	var eventTime = time.Date(2001, time.September, 30, 22, 45, 00, 0, time.UTC)

	type args struct {
		orderID types.ID
	}
	cases := []struct {
		name    string
		args    args
		before  tests.Callback
		want    []order.FlowUpdate
		wantErr error
	}{
		{
			name: "should return empty list of order flows",
			args: args{},
			want: []order.FlowUpdate{},
		},
		{
			name: "should return list with orders flows",

			args: args{
				orderID: 1,
			},
			want: []order.FlowUpdate{
				{ID: 1, OrderID: 1, Status: order.InTheCartStatus, OccurredAt: eventTime},
				{ID: 2, OrderID: 1, Status: order.PreparingStatus, OccurredAt: eventTime},
				{ID: 3, OrderID: 1, Status: order.WaitingDeliveryStatus, OccurredAt: eventTime},
			},
			before: func(ctx context.Context, d *tests.Database, t *testing.T) {
				_, err := d.InsertOrders(ctx,
					order.Order{ID: 1},
				)
				if err != nil {
					t.Error(err)
				}

				_, err = d.InsertOrdersFlow(ctx,
					order.FlowUpdate{ID: 1, OrderID: 1, Status: order.InTheCartStatus, OccurredAt: eventTime},
					order.FlowUpdate{ID: 2, OrderID: 1, Status: order.PreparingStatus, OccurredAt: eventTime},
					order.FlowUpdate{ID: 3, OrderID: 1, Status: order.WaitingDeliveryStatus, OccurredAt: eventTime},
				)
				if err != nil {
					t.Error(err)
				}
			},
		},
	}
	for _, tt := range cases {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctx, db := tests.FetchTestDB(t, tt.before)
			defer db.Drop()

			got, err := actions{db: db.Pool}.ListFlow(ctx, tt.args.orderID)
			assert.ErrorIs(t, err, tt.wantErr)
			assert.Equal(t, tt.want, got)
		})
	}
}
