package stocking

import (
	"context"
	"gomies/app/core/entities/movement"
	"gomies/app/core/entities/stock"
	"gomies/app/sdk/fault"
	"gomies/app/sdk/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWorkflow_ReserveResources(t *testing.T) {
	t.Parallel()

	type (
		args struct {
			reservationID types.ID
			reservations  []Reservation
		}

		opts struct {
			movements *movement.ActionsMock
			stocks    *stock.ActionsMock
		}

		test struct {
			name    string
			args    args
			want    []ReservationResult
			wantErr error
			opts    opts
		}
	)

	for _, c := range []test{
		{
			name: "should reserve all resources",
			args: args{
				reservationID: 444,
				reservations: []Reservation{
					{
						ResourceID: 1,
						Quantity:   100,
					},
					{
						ResourceID: 2,
						Quantity:   50,
					},
				},
			},
			want: []ReservationResult{
				{
					ResourceID: 1,
					Want:       100,
					Got:        300,
				},
				{
					ResourceID: 2,
					Want:       50,
					Got:        5000,
				},
			},
			opts: opts{
				movements: &movement.ActionsMock{
					GetBalanceByResourceIDFunc: func(ctx context.Context, resourceID types.ID, filter movement.Filter) (types.Quantity, error) {
						if filter.ResourceID == 1 {
							return 300, nil
						}

						return 5000, nil
					},

					CreateFunc: func(ctx context.Context, movement movement.Movement) (movement.Movement, error) {
						return movement, nil
					},
				},
				stocks: &stock.ActionsMock{
					GetStockByIDFunc: func(ctx context.Context, resourceID types.ID) (stock.Stock, error) {
						return stock.Stock{
							MaximumQuantity: 5000,
							MinimumQuantity: 0,
						}, nil
					},
				},
			},
		},
		{
			name: "should reserve all resources available and return failed reservations",
			args: args{
				reservationID: 444,
				reservations: []Reservation{
					{
						ResourceID: 1,
						Quantity:   100,
					},
					{
						ResourceID: 2,
						Quantity:   50,
					},
				},
			},
			want: []ReservationResult{
				{
					ResourceID: 1,
					Want:       100,
					Got:        300,
				},
				{
					ResourceID: 2,
					Want:       50,
					Got:        20,
					Error:      stock.ErrStockEmpty,
				},
			},
			opts: opts{
				movements: &movement.ActionsMock{
					GetBalanceByResourceIDFunc: func(ctx context.Context, resourceID types.ID, filter movement.Filter) (types.Quantity, error) {
						if filter.ResourceID == 1 {
							return 300, nil
						}

						return 20, nil
					},

					CreateFunc: func(ctx context.Context, movement movement.Movement) (movement.Movement, error) {
						return movement, nil
					},
				},
				stocks: &stock.ActionsMock{
					GetStockByIDFunc: func(ctx context.Context, resourceID types.ID) (stock.Stock, error) {
						return stock.Stock{
							MaximumQuantity: 5000,
							MinimumQuantity: 0,
						}, nil
					},
				},
			},
		},
		{
			name: "should return error for stock not found",
			args: args{
				reservationID: 444,
				reservations: []Reservation{
					{
						ResourceID: 1,
						Quantity:   100,
					},
					{
						ResourceID: 2,
						Quantity:   50,
					},
				},
			},
			wantErr: fault.ErrNotFound,
			opts: opts{
				movements: &movement.ActionsMock{
					GetBalanceByResourceIDFunc: func(ctx context.Context, resourceID types.ID, filter movement.Filter) (types.Quantity, error) {
						if filter.ResourceID == 1 {
							return 300, nil
						}

						return 20, nil
					},
					CreateFunc: func(ctx context.Context, movement movement.Movement) (movement.Movement, error) {
						return movement, nil
					},
				},
				stocks: &stock.ActionsMock{
					GetStockByIDFunc: func(ctx context.Context, resourceID types.ID) (stock.Stock, error) {
						return stock.Stock{}, fault.ErrNotFound
					},
				},
			},
		},
		{
			name: "should return ordered result for multiple reservations",
			args: args{
				reservationID: 444,
				reservations: func() []Reservation {
					arr := make([]Reservation, 1000000)
					for i := 0; i < len(arr); i++ {
						arr[i] = Reservation{
							ResourceID: types.ID(i + 1),
							Quantity:   10,
						}
					}

					return arr
				}(),
			},
			want: func() []ReservationResult {
				arr := make([]ReservationResult, 1000000)
				for i := 0; i < len(arr); i++ {
					arr[i] = ReservationResult{
						ResourceID: types.ID(i + 1),
						Want:       10,
						Got:        100,
					}
				}

				return arr
			}(),
			opts: opts{
				movements: &movement.ActionsMock{
					GetBalanceByResourceIDFunc: func(ctx context.Context, resourceID types.ID, filter movement.Filter) (types.Quantity, error) {
						return 100, nil
					},
					CreateFunc: func(ctx context.Context, movement movement.Movement) (movement.Movement, error) {
						return movement, nil
					},
				},
				stocks: &stock.ActionsMock{
					GetStockByIDFunc: func(ctx context.Context, resourceID types.ID) (stock.Stock, error) {
						return stock.Stock{}, nil
					},
				},
			},
		},
	} {
		c := c

		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			got, gotErr := workflow{stocks: c.opts.stocks, movements: c.opts.movements}.
				ReserveResources(context.Background(), c.args.reservationID, c.args.reservations)

			assert.ErrorIs(t, gotErr, c.wantErr)
			assert.Equal(t, c.want, got)
		})
	}
}
