package product

import (
	"context"
	"github.com/stretchr/testify/assert"
	"gomies/app/core/entities/catalog/product"
	"gomies/pkg/sdk/fault"
	"testing"
)

func TestWorkflow_GetProduct(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	type (
		args struct {
			key product.Key
		}

		opts struct {
			products *product.ActionsMock
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
			name: "should return product",
			args: args{
				key: product.Key{ID: idExample1},
			},
			want: product.Product{
				Code: "PRD",
				Name: "PRD",
				Type: product.OutputType,
			},
			opts: opts{
				products: &product.ActionsMock{
					GetFunc: func(ctx context.Context, key product.Key) (product.Product, error) {
						return product.Product{
							Code: "PRD",
							Name: "PRD",
							Type: product.OutputType,
						}, nil
					},
				},
			},
		},
		{
			name: "should return product not found",
			args: args{
				key: product.Key{ID: idExample1},
			},
			wantErr: fault.ErrNotFound,
			opts: opts{
				products: &product.ActionsMock{
					GetFunc: func(ctx context.Context, key product.Key) (product.Product, error) {
						return product.Product{}, fault.ErrNotFound
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
			got, err := wf.GetProduct(ctx, c.args.key)

			assert.ErrorIs(t, err, c.wantErr)
			assert.Equal(t, c.want, got)

		})
	}
}
