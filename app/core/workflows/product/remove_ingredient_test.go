package product

import (
	"context"
	"gomies/app/core/entities/catalog/product"
	"gomies/pkg/sdk/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWorkflow_RemoveIngredient(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	fakeID := types.UID("1bdcafba-9deb-48b4-8a0e-ecea4c99b0e3")

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
				key: product.Key{ID: fakeID},
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

			wf := workflow{
				products: c.opts.products,
			}
			err := wf.RemoveIngredient(ctx, c.args.key, c.args.id)

			assert.ErrorIs(t, err, c.wantErr)

		})
	}
}
