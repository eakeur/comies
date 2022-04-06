package product

import (
	"context"
	"gomies/app/core/entities/catalog/category"
	"gomies/app/core/entities/catalog/product"
	"gomies/pkg/sdk/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWorkflow_UpdateProduct(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	fakeID := types.UID("1bdcafba-9deb-48b4-8a0e-ecea4c99b0e3")

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
			wantErr error
		}
	)

	cases := []test{
		{
			name: "should return nil for successful update",
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
			opts: opts{
				products: &product.ActionsMock{
					UpdateProductFunc: func(ctx context.Context, prd product.Product) error {
						return nil
					},
				},
				categories: &category.ActionsMock{
					GetCategoryFunc: func(ctx context.Context, categoryKey category.Key) (category.Category, error) {
						return category.Category{}, nil
					},
				},
			},
		},
		{
			name: "should return nil for updated product with category",
			args: args{
				product: product.Product{
					Code:               "PRD",
					Name:               "Product",
					CategoryExternalID: fakeID,
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
			opts: opts{
				products: &product.ActionsMock{
					UpdateProductFunc: func(ctx context.Context, prd product.Product) error {
						return nil
					},
				},
				categories: &category.ActionsMock{
					GetCategoryFunc: func(ctx context.Context, categoryKey category.Key) (category.Category, error) {
						return category.Category{Entity: types.Entity{ID: 2}}, nil
					},
				},
			},
		},
		{
			name: "should return nil for successful update without category",
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
			opts: opts{
				products: &product.ActionsMock{
					UpdateProductFunc: func(ctx context.Context, prd product.Product) error {
						return nil
					},
				},
				categories: &category.ActionsMock{
					GetCategoryFunc: func(ctx context.Context, categoryKey category.Key) (category.Category, error) {
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
					UpdateProductFunc: func(ctx context.Context, prd product.Product) error {
						return nil
					},
				},
				categories: &category.ActionsMock{
					GetCategoryFunc: func(ctx context.Context, categoryKey category.Key) (category.Category, error) {
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
					UpdateProductFunc: func(ctx context.Context, prd product.Product) error {
						return nil
					},
				},
				categories: &category.ActionsMock{
					GetCategoryFunc: func(ctx context.Context, categoryKey category.Key) (category.Category, error) {
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
					UpdateProductFunc: func(ctx context.Context, prd product.Product) error {
						return nil
					},
				},
				categories: &category.ActionsMock{
					GetCategoryFunc: func(ctx context.Context, categoryKey category.Key) (category.Category, error) {
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
					UpdateProductFunc: func(ctx context.Context, prd product.Product) error {
						return nil
					},
				},
				categories: &category.ActionsMock{
					GetCategoryFunc: func(ctx context.Context, categoryKey category.Key) (category.Category, error) {
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

			wf := workflow{
				products:   c.opts.products,
				categories: c.opts.categories,
			}
			err := wf.UpdateProduct(ctx, c.args.product)
			assert.ErrorIs(t, err, c.wantErr)
		})
	}

}
