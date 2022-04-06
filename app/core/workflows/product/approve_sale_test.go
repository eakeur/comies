package product

import (
	"context"
	"gomies/app/core/entities/catalog/product"
	"gomies/pkg/sdk/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWorkflow_ApproveSale(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	fakeID := types.UID("1bdcafba-9deb-48b4-8a0e-ecea4c99b0e3")

	type (
		args struct {
			req product.ApproveSaleRequest
		}

		opts struct {
			products *product.ActionsMock
			stocks   *StockServiceMock
		}

		test struct {
			name    string
			args    args
			opts    opts
			wantErr error
		}
	)

	cases := []test{
		{
			name: "should return nil for approved sale",
			args: args{
				req: product.ApproveSaleRequest{
					Key:      product.Key{ID: fakeID},
					Quantity: 10,
					Price:    1000,
				},
			},
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
				stocks: &StockServiceMock{
					ComputeFunc: func(ctx context.Context, productID types.UID) (types.Quantity, error) {
						return 50, nil
					},
				},
			},
		},
		{
			name: "should return nil for approved sale with ingredients",
			args: args{
				req: product.ApproveSaleRequest{
					Key:      product.Key{ID: fakeID},
					Quantity: 10,
					Price:    1000,
				},
			},
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
				stocks: &StockServiceMock{
					ComputeSomeFunc: func(ctx context.Context, productsIDs ...types.UID) ([]types.Quantity, error) {
						return []types.Quantity{50, 50, 50, 50}, nil
					},
				},
			},
		},
		{
			name: "should return error for price not allowed",
			args: args{
				req: product.ApproveSaleRequest{
					Key:      product.Key{ID: fakeID},
					Quantity: 10,
					Price:    1000,
				},
			},
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
					Key:      product.Key{ID: fakeID},
					Quantity: 1,
					Price:    1500,
				},
			},
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
					Key:      product.Key{ID: fakeID},
					Quantity: 10,
					Price:    1000,
				},
			},
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
				stocks: &StockServiceMock{
					ComputeSomeFunc: func(ctx context.Context, productsIDs ...types.UID) ([]types.Quantity, error) {
						return []types.Quantity{50, 50, 50, 50}, nil
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
			err := wf.ApproveSale(ctx, c.args.req)

			assert.ErrorIs(t, err, c.wantErr)

		})
	}
}
