package stock

import (
	"context"
	"github.com/stretchr/testify/assert"
	"gomies/app/core/entities/stock"
	"gomies/app/sdk/fault"
	"gomies/app/sdk/types"
	"testing"
)

func TestWorkflow_ReserveResources(t *testing.T) {
	t.Parallel()

	type (
		args struct {
			reservationID types.ID
			reservations  []Reservation
		}

		opts struct {
			stocks *stock.ActionsMock
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
				stocks: &stock.ActionsMock{
					ComputeStockFunc: func(ctx context.Context, filter stock.Filter) (types.Quantity, error) {
						if filter.ResourceID == 1 {
							return 300, nil
						}

						return 5000, nil
					},

					GetStockByIDFunc: func(ctx context.Context, resourceID types.ID) (stock.Stock, error) {
						return stock.Stock{
							MaximumQuantity: 5000,
							MinimumQuantity: 0,
						}, nil
					},

					SaveMovementsFunc: func(ctx context.Context, movement stock.Movement) (stock.Movement, error) {
						return movement, nil
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
				stocks: &stock.ActionsMock{
					ComputeStockFunc: func(ctx context.Context, filter stock.Filter) (types.Quantity, error) {
						if filter.ResourceID == 1 {
							return 300, nil
						}

						return 20, nil
					},

					GetStockByIDFunc: func(ctx context.Context, resourceID types.ID) (stock.Stock, error) {
						return stock.Stock{
							MaximumQuantity: 5000,
							MinimumQuantity: 0,
						}, nil
					},

					SaveMovementsFunc: func(ctx context.Context, movement stock.Movement) (stock.Movement, error) {
						return movement, nil
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
				stocks: &stock.ActionsMock{
					ComputeStockFunc: func(ctx context.Context, filter stock.Filter) (types.Quantity, error) {
						if filter.ResourceID == 1 {
							return 300, nil
						}

						return 20, nil
					},

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
				stocks: &stock.ActionsMock{
					ComputeStockFunc: func(ctx context.Context, filter stock.Filter) (types.Quantity, error) {
						return 100, nil
					},

					GetStockByIDFunc: func(ctx context.Context, resourceID types.ID) (stock.Stock, error) {
						return stock.Stock{}, nil
					},

					SaveMovementsFunc: func(ctx context.Context, movement stock.Movement) (stock.Movement, error) {
						return movement, nil
					},
				},
			},
		},
	} {
		c := c

		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			got, gotErr := workflow{stocks: c.opts.stocks}.
				ReserveResources(context.Background(), c.args.reservationID, c.args.reservations)

			assert.ErrorIs(t, gotErr, c.wantErr)
			assert.Equal(t, c.want, got)
		})
	}
}
