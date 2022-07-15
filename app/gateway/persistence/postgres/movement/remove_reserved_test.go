package movement

import (
	"comies/app/core/entities/movement"
	"comies/app/core/entities/stock"
	"comies/app/gateway/persistence/postgres/tests"
	"comies/app/sdk/fault"
	"comies/app/sdk/types"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_actions_RemoveReserved(t *testing.T) {
	t.Parallel()

	var date = time.Now().UTC()

	type args struct {
		agentID types.ID
	}
	cases := []struct {
		name    string
		args    args
		wantErr error
		before  tests.Callback
		after   tests.Callback
	}{
		{
			name: "should delete movements successfully",
			args: args{
				agentID: 1544474558856547556,
			},
			before: func(ctx context.Context, d *tests.Database, t *testing.T) {
				_, err := d.InsertStocks(ctx, stock.Stock{
					ID:              1,
					TargetID:        22345666,
					MaximumQuantity: 10,
					MinimumQuantity: 100,
					Location:        "Under the table",
				}, stock.Stock{
					ID:              2,
					TargetID:        765434,
					MaximumQuantity: 10,
					MinimumQuantity: 100,
					Location:        "Under the table",
				}, stock.Stock{
					ID:              3,
					TargetID:        223345,
					MaximumQuantity: 10,
					MinimumQuantity: 100,
					Location:        "Under the table",
				}, stock.Stock{
					ID:              4,
					TargetID:        3232323,
					MaximumQuantity: 10,
					MinimumQuantity: 100,
					Location:        "Under the table",
				}, stock.Stock{
					ID:              5,
					TargetID:        334455,
					MaximumQuantity: 10,
					MinimumQuantity: 100,
					Location:        "Under the table",
				})
				if err != nil {
					t.Error(err)
				}

				_, err = d.InsertMovements(ctx, movement.Movement{
					ID:        1,
					StockID:   5,
					Type:      movement.OutputMovement,
					Date:      date,
					Quantity:  100,
					PaidValue: 50,
					AgentID:   1544474558856547556,
				}, movement.Movement{
					ID:        2,
					StockID:   4,
					Type:      movement.ReservedMovement,
					Date:      date,
					Quantity:  100,
					PaidValue: 50,
					AgentID:   1544474558856547556,
				}, movement.Movement{
					ID:        3,
					StockID:   5,
					Type:      movement.ReservedMovement,
					Date:      date,
					Quantity:  100,
					PaidValue: 50,
					AgentID:   1544474558856547556,
				}, movement.Movement{
					ID:        4,
					StockID:   3,
					Type:      movement.OutputMovement,
					Date:      date,
					Quantity:  100,
					PaidValue: 50,
					AgentID:   56547556444444444,
				}, movement.Movement{
					ID:        5,
					StockID:   2,
					Type:      movement.OutputMovement,
					Date:      date,
					Quantity:  100,
					PaidValue: 50,
					AgentID:   547556444444444,
				})
				if err != nil {
					t.Error(err)
				}
			},
			after: func(ctx context.Context, d *tests.Database, t *testing.T) {
				d.CheckValue(ctx, "select count(id) from movements", int64(3))
			},
		},
		{
			name: "should fail for nonexistent movement",
			args: args{
				agentID: 1,
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
			err := a.RemoveReserved(ctx, tt.args.agentID)
			assert.ErrorIs(t, err, tt.wantErr)
		})
	}
}
