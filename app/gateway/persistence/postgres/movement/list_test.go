package movement

import (
	"comies/app/core/entities/movement"
	"comies/app/core/entities/product"
	"comies/app/gateway/persistence/postgres/tests"
	"comies/app/sdk/types"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_actions_ListByResourceID(t *testing.T) {
	t.Parallel()

	type args struct {
		resourceID types.ID
		filter     movement.Filter
	}
	cases := []struct {
		name    string
		args    args
		before  tests.Callback
		want    []movement.Movement
		wantErr error
	}{
		{
			name: "should return empty list of orders",
			args: args{
				resourceID: 2000,
			},
			want: []movement.Movement{},
		},
		{
			name: "should return list with date filtered movements",

			args: args{
				resourceID: 1,
				filter: movement.Filter{
					InitialDate: time.Date(2020, 01, 01, 00, 00, 00, 00, time.UTC),
					FinalDate:   time.Date(2020, 01, 02, 00, 00, 00, 00, time.UTC),
				},
			},
			want: []movement.Movement{
				{
					ID:        3,
					ProductID: 1,
					Type:      movement.ReservedType,
					Date:      time.Date(2020, 01, 01, 00, 00, 00, 00, time.UTC),
					Quantity:  100,
					PaidValue: 50,
					AgentID:   1544474558856547556,
				}, {
					ID:        4,
					ProductID: 1,
					Type:      movement.OutputType,
					Date:      time.Date(2020, 01, 01, 15, 00, 00, 00, time.UTC),
					Quantity:  100,
					PaidValue: 50,
					AgentID:   56547556444444444,
				},
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
				if err != nil {
					t.Error(err)
				}

				_, err = d.InsertMovements(ctx, movement.Movement{
					ID:        1,
					ProductID: 1,
					Type:      movement.OutputType,
					Date:      time.Date(2019, 12, 31, 00, 00, 00, 00, time.UTC),
					Quantity:  100,
					PaidValue: 50,
					AgentID:   1544474558856547556,
				}, movement.Movement{
					ID:        2,
					ProductID: 1,
					Type:      movement.ReservedType,
					Date:      time.Date(2019, 12, 31, 15, 00, 00, 00, time.UTC),
					Quantity:  100,
					PaidValue: 50,
					AgentID:   1544474558856547556,
				}, movement.Movement{
					ID:        3,
					ProductID: 1,
					Type:      movement.ReservedType,
					Date:      time.Date(2020, 01, 01, 00, 00, 00, 00, time.UTC),
					Quantity:  100,
					PaidValue: 50,
					AgentID:   1544474558856547556,
				}, movement.Movement{
					ID:        4,
					ProductID: 1,
					Type:      movement.OutputType,
					Date:      time.Date(2020, 01, 01, 15, 00, 00, 00, time.UTC),
					Quantity:  100,
					PaidValue: 50,
					AgentID:   56547556444444444,
				}, movement.Movement{
					ID:        5,
					ProductID: 1,
					Type:      movement.OutputType,
					Date:      time.Date(2020, 01, 02, 01, 00, 00, 00, time.UTC),
					Quantity:  100,
					PaidValue: 50,
					AgentID:   547556444444444,
				})
				if err != nil {
					t.Error(err)
				}
			},
		},
		{
			name: "should return list with no filters",

			args: args{
				resourceID: 1,
			},
			want: []movement.Movement{
				{
					ID:        1,
					ProductID: 1,
					Type:      movement.OutputType,
					Date:      time.Date(2019, 12, 31, 00, 00, 00, 00, time.UTC),
					Quantity:  100,
					PaidValue: 50,
					AgentID:   1544474558856547556,
				}, {
					ID:        2,
					ProductID: 1,
					Type:      movement.ReservedType,
					Date:      time.Date(2019, 12, 31, 15, 00, 00, 00, time.UTC),
					Quantity:  100,
					PaidValue: 50,
					AgentID:   1544474558856547556,
				}, {
					ID:        3,
					ProductID: 1,
					Type:      movement.ReservedType,
					Date:      time.Date(2020, 01, 01, 00, 00, 00, 00, time.UTC),
					Quantity:  100,
					PaidValue: 50,
					AgentID:   1544474558856547556,
				}, {
					ID:        4,
					ProductID: 1,
					Type:      movement.OutputType,
					Date:      time.Date(2020, 01, 01, 15, 00, 00, 00, time.UTC),
					Quantity:  100,
					PaidValue: 50,
					AgentID:   56547556444444444,
				}, {
					ID:        5,
					ProductID: 1,
					Type:      movement.OutputType,
					Date:      time.Date(2020, 01, 02, 01, 00, 00, 00, time.UTC),
					Quantity:  100,
					PaidValue: 50,
					AgentID:   547556444444444,
				},
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
				if err != nil {
					t.Error(err)
				}

				_, err = d.InsertMovements(ctx, movement.Movement{
					ID:        1,
					ProductID: 1,
					Type:      movement.OutputType,
					Date:      time.Date(2019, 12, 31, 00, 00, 00, 00, time.UTC),
					Quantity:  100,
					PaidValue: 50,
					AgentID:   1544474558856547556,
				}, movement.Movement{
					ID:        2,
					ProductID: 1,
					Type:      movement.ReservedType,
					Date:      time.Date(2019, 12, 31, 15, 00, 00, 00, time.UTC),
					Quantity:  100,
					PaidValue: 50,
					AgentID:   1544474558856547556,
				}, movement.Movement{
					ID:        3,
					ProductID: 1,
					Type:      movement.ReservedType,
					Date:      time.Date(2020, 01, 01, 00, 00, 00, 00, time.UTC),
					Quantity:  100,
					PaidValue: 50,
					AgentID:   1544474558856547556,
				}, movement.Movement{
					ID:        4,
					ProductID: 1,
					Type:      movement.OutputType,
					Date:      time.Date(2020, 01, 01, 15, 00, 00, 00, time.UTC),
					Quantity:  100,
					PaidValue: 50,
					AgentID:   56547556444444444,
				}, movement.Movement{
					ID:        5,
					ProductID: 1,
					Type:      movement.OutputType,
					Date:      time.Date(2020, 01, 02, 01, 00, 00, 00, time.UTC),
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
			defer db.Drop()

			got, err := actions{db: db.Pool}.ListByProductID(ctx, tt.args.resourceID, tt.args.filter)
			assert.ErrorIs(t, err, tt.wantErr)
			assert.Equal(t, tt.want, got)
		})
	}
}
