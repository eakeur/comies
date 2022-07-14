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

func Test_actions_SetObservation(t *testing.T) {
	t.Parallel()

	type args struct {
		id          types.ID
		observation string
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
				id:          1,
				observation: "Remove onions and tomatoes",
			},
			before: func(ctx context.Context, d *tests.Database, t *testing.T) {
				_, err := d.InsertOrders(ctx, order.Order{ID: 1, Observations: "Remove onions"})
				if err != nil {
					t.Error(err)
				}
			},
			after: func(ctx context.Context, d *tests.Database, t *testing.T) {
				d.CheckValue(ctx, "select max(observations) from orders", "Remove onions and tomatoes")
			},
		},
		{
			name: "should fail for nonexistent order",
			args: args{
				id:          1,
				observation: "Remove onions",
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
			err := a.SetObservation(ctx, tt.args.id, tt.args.observation)
			assert.ErrorIs(t, err, tt.wantErr)
		})
	}
}