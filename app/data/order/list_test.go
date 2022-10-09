package order

import (
	"comies/app/core/entities/item"
	"comies/app/core/entities/order"
	"comies/app/gateway/persistence/postgres/tests"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_actions_List(t *testing.T) {
	t.Parallel()

	type args struct {
		filter order.Filter
	}
	cases := []struct {
		name    string
		args    args
		before  tests.Callback
		want    []order.Order
		wantErr error
	}{
		{
			name: "should return empty list of orders",
			args: args{},
			want: []order.Order{},
		},
		{
			name: "should return list with status filtered orders",

			args: args{
				filter: order.Filter{
					Status: []order.Status{
						order.WaitingDeliveryStatus,
						order.CanceledStatus,
					},
				},
			},
			want: []order.Order{
				{ID: 1, Status: order.WaitingDeliveryStatus, FinalPrice: 20},
				{ID: 3, Status: order.CanceledStatus},
				{ID: 4, Status: order.WaitingDeliveryStatus, FinalPrice: 30},
			},
			before: func(ctx context.Context, d *tests.Database, t *testing.T) {
				_, err := d.InsertOrders(ctx,
					order.Order{ID: 1},
					order.Order{ID: 2},
					order.Order{ID: 3},
					order.Order{ID: 4},
				)
				if err != nil {
					t.Error(err)
				}

				_, err = d.InsertOrdersFlow(ctx,
					order.FlowUpdate{ID: 1, OrderID: 1, Status: order.InTheCartStatus, OccurredAt: time.Now().UTC()},
					order.FlowUpdate{ID: 2, OrderID: 1, Status: order.PreparingStatus, OccurredAt: time.Now().UTC().Add(time.Second * 1)},
					order.FlowUpdate{ID: 3, OrderID: 1, Status: order.WaitingDeliveryStatus, OccurredAt: time.Now().UTC().Add(time.Second * 2)},
					order.FlowUpdate{ID: 4, OrderID: 2, Status: order.InTheCartStatus, OccurredAt: time.Now().UTC()},
					order.FlowUpdate{ID: 5, OrderID: 2, Status: order.PreparingStatus, OccurredAt: time.Now().UTC().Add(time.Second * 1)},
					order.FlowUpdate{ID: 6, OrderID: 2, Status: order.FinishedStatus, OccurredAt: time.Now().UTC().Add(time.Second * 2)},
					order.FlowUpdate{ID: 7, OrderID: 3, Status: order.InTheCartStatus, OccurredAt: time.Now().UTC()},
					order.FlowUpdate{ID: 8, OrderID: 3, Status: order.PreparingStatus, OccurredAt: time.Now().UTC().Add(time.Second * 1)},
					order.FlowUpdate{ID: 9, OrderID: 3, Status: order.CanceledStatus, OccurredAt: time.Now().UTC().Add(time.Second * 2)},
					order.FlowUpdate{ID: 10, OrderID: 4, Status: order.InTheCartStatus, OccurredAt: time.Now().UTC()},
					order.FlowUpdate{ID: 11, OrderID: 4, Status: order.PreparingStatus, OccurredAt: time.Now().UTC().Add(time.Second * 1)},
					order.FlowUpdate{ID: 12, OrderID: 4, Status: order.WaitingDeliveryStatus, OccurredAt: time.Now().UTC().Add(time.Second * 2)},
				)
				if err != nil {
					t.Error(err)
				}

				_, err = d.InsertItems(ctx,
					item.Item{ID: 1, OrderID: 1, ProductID: 1, Price: 10},
					item.Item{ID: 2, OrderID: 1, ProductID: 2, Price: 10},
					item.Item{ID: 3, OrderID: 4, ProductID: 44, Price: 15},
					item.Item{ID: 4, OrderID: 4, ProductID: 55, Price: 15},
				)
				if err != nil {
					t.Error(err)
				}
			},
		},
		{
			name: "should return list with time filtered orders",

			args: args{
				filter: order.Filter{
					PlacedAfter:  time.Date(2001, time.September, 29, 12, 45, 00, 0, time.UTC),
					PlacedBefore: time.Date(2001, time.September, 30, 22, 45, 00, 0, time.UTC),
				},
			},
			want: []order.Order{
				{ID: 1, FinalPrice: 201, Status: order.PreparingStatus, PlacedAt: time.Date(2001, time.September, 29, 12, 45, 00, 0, time.UTC)},
				{ID: 2, FinalPrice: 150, Status: order.PreparingStatus, PlacedAt: time.Date(2001, time.September, 30, 22, 45, 00, 0, time.UTC)},
			},
			before: func(ctx context.Context, d *tests.Database, t *testing.T) {
				_, err := d.InsertOrders(ctx,
					order.Order{ID: 1, PlacedAt: time.Date(2001, time.September, 29, 12, 45, 00, 0, time.UTC)},
					order.Order{ID: 2, PlacedAt: time.Date(2001, time.September, 30, 22, 45, 00, 0, time.UTC)},
					order.Order{ID: 3, PlacedAt: time.Now().UTC()},
					order.Order{ID: 4, PlacedAt: time.Now().UTC()},
					order.Order{ID: 5, PlacedAt: time.Date(2001, time.August, 30, 13, 00, 00, 0, time.UTC)},
				)
				if err != nil {
					t.Error(err)
				}

				_, err = d.InsertOrdersFlow(ctx,
					order.FlowUpdate{ID: 1, OrderID: 1, Status: order.InTheCartStatus, OccurredAt: time.Now().UTC()},
					order.FlowUpdate{ID: 2, OrderID: 1, Status: order.PreparingStatus, OccurredAt: time.Now().UTC()},
					order.FlowUpdate{ID: 3, OrderID: 2, Status: order.InTheCartStatus, OccurredAt: time.Now().UTC()},
					order.FlowUpdate{ID: 4, OrderID: 2, Status: order.PreparingStatus, OccurredAt: time.Now().UTC()},
					order.FlowUpdate{ID: 5, OrderID: 3, Status: order.InTheCartStatus, OccurredAt: time.Now().UTC()},
					order.FlowUpdate{ID: 6, OrderID: 3, Status: order.PendingStatus, OccurredAt: time.Now().UTC()},
					order.FlowUpdate{ID: 7, OrderID: 4, Status: order.InTheCartStatus, OccurredAt: time.Now().UTC()},
					order.FlowUpdate{ID: 8, OrderID: 4, Status: order.WaitingTakeoutStatus, OccurredAt: time.Now().UTC()},
					order.FlowUpdate{ID: 9, OrderID: 5, Status: order.InTheCartStatus, OccurredAt: time.Now().UTC()},
					order.FlowUpdate{ID: 10, OrderID: 5, Status: order.FinishedStatus, OccurredAt: time.Now().UTC()},
				)
				if err != nil {
					t.Error(err)
				}

				_, err = d.InsertItems(ctx,
					item.Item{ID: 1, OrderID: 1, ProductID: 1, Price: 100},
					item.Item{ID: 2, OrderID: 1, ProductID: 2, Price: 101},
					item.Item{ID: 3, OrderID: 2, ProductID: 44, Price: 150},
				)
				if err != nil {
					t.Error(err)
				}
			},
		},
		{
			name: "should return list with delivery mode filtered orders",

			args: args{
				filter: order.Filter{
					DeliveryMode: order.TakeoutDeliveryMode,
				},
			},
			want: []order.Order{
				{ID: 1, DeliveryMode: order.TakeoutDeliveryMode, Status: order.DeliveringStatus, FinalPrice: 101},
			},
			before: func(ctx context.Context, d *tests.Database, t *testing.T) {
				_, err := d.InsertOrders(ctx,
					order.Order{ID: 1, DeliveryMode: order.TakeoutDeliveryMode},
					order.Order{ID: 2, DeliveryMode: order.DeliveryDeliveryMode},
				)
				if err != nil {
					t.Error(err)
				}

				_, err = d.InsertOrdersFlow(ctx,
					order.FlowUpdate{ID: 1, OrderID: 1, Status: order.InTheCartStatus, OccurredAt: time.Now().UTC()},
					order.FlowUpdate{ID: 2, OrderID: 1, Status: order.DeliveringStatus, OccurredAt: time.Now().UTC()},
					order.FlowUpdate{ID: 3, OrderID: 2, Status: order.InTheCartStatus, OccurredAt: time.Now().UTC()},
					order.FlowUpdate{ID: 4, OrderID: 2, Status: order.WaitingDeliveryStatus, OccurredAt: time.Now().UTC()},
				)
				if err != nil {
					t.Error(err)
				}

				_, err = d.InsertItems(ctx,
					item.Item{ID: 1, OrderID: 1, ProductID: 1, Price: 101},
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

			got, err := actions{db: db.Pool}.List(ctx, tt.args.filter)
			assert.ErrorIs(t, err, tt.wantErr)
			assert.Equal(t, tt.want, got)
		})
	}
}
