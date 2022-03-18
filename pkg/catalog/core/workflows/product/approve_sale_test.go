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

func TestWorkflow_ApproveSale(t *testing.T) {
	const operation = "Workflows.Product.ApproveSale"
	t.Parallel()

	ctx := tests.WorkflowContext(idExample1, idExample2)
	managers := tests.Managers()

	type (
		args struct {
			req product.ApproveSaleRequest
		}

		opts struct {
			products *product.ActionsMock
			stocks   *stock.ActionsMock
		}

		test struct {
			name    string
			args    args
			opts    opts
			wantKey product.Key
			wantErr error
		}
	)

	cases := []test{
		{
			name: "should return nil for approved sale",
			args: args{
				req: product.ApproveSaleRequest{
					Key:      product.Key{ID: idExample1},
					Quantity: 10,
					Price:    1000,
				},
			},
			wantKey: product.Key{ID: idExample1, Store: types.Store{StoreID: idExample2}},
			opts: opts{
				products: &product.ActionsMock{
					GetProductSaleInfoFunc: func(ctx context.Context, key product.Key) (product.Sale, error) {
						return product.Sale{
							SalePrice:      1000,
							MinimumSale:    1,
							HasIngredients: false,
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
				},
			},
		},
		{
			name: "should return nil for approved sale with ingredients",
			args: args{
				req: product.ApproveSaleRequest{
					Key:      product.Key{ID: idExample1},
					Quantity: 10,
					Price:    1000,
				},
			},
			wantKey: product.Key{ID: idExample1, Store: types.Store{StoreID: idExample2}},
			opts: opts{
				products: &product.ActionsMock{
					GetProductSaleInfoFunc: func(ctx context.Context, key product.Key) (product.Sale, error) {
						return product.Sale{
							SalePrice:      1000,
							MinimumSale:    1,
							HasIngredients: true,
						}, nil
					},
					ListIngredientsFunc: func(ctx context.Context, productKey product.Key) ([]product.Ingredient, error) {
						return []product.Ingredient{
							{Quantity: 1}, {Quantity: 2}, {Quantity: 3}, {Quantity: 4},
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
				},
			},
		},
		{
			name: "should return error for price not allowed",
			args: args{
				req: product.ApproveSaleRequest{
					Key:      product.Key{ID: idExample1},
					Quantity: 10,
					Price:    1000,
				},
			},
			wantKey: product.Key{ID: idExample1, Store: types.Store{StoreID: idExample2}},
			wantErr: product.ErrInvalidSalePrice,
			opts: opts{
				products: &product.ActionsMock{
					GetProductSaleInfoFunc: func(ctx context.Context, key product.Key) (product.Sale, error) {
						return product.Sale{
							SalePrice:   1500,
							MinimumSale: 1,
						}, nil
					},
				},
			},
		},
		{
			name: "should return error for quantity not allowed",
			args: args{
				req: product.ApproveSaleRequest{
					Key:      product.Key{ID: idExample1},
					Quantity: 1,
					Price:    1500,
				},
			},
			wantKey: product.Key{ID: idExample1, Store: types.Store{StoreID: idExample2}},
			wantErr: product.ErrInvalidSaleQuantity,
			opts: opts{
				products: &product.ActionsMock{
					GetProductSaleInfoFunc: func(ctx context.Context, key product.Key) (product.Sale, error) {
						return product.Sale{
							SalePrice:   1500,
							MinimumSale: 5,
						}, nil
					},
				},
			},
		},
		{
			name: "should return error for insufficient ingredients",
			args: args{
				req: product.ApproveSaleRequest{
					Key:      product.Key{ID: idExample1},
					Quantity: 10,
					Price:    1000,
				},
			},
			wantKey: product.Key{ID: idExample1, Store: types.Store{StoreID: idExample2}},
			wantErr: product.ErrNotEnoughStocked,
			opts: opts{
				products: &product.ActionsMock{
					GetProductSaleInfoFunc: func(ctx context.Context, key product.Key) (product.Sale, error) {
						return product.Sale{
							SalePrice:      1000,
							MinimumSale:    1,
							HasIngredients: true,
						}, nil
					},
					ListIngredientsFunc: func(ctx context.Context, productKey product.Key) ([]product.Ingredient, error) {
						return []product.Ingredient{
							{Quantity: 2}, {Quantity: 4}, {Quantity: 6}, {Quantity: 8},
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
				},
			},
		},
	}

	for _, c := range cases {
		c := c

		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			wf := NewWorkflow(c.opts.products, c.opts.stocks, nil, managers.Transactions)
			err := wf.ApproveSale(ctx, c.args.req)

			assert.ErrorIs(t, err, c.wantErr)

			if err == nil && c.wantErr == nil {
				assert.Equal(t, c.wantKey, c.opts.products.GetProductSaleInfoCalls()[0].Key)
			}

		})
	}
}
