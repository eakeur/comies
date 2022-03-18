package product

import (
	"context"
	"github.com/stretchr/testify/assert"
	"gomies/pkg/catalog/core/entities/product"
	"gomies/pkg/sdk/tests"
	"gomies/pkg/sdk/types"
	"gomies/pkg/stocking/core/entities/stock"
	"testing"
)

func TestWorkflow_AddToStock(t *testing.T) {
	const operation = "Workflows.Product.AddToStock"
	t.Parallel()

	ctx := tests.WorkflowContext(idExample1, idExample2)
	managers := tests.Managers()

	type (
		args struct {
			key      product.Key
			movement stock.Movement
		}

		opts struct {
			products *product.ActionsMock
			stocks   *stock.ActionsMock
		}

		test struct {
			name    string
			args    args
			opts    opts
			want    product.StockAdditionResult
			wantKey product.Key
			wantErr error
		}
	)

	cases := []test{
		{
			name: "should return movement created",
			args: args{
				key: product.Key{ID: idExample1},
				movement: stock.Movement{
					TargetID: idExample1,
					Type:     stock.Output,
					Quantity: 10,
				},
			},
			want: product.StockAdditionResult{
				Movement: stock.Movement{
					Entity: types.Entity{
						History: types.History{
							By:        idExample1,
							Operation: operation,
						},
					},
					TargetID: idExample1,
					Type:     stock.Output,
					Quantity: 10,
				},
				RemainingStock: 40,
			},
			wantKey: product.Key{ID: idExample1, Store: types.Store{StoreID: idExample2}},
			opts: opts{
				products: &product.ActionsMock{
					GetProductStockInfoFunc: func(ctx context.Context, key product.Key) (product.Stock, error) {
						return product.Stock{
							CostPrice:       750,
							MaximumQuantity: 100,
							MinimumQuantity: 0,
						}, nil
					},
				},
				stocks: &stock.ActionsMock{
					ComputeStockFunc: func(ctx context.Context, target types.External, stockFilter stock.Filter) (stock.Actual, error) {
						return stock.Actual{
							TargetID:    target,
							InitialDate: stockFilter.InitialDate,
							FinalDate:   stockFilter.FinalDate,
							Actual:      50,
						}, nil
					},
					AddToStockFunc: func(ctx context.Context, target types.External, mov stock.Movement) (stock.Movement, error) {
						return mov, nil
					},
				},
			},
		},
		{
			name: "should return input movement created with full stock",
			args: args{
				key: product.Key{ID: idExample1},
				movement: stock.Movement{
					TargetID: idExample1,
					Type:     stock.Input,
					Quantity: 10,
				},
			},
			want: product.StockAdditionResult{
				Movement: stock.Movement{
					Entity: types.Entity{
						History: types.History{
							By:        idExample1,
							Operation: operation,
						},
					},
					TargetID: idExample1,
					Type:     stock.Input,
					Quantity: 10,
				},
				RemainingStock: 100,
			},
			wantKey: product.Key{ID: idExample1, Store: types.Store{StoreID: idExample2}},
			opts: opts{
				products: &product.ActionsMock{
					GetProductStockInfoFunc: func(ctx context.Context, key product.Key) (product.Stock, error) {
						return product.Stock{
							CostPrice:       750,
							MaximumQuantity: 100,
							MinimumQuantity: 0,
						}, nil
					},
				},
				stocks: &stock.ActionsMock{
					ComputeStockFunc: func(ctx context.Context, target types.External, stockFilter stock.Filter) (stock.Actual, error) {
						return stock.Actual{
							TargetID:    target,
							InitialDate: stockFilter.InitialDate,
							FinalDate:   stockFilter.FinalDate,
							Actual:      90,
						}, nil
					},
					AddToStockFunc: func(ctx context.Context, target types.External, mov stock.Movement) (stock.Movement, error) {
						return mov, nil
					},
				},
			},
		},
		{
			name: "should return output movement created with full stock",
			args: args{
				key: product.Key{ID: idExample1},
				movement: stock.Movement{
					TargetID: idExample1,
					Type:     stock.Output,
					Quantity: 10,
				},
			},
			want: product.StockAdditionResult{
				Movement: stock.Movement{
					Entity: types.Entity{
						History: types.History{
							By:        idExample1,
							Operation: operation,
						},
					},
					TargetID: idExample1,
					Type:     stock.Output,
					Quantity: 10,
				},
				RemainingStock: 90,
			},
			wantKey: product.Key{ID: idExample1, Store: types.Store{StoreID: idExample2}},
			opts: opts{
				products: &product.ActionsMock{
					GetProductStockInfoFunc: func(ctx context.Context, key product.Key) (product.Stock, error) {
						return product.Stock{
							CostPrice:       750,
							MaximumQuantity: 100,
							MinimumQuantity: 0,
						}, nil
					},
				},
				stocks: &stock.ActionsMock{
					ComputeStockFunc: func(ctx context.Context, target types.External, stockFilter stock.Filter) (stock.Actual, error) {
						return stock.Actual{
							TargetID:    target,
							InitialDate: stockFilter.InitialDate,
							FinalDate:   stockFilter.FinalDate,
							Actual:      100,
						}, nil
					},
					AddToStockFunc: func(ctx context.Context, target types.External, mov stock.Movement) (stock.Movement, error) {
						return mov, nil
					},
				},
			},
		},
		{
			name: "should fail because stock is already full",
			args: args{
				key: product.Key{ID: idExample1},
				movement: stock.Movement{
					TargetID: idExample1,
					Type:     stock.Input,
					Quantity: 10,
				},
			},
			wantErr: product.ErrStockAlreadyFull,
			wantKey: product.Key{ID: idExample1, Store: types.Store{StoreID: idExample2}},
			opts: opts{
				products: &product.ActionsMock{
					GetProductStockInfoFunc: func(ctx context.Context, key product.Key) (product.Stock, error) {
						return product.Stock{
							CostPrice:       750,
							MaximumQuantity: 100,
							MinimumQuantity: 0,
						}, nil
					},
				},
				stocks: &stock.ActionsMock{
					ComputeStockFunc: func(ctx context.Context, target types.External, stockFilter stock.Filter) (stock.Actual, error) {
						return stock.Actual{
							TargetID:    target,
							InitialDate: stockFilter.InitialDate,
							FinalDate:   stockFilter.FinalDate,
							Actual:      100,
						}, nil
					},
					AddToStockFunc: func(ctx context.Context, target types.External, mov stock.Movement) (stock.Movement, error) {
						return mov, nil
					},
				},
			},
		},
		{
			name: "should fail because stock would be lower than allowed",
			args: args{
				key: product.Key{ID: idExample1},
				movement: stock.Movement{
					TargetID: idExample1,
					Type:     stock.Output,
					Quantity: 10,
				},
			},
			wantErr: product.ErrStockNegative,
			wantKey: product.Key{ID: idExample1, Store: types.Store{StoreID: idExample2}},
			opts: opts{
				products: &product.ActionsMock{
					GetProductStockInfoFunc: func(ctx context.Context, key product.Key) (product.Stock, error) {
						return product.Stock{
							CostPrice:       750,
							MaximumQuantity: 100,
							MinimumQuantity: 15,
						}, nil
					},
				},
				stocks: &stock.ActionsMock{
					ComputeStockFunc: func(ctx context.Context, target types.External, stockFilter stock.Filter) (stock.Actual, error) {
						return stock.Actual{
							TargetID:    target,
							InitialDate: stockFilter.InitialDate,
							FinalDate:   stockFilter.FinalDate,
							Actual:      20,
						}, nil
					},
					AddToStockFunc: func(ctx context.Context, target types.External, mov stock.Movement) (stock.Movement, error) {
						return mov, nil
					},
				},
			},
		},
	}

	for _, c := range cases {
		c := c

		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			wf := NewWorkflow(c.opts.products, c.opts.stocks, nil, managers.Transactions)
			stk, err := wf.AddToStock(ctx, c.args.key, c.args.movement)

			assert.ErrorIs(t, err, c.wantErr)
			assert.Equal(t, c.want, stk)

			if err == nil && c.wantErr == nil {
				assert.Equal(t, c.wantKey, c.opts.products.GetProductStockInfoCalls()[0].Key)
			}

		})
	}
}
