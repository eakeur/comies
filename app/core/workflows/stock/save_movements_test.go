package stock

import (
	"context"
	"gomies/app/core/entities/stock"
	"gomies/pkg/sdk/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWorkflow_AddToStock(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	fakeID := types.UID("1bdcafba-9deb-48b4-8a0e-ecea4c99b0e3")

	type (
		args struct {
			resourceID types.UID
			movement   stock.Movement
			config     stock.Config
		}

		opts struct {
			stocks *stock.ActionsMock
		}

		test struct {
			name    string
			args    args
			opts    opts
			want    stock.AdditionResult
			wantErr error
		}
	)

	cases := []test{
		{
			name: "should return movement created",
			args: args{
				resourceID: fakeID,
				movement: stock.Movement{
					TargetID: fakeID,
					Type:     stock.OutputMovement,
					Quantity: 10,
				},
				config: stock.Config{
					MaximumQuantity: 100,
				},
			},
			want: stock.AdditionResult{Count: 40},
			opts: opts{
				stocks: &stock.ActionsMock{
					ComputeFunc: func(ctx context.Context, filter stock.Filter) (types.Quantity, error) {
						return 50, nil
					},
					SaveMovementsFunc: func(ctx context.Context, movement ...stock.Movement) ([]stock.Movement, error) {
						return movement, nil
					},
				},
			},
		},
		{
			name: "should return input movement created with full stock",
			args: args{
				resourceID: fakeID,
				movement: stock.Movement{
					TargetID: fakeID,
					Type:     stock.InputMovement,
					Quantity: 10,
				},
				config: stock.Config{
					MaximumQuantity: 100,
				},
			},
			want: stock.AdditionResult{Count: 100},
			opts: opts{
				stocks: &stock.ActionsMock{
					ComputeFunc: func(ctx context.Context, filter stock.Filter) (types.Quantity, error) {
						return 90, nil
					},
					SaveMovementsFunc: func(ctx context.Context, movement ...stock.Movement) ([]stock.Movement, error) {
						return movement, nil
					},
				},
			},
		},
		{
			name: "should return output movement created with full stock",
			args: args{
				resourceID: fakeID,
				movement: stock.Movement{
					TargetID: fakeID,
					Type:     stock.OutputMovement,
					Quantity: 10,
				},
				config: stock.Config{
					MaximumQuantity: 100,
				},
			},
			want: stock.AdditionResult{Count: 90},
			opts: opts{
				stocks: &stock.ActionsMock{
					ComputeFunc: func(ctx context.Context, filter stock.Filter) (types.Quantity, error) {
						return 100, nil
					},
					SaveMovementsFunc: func(ctx context.Context, movement ...stock.Movement) ([]stock.Movement, error) {
						return movement, nil
					},
				},
			},
		},
		{
			name: "should fail because stock is already full",
			args: args{
				resourceID: fakeID,
				movement: stock.Movement{
					TargetID: fakeID,
					Type:     stock.InputMovement,
					Quantity: 10,
				},
				config: stock.Config{
					MaximumQuantity: 100,
				},
			},
			wantErr: stock.ErrStockFull,
			opts: opts{
				stocks: &stock.ActionsMock{
					ComputeFunc: func(ctx context.Context, filter stock.Filter) (types.Quantity, error) {
						return 100, nil
					},
					SaveMovementsFunc: func(ctx context.Context, movement ...stock.Movement) ([]stock.Movement, error) {
						return movement, nil
					},
				},
			},
		},
		{
			name: "should fail because stock would be lower than allowed",
			args: args{
				resourceID: fakeID,
				movement: stock.Movement{
					TargetID: fakeID,
					Type:     stock.OutputMovement,
					Quantity: 10,
				},
				config: stock.Config{
					MaximumQuantity: 100,
					MinimumQuantity: 15,
				},
			},
			wantErr: stock.ErrStockEmpty,
			opts: opts{
				stocks: &stock.ActionsMock{
					ComputeFunc: func(ctx context.Context, filter stock.Filter) (types.Quantity, error) {
						return 20, nil
					},
					SaveMovementsFunc: func(ctx context.Context, movement ...stock.Movement) ([]stock.Movement, error) {
						return movement, nil
					},
				},
			},
		},
	}

	for _, c := range cases {
		c := c

		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			wf := NewWorkflow(c.opts.stocks)
			stk, err := wf.SaveMovements(ctx, c.args.config, c.args.resourceID, c.args.movement)

			assert.ErrorIs(t, err, c.wantErr)
			assert.Equal(t, c.want, stk)

		})
	}
}
