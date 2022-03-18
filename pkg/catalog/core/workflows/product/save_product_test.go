package product

import (
	"context"
	"github.com/stretchr/testify/assert"
	"gomies/pkg/catalog/core/entities/category"
	"gomies/pkg/catalog/core/entities/product"
	"gomies/pkg/sdk/types"
	"testing"
)

func TestWorkflow_SaveProduct(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	type (
		args struct {
			product product.Product
		}

		opts struct {
			categories *category.ActionsMock
			products   *product.ActionsMock
		}

		test struct {
			name    string
			args    args
			opts    opts
			want    product.Product
			wantErr error
		}
	)

	cases := []test{
		{
			name: "should return product created",
			args: args{
				product: product.Product{
					Code: "PRD",
					Name: "Product",
					Type: product.OutputType,
					Stock: product.Stock{
						CostPrice: 1,
					},
					Sale: product.Sale{
						SalePrice:   2,
						MinimumSale: 2,
					},
				},
			},
			want: product.Product{
				Code: "PRD",
				Name: "Product",
				Type: product.OutputType,
				Stock: product.Stock{
					CostPrice: 1,
				},
				Sale: product.Sale{
					SalePrice:   2,
					MinimumSale: 2,
				},
			},
			opts: opts{
				products: &product.ActionsMock{
					SaveFunc: func(ctx context.Context, prd product.Product, flag ...types.WritingFlag) (product.Product, error) {
						return prd, nil
					},
				},
				categories: &category.ActionsMock{
					GetFunc: func(ctx context.Context, categoryKey category.Key) (category.Category, error) {
						return category.Category{}, nil
					},
				},
			},
		},
		{
			name: "should return product with category created",
			args: args{
				product: product.Product{
					Code:               "PRD",
					Name:               "Product",
					CategoryExternalID: idExample2,
					Type:               product.OutputType,
					Stock: product.Stock{
						CostPrice: 1,
					},
					Sale: product.Sale{
						SalePrice:   2,
						MinimumSale: 2,
					},
				},
			},
			want: product.Product{
				Code:               "PRD",
				Name:               "Product",
				CategoryID:         2,
				CategoryExternalID: idExample2,
				Type:               product.OutputType,
				Stock: product.Stock{
					CostPrice: 1,
				},
				Sale: product.Sale{
					SalePrice:   2,
					MinimumSale: 2,
				},
			},
			opts: opts{
				products: &product.ActionsMock{
					SaveFunc: func(ctx context.Context, prd product.Product, flag ...types.WritingFlag) (product.Product, error) {
						return prd, nil
					},
				},
				categories: &category.ActionsMock{
					GetFunc: func(ctx context.Context, categoryKey category.Key) (category.Category, error) {
						return category.Category{Entity: types.Entity{ID: 2}}, nil
					},
				},
			},
		},
		{
			name: "should return input product created ",
			args: args{
				product: product.Product{
					Code: "PRD",
					Name: "Product",
					Type: product.InputType,
					Stock: product.Stock{
						CostPrice: 1,
					},
				},
			},
			want: product.Product{
				Code: "PRD",
				Name: "Product",
				Type: product.InputType,
				Stock: product.Stock{
					CostPrice: 1,
				},
			},
			opts: opts{
				products: &product.ActionsMock{
					SaveFunc: func(ctx context.Context, prd product.Product, flag ...types.WritingFlag) (product.Product, error) {
						return prd, nil
					},
				},
				categories: &category.ActionsMock{
					GetFunc: func(ctx context.Context, categoryKey category.Key) (category.Category, error) {
						return category.Category{}, nil
					},
				},
			},
		},
		{
			name: "should return error for invalid code",
			args: args{
				product: product.Product{
					Code: "P",
					Name: "Product",
					Type: product.OutputType,
					Stock: product.Stock{
						CostPrice: 1,
					},
					Sale: product.Sale{
						SalePrice:   2,
						MinimumSale: 2,
					},
				},
			},
			wantErr: product.ErrInvalidCode,
			opts: opts{
				products: &product.ActionsMock{
					SaveFunc: func(ctx context.Context, prd product.Product, flag ...types.WritingFlag) (product.Product, error) {
						return prd, nil
					},
				},
				categories: &category.ActionsMock{
					GetFunc: func(ctx context.Context, categoryKey category.Key) (category.Category, error) {
						return category.Category{}, nil
					},
				},
			},
		},
		{
			name: "should return error for invalid name",
			args: args{
				product: product.Product{
					Code: "PRD",
					Name: "-",
					Type: product.OutputType,
					Stock: product.Stock{
						CostPrice: 1,
					},
					Sale: product.Sale{
						SalePrice:   2,
						MinimumSale: 2,
					},
				},
			},
			wantErr: product.ErrInvalidName,
			opts: opts{
				products: &product.ActionsMock{
					SaveFunc: func(ctx context.Context, prd product.Product, flag ...types.WritingFlag) (product.Product, error) {
						return prd, nil
					},
				},
				categories: &category.ActionsMock{
					GetFunc: func(ctx context.Context, categoryKey category.Key) (category.Category, error) {
						return category.Category{}, nil
					},
				},
			},
		},
		{
			name: "should return error for invalid price",
			args: args{
				product: product.Product{
					Code: "PRD",
					Name: "Product",
					Type: product.OutputType,
					Stock: product.Stock{
						CostPrice: 1,
					},
					Sale: product.Sale{
						MinimumSale: 2,
					},
				},
			},
			wantErr: product.ErrInvalidPrice,
			opts: opts{
				products: &product.ActionsMock{
					SaveFunc: func(ctx context.Context, prd product.Product, flag ...types.WritingFlag) (product.Product, error) {
						return prd, nil
					},
				},
				categories: &category.ActionsMock{
					GetFunc: func(ctx context.Context, categoryKey category.Key) (category.Category, error) {
						return category.Category{}, nil
					},
				},
			},
		},
		{
			name: "should return error for invalid quantity",
			args: args{
				product: product.Product{
					Code: "PRD",
					Name: "Product",
					Type: product.OutputType,
					Stock: product.Stock{
						CostPrice: 1,
					},
					Sale: product.Sale{
						SalePrice: 3,
					},
				},
			},
			wantErr: product.ErrMinimumSaleQuantity,
			opts: opts{
				products: &product.ActionsMock{
					SaveFunc: func(ctx context.Context, prd product.Product, flag ...types.WritingFlag) (product.Product, error) {
						return prd, nil
					},
				},
				categories: &category.ActionsMock{
					GetFunc: func(ctx context.Context, categoryKey category.Key) (category.Category, error) {
						return category.Category{}, nil
					},
				},
			},
		},
	}

	for _, c := range cases {
		c := c

		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			wf := NewWorkflow(c.opts.products, c.opts.categories, nil)
			ingredient, err := wf.SaveProduct(ctx, c.args.product)

			assert.ErrorIs(t, err, c.wantErr)
			assert.Equal(t, c.want, ingredient)

		})
	}

}
