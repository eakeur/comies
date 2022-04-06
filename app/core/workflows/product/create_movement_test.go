package product

import (
	"context"
	"errors"
	"gomies/app/core/entities/catalog/product"
	"gomies/pkg/sdk/fault"
	"gomies/pkg/sdk/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWorkflow_AddMovement(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	fakeID := types.UID("1bdcafba-9deb-48b4-8a0e-ecea4c99b0e3")
	fakeErr := errors.New("fakeErr")

	type (
		args struct {
			productID types.UID
			movement  Movement
			config    product.Stock
		}

		opts struct {
			products *product.ActionsMock
			stocks   *StockServiceMock
		}

		test struct {
			name    string
			args    args
			opts    opts
			want    types.Quantity
			wantErr error
		}
	)

	cases := []test{
		{
			name: "should return remaining quantity for stock created",
			args: args{
				productID: fakeID,
				movement:  Movement{},
			},
			want: 200000,
			opts: opts{
				stocks: &StockServiceMock{
					CreateMovementFunc: func(ctx context.Context, config product.Stock, resourceID types.UID, movement Movement) (types.Quantity, error) {
						return 200000, nil
					},
				},
				products: &product.ActionsMock{
					GetProductStockInfoFunc: func(ctx context.Context, key product.Key) (product.Stock, error) {
						return product.Stock{
							MaximumQuantity: 1000,
							MinimumQuantity: 50,
						}, nil
					},
				},
			},
		},
		{
			name: "should fail because product does not exist",
			args: args{
				productID: fakeID,
				movement:  Movement{},
			},
			wantErr: fault.ErrNotFound,
			opts: opts{
				stocks: &StockServiceMock{
					CreateMovementFunc: func(ctx context.Context, config product.Stock, resourceID types.UID, movement Movement) (types.Quantity, error) {
						return 200000, nil
					},
				},
				products: &product.ActionsMock{
					GetProductStockInfoFunc: func(ctx context.Context, key product.Key) (product.Stock, error) {
						return product.Stock{}, fault.ErrNotFound
					},
				},
			},
		},
		{
			name: "should fail because service failed",
			args: args{
				productID: fakeID,
				movement:  Movement{},
			},
			wantErr: fakeErr,
			opts: opts{
				stocks: &StockServiceMock{
					CreateMovementFunc: func(ctx context.Context, config product.Stock, resourceID types.UID, movement Movement) (types.Quantity, error) {
						return 0, fakeErr
					},
				},
				products: &product.ActionsMock{
					GetProductStockInfoFunc: func(ctx context.Context, key product.Key) (product.Stock, error) {
						return product.Stock{}, nil
					},
				},
			},
		},
	}

	for _, c := range cases {
		c := c

		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			wf := workflow{
				products: c.opts.products,
				stocks:   c.opts.stocks,
			}

			ingredient, err := wf.CreateMovement(ctx, c.args.productID, c.args.movement)

			assert.ErrorIs(t, err, c.wantErr)
			assert.Equal(t, c.want, ingredient)

		})
	}

}
