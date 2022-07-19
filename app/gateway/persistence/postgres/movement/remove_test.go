package movement

import (
	"comies/app/core/entities/movement"
	"comies/app/core/entities/product"
	"comies/app/gateway/persistence/postgres/tests"
	"comies/app/sdk/throw"
	"comies/app/sdk/types"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_actions_Remove(t *testing.T) {
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
			name: "should delete movement successfully",
			args: args{
				movementID: 1,
			},
			before: func(ctx context.Context, d *tests.Database, t *testing.T) {
				_, err := d.InsertProducts(ctx, product.Product{
					ID: 1,
					Stock: product.Stock{
						MaximumQuantity: 10,
						MinimumQuantity: 100,
						Location:        "Under the table",
					},
				})

				_, err = d.InsertMovements(ctx, movement.Movement{
					ID:        1,
					ProductID: 1,
					Type:      movement.OutputType,
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
				d.CheckValue(ctx, "select count(id) from movements", int64(0))
			},
		},
		{
			name: "should fail for nonexistent movement",
			args: args{
				movementID: 1,
			},
			wantErr: throw.ErrNotFound,
		},
	}
	for _, tt := range cases {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctx, db := tests.FetchTestTX(t, tt.before)
			defer db.Drop(tt.after)

			a := actions{}
			err := a.Remove(ctx, tt.args.movementID)
			assert.ErrorIs(t, err, tt.wantErr)
		})
	}
}
