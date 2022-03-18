package product

import (
	"context"
	"github.com/stretchr/testify/assert"
	"gomies/pkg/catalog/core/entities/product"
	"testing"
)

func TestWorkflow_RemoveProduct(t *testing.T) {
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
			wantErr error
		}
	)

	cases := []test{
		{
			name: "should return product",
			args: args{
				key: product.Key{ID: idExample1},
			},
			opts: opts{
				products: &product.ActionsMock{
					RemoveFunc: func(ctx context.Context, key product.Key) error {
						return nil
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
			err := wf.RemoveProduct(ctx, c.args.key)

			assert.ErrorIs(t, err, c.wantErr)

		})
	}
}
