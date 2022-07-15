package item

import (
	"comies/app/core/entities/item"
	"comies/app/core/entities/order"
	"comies/app/gateway/persistence/postgres/tests"
	"comies/app/sdk/fault"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_actions_Create(t *testing.T) {
	t.Parallel()

	type args struct {
		i item.Item
	}
	cases := []struct {
		name    string
		before  tests.Callback
		after   tests.Callback
		args    args
		want    item.Item
		wantErr error
	}{
		{
			name: "should create item successfully",
			args: args{i: item.Item{
				ID:        1,
				OrderID:   1,
				ProductID: 4,
			}},
			want: item.Item{
				ID:        1,
				OrderID:   1,
				ProductID: 4,
			},
			before: func(ctx context.Context, d *tests.Database, t *testing.T) {
				_, err := d.InsertOrders(ctx, order.Order{ID: 1})
				if err != nil {
					t.Error(err)
				}
			},
			after: func(ctx context.Context, d *tests.Database, t *testing.T) {
				var quantityInserted int
				r := d.Pool.QueryRow(ctx, "select count(id) from items")
				if err := r.Scan(&quantityInserted); err != nil && quantityInserted != 1 {
					t.Errorf("the item was not inserted or an error occurred when fetching it: %v", err)
				}
			},
		},
		{
			name: "should fail because the order id is invalid",
			args: args{i: item.Item{
				ID:        1,
				OrderID:   1,
				ProductID: 4,
			}},
			wantErr: fault.ErrNotFound,
		},
		{
			name: "should fail because the item id is duplicate",
			args: args{i: item.Item{
				ID:        1,
				OrderID:   1,
				ProductID: 4,
			}},
			before: func(ctx context.Context, d *tests.Database, t *testing.T) {
				_, err := d.InsertOrders(ctx, order.Order{ID: 1})
				if err != nil {
					t.Error(err)
				}

				_, err = d.InsertItems(ctx,
					item.Item{ID: 1, OrderID: 1, ProductID: 1},
					item.Item{ID: 2, OrderID: 1, ProductID: 2},
					item.Item{ID: 3, OrderID: 1, ProductID: 3},
				)
				if err != nil {
					t.Error(err)
				}
			},
			wantErr: fault.ErrAlreadyExists,
		},
	}
	for _, tt := range cases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctx, db := tests.FetchTestTX(t, tt.before)
			defer db.Drop(tt.after)

			a := actions{}
			got, err := a.Create(ctx, tt.args.i)
			assert.ErrorIs(t, err, tt.wantErr)
			assert.Equalf(t, tt.want, got, "Create(%v)", tt.args.i)
		})
	}
}
