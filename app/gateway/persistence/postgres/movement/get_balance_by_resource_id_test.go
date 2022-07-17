package movement

import (
	"comies/app/core/entities/movement"
	"comies/app/core/entities/stock"
	"comies/app/gateway/persistence/postgres/tests"
	"comies/app/sdk/types"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_actions_GetMovementByResourceID(t *testing.T) {
	t.Parallel()

	var date = time.Date(2001, time.September, 30, 22, 45, 00, 0, time.UTC)

	type args struct {
		resourceID types.ID
		filter     movement.Filter
	}
	cases := []struct {
		name    string
		args    args
		want    types.Quantity
		wantErr error
		before  tests.Callback
		after   tests.Callback
	}{
		{
			name: "should return sum of movements",
			args: args{
				resourceID: 22345666,
			},

			want: 500,
			before: func(ctx context.Context, db *tests.Database, t *testing.T) {
				_, err := db.InsertStocks(ctx, stock.Stock{
					ID:              1,
					TargetID:        22345666,
					MaximumQuantity: 10,
					MinimumQuantity: 100,
					Location:        "Under the table",
				})
				if err != nil {
					t.Error(err)
				}

				_, err = db.InsertMovements(ctx, movement.Movement{
					ID:        1,
					ProductID: 1,
					Type:      movement.OutputMovement,
					Date:      date,
					Quantity:  100,
					PaidValue: 50,
					AgentID:   1544474558856547556,
				}, movement.Movement{
					ID:        2,
					ProductID: 1,
					Type:      movement.ReservedMovement,
					Date:      date,
					Quantity:  100,
					PaidValue: 50,
					AgentID:   1544474558856547556,
				}, movement.Movement{
					ID:        3,
					ProductID: 1,
					Type:      movement.ReservedMovement,
					Date:      date,
					Quantity:  100,
					PaidValue: 50,
					AgentID:   1544474558856547556,
				}, movement.Movement{
					ID:        4,
					ProductID: 1,
					Type:      movement.OutputMovement,
					Date:      date,
					Quantity:  100,
					PaidValue: 50,
					AgentID:   56547556444444444,
				}, movement.Movement{
					ID:        5,
					ProductID: 1,
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
		},
	}
	for _, tt := range cases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctx, db := tests.FetchTestDB(t, tt.before)
			defer db.Drop(tt.after)

			a := actions{db: db.Pool}
			got, err := a.GetBalanceByProductID(ctx, tt.args.resourceID, tt.args.filter)
			assert.ErrorIs(t, err, tt.wantErr)
			assert.Equalf(t, tt.want, got, "GetByID(%v)", tt.args.resourceID)
		})
	}
}
