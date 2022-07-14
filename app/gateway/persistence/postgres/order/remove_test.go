package order

import (
	"context"
	"github.com/stretchr/testify/assert"
	"gomies/app/core/entities/order"
	"gomies/app/gateway/persistence/postgres/tests"
	"gomies/app/sdk/fault"
	"gomies/app/sdk/types"
	"testing"
)

func Test_actions_Remove(t *testing.T) {
	t.Parallel()

	type args struct {
		id types.ID
	}
	cases := []struct {
		name    string
		args    args
		wantErr error
		before  tests.Callback
		after   tests.Callback
	}{
		{
			name: "should delete order successfully",
			args: args{
				id: 1,
			},
			before: func(ctx context.Context, d *tests.Database, t *testing.T) {
				if _, err := d.InsertOrders(ctx, order.Order{ID: 1}); err != nil {
					t.Error(err)
				}
			},
			after: func(ctx context.Context, d *tests.Database, t *testing.T) {
				d.CheckValue(ctx, "select count(id) from orders", int64(0))
			},
		},
		{
			name: "should fail for nonexistent order",
			args: args{
				id: 1,
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
			err := a.Remove(ctx, tt.args.id)
			assert.ErrorIs(t, err, tt.wantErr)
		})
	}
}