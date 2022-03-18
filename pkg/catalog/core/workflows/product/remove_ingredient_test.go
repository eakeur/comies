package product

import (
	"context"
	"github.com/stretchr/testify/assert"
	"gomies/pkg/catalog/core/entities/product"
	"gomies/pkg/sdk/tests"
	"gomies/pkg/sdk/types"
	"testing"
)

func TestWorkflow_RemoveIngredient(t *testing.T) {

	const operation = "Workflows.Product.RemoveIngredient"
	t.Parallel()

	ctx := tests.WorkflowContext(idExample1, idExample2)
	managers := tests.Managers()

	type (
		args struct {
			key product.Key
			id  types.External
		}

		opts struct {
			products *product.ActionsMock
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
			name: "should remove ingredient",
			args: args{
				key: product.Key{ID: idExample1},
			},
			wantKey: product.Key{ID: idExample1, Store: types.Store{StoreID: idExample2}},
			opts: opts{
				products: &product.ActionsMock{
					RemoveIngredientFunc: func(ctx context.Context, productKey product.Key, ingredientID types.External) error {
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

			wf := NewWorkflow(c.opts.products, nil, nil, managers.Transactions)
			err := wf.RemoveIngredient(ctx, c.args.key, c.args.id)

			assert.ErrorIs(t, err, c.wantErr)

			if err == nil && c.wantErr == nil {
				assert.Equal(t, c.wantKey, c.opts.products.RemoveIngredientCalls()[0].ProductKey)
			}

		})
	}
}
