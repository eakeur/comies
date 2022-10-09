package order

import (
	"comies/app/core/entities/order"
	"comies/app/core/types"
	"comies/app/gateway/persistence/postgres/tests"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_actions_SetDeliveryMode(t *testing.T) {
	t.Parallel()

	type args struct {
		id   types.ID
		mode order.DeliveryMode
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
				id:   1,
				mode: order.DeliveryDeliveryMode,
			},
			before: func(ctx context.Context, d *tests.Database, t *testing.T) {
				_, err := d.InsertOrders(ctx, order.Order{ID: 1, DeliveryMode: order.TakeoutDeliveryMode})
				if err != nil {
					t.Error(err)
				}
			},
			after: func(ctx context.Context, d *tests.Database, _ *testing.T) {
				d.CheckValue(t, ctx, "select max(delivery_mode) from orders", order.DeliveryDeliveryMode)
			},
		},
		{
			name: "should fail for nonexistent order",
			args: args{
				id:   1,
				mode: order.DeliveryDeliveryMode,
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
			err := a.SetDeliveryMode(ctx, tt.args.id, tt.args.mode)
			assert.ErrorIs(t, err, tt.wantErr)
		})
	}
}
