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
			wantErr: types.ErrNotFound,
		},
	}
	for _, tt := range cases {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctx, _ := tests.FetchTestTX(t, tt.before, tt.after)

			a := actions{}
			err := a.SetPlacementDate(ctx, tt.args.id, tt.args.date)
			assert.ErrorIs(t, err, tt.wantErr)
		})
	}
}
