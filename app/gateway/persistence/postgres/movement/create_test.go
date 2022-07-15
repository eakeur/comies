package movement

import (
	"comies/app/core/entities/movement"
	"comies/app/core/entities/stock"
	"comies/app/gateway/persistence/postgres/tests"
	"comies/app/sdk/fault"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_actions_Create(t *testing.T) {
	t.Parallel()

	var date = time.Now().UTC()

	type args struct {
		movement movement.Movement
	}
	cases := []struct {
		name    string
		args    args
		want    movement.Movement
		wantErr error
		before  tests.Callback
		after   tests.Callback
	}{
		{
			name: "should return created movement",
			args: args{
				movement: movement.Movement{
					ID:        1,
					StockID:   1,
					Type:      movement.OutputMovement,
					Date:      date,
					Quantity:  100,
					PaidValue: 50,
					AgentID:   1544474558856547556,
				},
			},
			want: movement.Movement{
				ID:        1,
				StockID:   1,
				Type:      movement.OutputMovement,
				Date:      date,
				Quantity:  100,
				PaidValue: 50,
				AgentID:   1544474558856547556,
			},
			before: func(ctx context.Context, db *tests.Database, t *testing.T) {
				_, err := db.InsertStocks(ctx, stock.Stock{
					ID:              1,
					TargetID:        1,
					MaximumQuantity: 10,
					MinimumQuantity: 100,
					Location:        "Under the table",
				})
				if err != nil {
					t.Error(err)
				}
			},
			after: func(ctx context.Context, db *tests.Database, t *testing.T) {
				const script = `
					select 
						id = $1 and 
						stock_id = $2 and 
						type = $3 and 
						date = $4 and 
						quantity = $5 and 
						value = $6 and 
						agent_id = $7 
					as equal
					from movements where id = $1
				`
				db.CheckValue(ctx, script, true, 1, 1, movement.OutputMovement, date, 100, 50, 1544474558856547556)
			},
		},
		{
			name: "should fail for nonexistent stock",
			args: args{
				movement: movement.Movement{
					ID:        1,
					StockID:   1,
					Type:      movement.OutputMovement,
					Date:      date,
					Quantity:  100,
					PaidValue: 50,
					AgentID:   1544474558856547556,
				},
			},
			wantErr: fault.ErrNotFound,
		},
		{
			name: "should fail for repeated movement id",
			args: args{
				movement: movement.Movement{
					ID:        1,
					StockID:   1,
					Type:      movement.OutputMovement,
					Date:      date,
					Quantity:  100,
					PaidValue: 50,
					AgentID:   1544474558856547556,
				},
			},
			wantErr: fault.ErrAlreadyExists,
			before: func(ctx context.Context, db *tests.Database, t *testing.T) {
				_, err := db.InsertStocks(ctx, stock.Stock{
					ID:              1,
					TargetID:        1,
					MaximumQuantity: 10,
					MinimumQuantity: 100,
					Location:        "Under the table",
				})
				if err != nil {
					t.Error(err)
				}

				_, err = db.InsertMovements(ctx, movement.Movement{
					ID:        1,
					StockID:   1,
					Type:      movement.OutputMovement,
					Date:      date,
					Quantity:  100,
					PaidValue: 50,
					AgentID:   1544474558856547556,
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
			got, err := a.Create(ctx, tt.args.movement)
			assert.ErrorIs(t, err, tt.wantErr)
			assert.Equalf(t, tt.want, got, "Create(%v)", tt.args.movement)
		})
	}
}