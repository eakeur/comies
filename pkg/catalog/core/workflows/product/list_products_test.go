package product

import (
	"context"
	"github.com/stretchr/testify/assert"
	"gomies/pkg/catalog/core/entities/product"
	"gomies/pkg/sdk/tests"
	"gomies/pkg/sdk/types"
	"testing"
)

func TestWorkflow_ListProducts(t *testing.T) {

	const operation = "Workflows.Product.ListProducts"
	t.Parallel()

	ctx := tests.WorkflowContext(idExample1, idExample2)
	managers := tests.Managers()

	type (
		args struct {
			filter product.Filter
		}

		opts struct {
			products *product.ActionsMock
		}

		test struct {
			name       string
			args       args
			opts       opts
			want       []product.Product
			wantFilter product.Filter
			wantErr    error
		}
	)

	cases := []test{
		{
			name: "should return product",
			want: []product.Product{
				{}, {},
			},
			wantFilter: product.Filter{Filter: types.Filter{Store: types.Store{StoreID: idExample2}}},
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

			wf := NewWorkflow(c.opts.products, nil, nil, managers.Transactions)
			got, err := wf.ListProducts(ctx, c.args.filter)

			assert.ErrorIs(t, err, c.wantErr)
			assert.Equal(t, c.want, got)

			if err == nil && c.wantErr == nil {
				assert.Equal(t, c.wantFilter, c.opts.products.ListCalls()[0].ProductFilter)
			}

		})
	}
}
