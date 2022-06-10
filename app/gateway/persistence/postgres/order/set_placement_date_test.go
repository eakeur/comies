package order

import (
	"context"
	"github.com/stretchr/testify/assert"
	"gomies/app/core/entities/order"
	"gomies/app/gateway/persistence/postgres/tests"
	"gomies/app/sdk/fault"
	"gomies/app/sdk/types"
	"testing"
	"time"
)

func Test_actions_SetPlacementDate(t *testing.T) {
	t.Parallel()

	type args struct {
		id   types.ID
		date time.Time
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
				date: time.Date(2001, time.September, 30, 22, 45, 00, 0, time.UTC),
			},
			before: func(ctx context.Context, d *tests.Database, t *testing.T) {
				_, err := d.InsertOrders(ctx, order.Order{ID: 1})
				if err != nil {
					t.Error(err)
				}
			},
		},
		{
			name: "should fail for nonexistent order",
			args: args{
				id:   1,
				date: time.Date(2001, time.September, 30, 22, 45, 00, 0, time.UTC),
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
			err := a.SetPlacementDate(ctx, tt.args.id, tt.args.date)
			assert.ErrorIs(t, err, tt.wantErr)
		})
	}
}
