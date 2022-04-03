package product

import (
	"context"
	"github.com/stretchr/testify/assert"
	"gomies/app/core/entities/catalog/product"
	"testing"
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
					ListFunc: func(ctx context.Context, productFilter product.Filter) ([]product.Product, error) {
						return []product.Product{
							{}, {},
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

			wf := NewWorkflow(c.opts.products, nil, nil)
			got, err := wf.ListProducts(ctx, c.args.filter)

			assert.ErrorIs(t, err, c.wantErr)
			assert.Equal(t, c.want, got)

		})
	}
}
