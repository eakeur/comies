package product

import (
	"context"
	"github.com/stretchr/testify/assert"
	"gomies/app/core/entities/catalog/product"
	"gomies/pkg/sdk/types"
	"testing"
)

func TestWorkflow_RemoveIngredient(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	type (
		args struct {
			key product.Key
			id  types.UID
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
			name: "should remove ingredient",
			args: args{
				key: product.Key{ID: idExample1},
			},
			opts: opts{
				products: &product.ActionsMock{
					RemoveIngredientFunc: func(ctx context.Context, productKey product.Key, ingredientID types.UID) error {
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
			err := wf.RemoveIngredient(ctx, c.args.key, c.args.id)

			assert.ErrorIs(t, err, c.wantErr)

		})
	}
}
