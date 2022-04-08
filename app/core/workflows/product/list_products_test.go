package product

import (
	"context"
	"gomies/app/core/entities/catalog/product"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWorkflow_ListProducts(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	type (
		args struct {
			filter product.Filter
		}

		opts struct {
			products *product.ActionsMock
		}

		test struct {
			name    string
			args    args
			opts    opts
			want    []product.Product
			wantErr error
		}
	)

	cases := []test{
		{
			name: "should return product",
			want: []product.Product{
				{}, {},
			},
			opts: opts{
				products: &product.ActionsMock{
					ListProductsFunc: func(ctx context.Context, productFilter product.Filter) ([]product.Product, int, error) {
						return []product.Product{
							{}, {},
						}, 0, nil
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
			}
			got, _, err := wf.ListProducts(ctx, c.args.filter)

			assert.ErrorIs(t, err, c.wantErr)
			assert.Equal(t, c.want, got)

		})
	}
}
