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

func Test_actions_Create(t *testing.T) {
	t.Parallel()

	var placed = time.Now().UTC()

	type args struct {
		order order.Order
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
			name: "should return created order",
			args: args{
				order: order.Order{
					ID:       1,
					PlacedAt: placed,
				},
			},
			want: order.Order{
				ID:       1,
				PlacedAt: placed,
			},
			after: func(ctx context.Context, db *tests.Database, t *testing.T) {
				db.CheckValue(ctx, "select count(id) from orders", int64(1))
			},
		},
		{
			name: "should return ErrAlreadyExists error for existing ID",
			args: args{
				order: order.Order{
					ID:       1,
					PlacedAt: placed,
				},
			},
			wantErr: fault.ErrAlreadyExists,
			before: func(ctx context.Context, db *tests.Database, t *testing.T) {
				_, err := db.InsertOrders(ctx, order.Order{
					ID:       1,
					PlacedAt: placed,
				})
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

			ctx, db := tests.FetchTestTX(t, tt.before)
			defer db.Drop(tt.after)

			a := actions{}
			got, err := a.Create(ctx, tt.args.order)
			assert.ErrorIs(t, err, tt.wantErr)
			assert.Equalf(t, tt.want, got, "Create(%v)", tt.args.order)
		})
	}
}
