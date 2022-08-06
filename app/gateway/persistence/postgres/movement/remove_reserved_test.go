package movement

import (
	"comies/app/core/entities/movement"
	"comies/app/core/entities/product"
	"comies/app/core/throw"
	"comies/app/core/types"
	"comies/app/gateway/persistence/postgres/tests"
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
				_, err := d.InsertProducts(ctx, product.Product{
					ID:   1,
					Code: "A",
					Stock: product.Stock{
						MaximumQuantity: 10,
						MinimumQuantity: 100,
						Location:        "Under the table",
					},
				}, product.Product{
					ID:   2,
					Code: "B",
					Stock: product.Stock{
						MaximumQuantity: 10,
						MinimumQuantity: 100,
						Location:        "Under the table",
					},
				}, product.Product{
					ID:   3,
					Code: "C",
					Stock: product.Stock{
						MaximumQuantity: 10,
						MinimumQuantity: 100,
						Location:        "Under the table",
					},
				}, product.Product{
					ID:   4,
					Code: "D",
					Stock: product.Stock{
						MaximumQuantity: 10,
						MinimumQuantity: 100,
						Location:        "Under the table",
					},
				}, product.Product{
					ID:   5,
					Code: "E",
					Stock: product.Stock{
						MaximumQuantity: 10,
						MinimumQuantity: 100,
						Location:        "Under the table",
					},
				})
				if err != nil {
					t.Error(err)
				}

				_, err = d.InsertMovements(ctx, movement.Movement{
					ID:        1,
					ProductID: 5,
					Type:      movement.OutputType,
					Date:      date,
					Quantity:  100,
					PaidValue: 50,
					AgentID:   1544474558856547556,
				}, movement.Movement{
					ID:        2,
					ProductID: 4,
					Type:      movement.ReservedType,
					Date:      date,
					Quantity:  100,
					PaidValue: 50,
					AgentID:   1544474558856547556,
				}, movement.Movement{
					ID:        3,
					ProductID: 5,
					Type:      movement.ReservedType,
					Date:      date,
					Quantity:  100,
					PaidValue: 50,
					AgentID:   1544474558856547556,
				}, movement.Movement{
					ID:        4,
					ProductID: 3,
					Type:      movement.OutputType,
					Date:      date,
					Quantity:  100,
					PaidValue: 50,
					AgentID:   56547556444444444,
				}, movement.Movement{
					ID:        5,
					ProductID: 2,
					Type:      movement.OutputType,
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
			err := a.RemoveReserved(ctx, tt.args.agentID)
			assert.ErrorIs(t, err, tt.wantErr)
		})
	}
}
