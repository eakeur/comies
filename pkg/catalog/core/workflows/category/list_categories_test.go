package category

import (
	"context"
	"github.com/stretchr/testify/assert"
	"gomies/pkg/catalog/core/entities/category"
	"gomies/pkg/sdk/tests"
	"gomies/pkg/sdk/types"
	"testing"
)

func TestWorkflow_ListCategories(t *testing.T) {

	const operation = "Workflows.Category.ListCategories"
	t.Parallel()

	ctx := tests.WorkflowContext(idExample1, idExample2)
	managers := tests.Managers()

	type (
		args struct {
			filter category.Filter
		}

		opts struct {
			categories *category.ActionsMock
		}

		test struct {
			name       string
			args       args
			opts       opts
			want       []category.Category
			wantFilter category.Filter
			wantErr    error
		}
	)

	cases := []test{
		{
			name: "should return product",
			want: []category.Category{
				{}, {},
			},
			wantFilter: category.Filter{Filter: types.Filter{Store: types.Store{StoreID: idExample2}}},
			opts: opts{
				categories: &category.ActionsMock{
					ListFunc: func(ctx context.Context, productFilter category.Filter) ([]category.Category, error) {
						return []category.Category{
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

			wf := NewWorkflow(c.opts.categories, managers.Transactions)
			got, err := wf.ListCategories(ctx, c.args.filter)

			assert.ErrorIs(t, err, c.wantErr)
			assert.Equal(t, c.want, got)

			if err == nil && c.wantErr == nil {
				assert.Equal(t, c.wantFilter, c.opts.categories.ListCalls()[0].CategoryFilter)
			}

		})
	}
}
