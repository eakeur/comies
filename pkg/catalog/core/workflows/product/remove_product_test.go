package product

import (
	"context"
	"github.com/stretchr/testify/assert"
	"gomies/pkg/catalog/core/entities/product"
	"gomies/pkg/sdk/tests"
	"gomies/pkg/sdk/types"
	"testing"
)

func TestWorkflow_RemoveProduct(t *testing.T) {

	const operation = "Workflows.Product.RemoveProduct"
	t.Parallel()

	ctx := tests.WorkflowContext(idExample1, idExample2)
	managers := tests.Managers()

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
			wantKey product.Key
			wantErr error
		}
	)

	cases := []test{
		{
			name: "should return product",
			args: args{
				key: product.Key{ID: idExample1},
			},
			wantKey: product.Key{ID: idExample1, Store: types.Store{StoreID: idExample2}},
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

			wf := NewWorkflow(c.opts.products, nil, nil, managers.Transactions)
			err := wf.RemoveProduct(ctx, c.args.key)

			assert.ErrorIs(t, err, c.wantErr)

			if err == nil && c.wantErr == nil {
				assert.Equal(t, c.wantKey, c.opts.products.RemoveCalls()[0].Key)
			}

		})
	}
}
