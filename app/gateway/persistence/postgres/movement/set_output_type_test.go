package movement

import (
	"context"
	"github.com/stretchr/testify/assert"
	"gomies/app/core/entities/movement"
	"gomies/app/core/entities/stock"
	"gomies/app/gateway/persistence/postgres/tests"
	"gomies/app/sdk/fault"
	"gomies/app/sdk/types"
	"testing"
	"time"
)

func Test_actions_SetOutputType(t *testing.T) {
	t.Parallel()

	var date = time.Now().UTC()

	type args struct {
		movementID types.ID
	}
	cases := []struct {
		name    string
		args    args
		wantErr error
		before  tests.Callback
		after   tests.Callback
	}{
		{
			name: "should update movement successfully",
			args: args{
				movementID: 1,
			},
			before: func(ctx context.Context, d *tests.Database, t *testing.T) {
				_, err := d.InsertStocks(ctx, stock.Stock{
					ID:              1,
					TargetID:        1,
					MaximumQuantity: 10,
					MinimumQuantity: 100,
					Location:        "Under the table",
				})
				if err != nil {
					t.Error(err)
				}

				_, err = d.InsertMovements(ctx, movement.Movement{
					ID:        1,
					StockID:   1,
					Type:      movement.ReservedMovement,
					Date:      date,
					Quantity:  100,
					PaidValue: 50,
					AgentID:   1544474558856547556,
				})
				if err != nil {
					t.Error(err)
				}
			},
			after: func(ctx context.Context, d *tests.Database, t *testing.T) {
				d.CheckValue(ctx, "select type from movements where id = $1", movement.OutputMovement, 1)
			},
		},
		{
			name: "should fail for nonexistent movement",
			args: args{
				movementID: 1,
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
			err := a.SetOutputType(ctx, tt.args.movementID)
			assert.ErrorIs(t, err, tt.wantErr)
		})
	}
}
