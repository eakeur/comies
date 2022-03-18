package product

import (
	"context"
	"github.com/stretchr/testify/assert"
	"gomies/pkg/catalog/core/entities/product"
	"gomies/pkg/sdk/fault"
	"gomies/pkg/sdk/tests"
	"gomies/pkg/sdk/types"
	"testing"
)

func TestWorkflow_GetProduct(t *testing.T) {

	const operation = "Workflows.Product.GetProduct"
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
			want    product.Product
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
			want: product.Product{
				Code: "PRD",
				Name: "PRD",
				Type: product.OutputType,
			},
			wantKey: product.Key{ID: idExample1, Store: types.Store{StoreID: idExample2}},
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
			wantKey: product.Key{ID: idExample1, Store: types.Store{StoreID: idExample2}},
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

			wf := NewWorkflow(c.opts.products, nil, nil, managers.Transactions)
			got, err := wf.GetProduct(ctx, c.args.key)

			assert.ErrorIs(t, err, c.wantErr)
			assert.Equal(t, c.want, got)

			if err == nil && c.wantErr == nil {
				assert.Equal(t, c.wantKey, c.opts.products.GetCalls()[0].Key)
			}

		})
	}
}
