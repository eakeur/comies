package category

import (
	"context"
	"github.com/stretchr/testify/assert"
	"gomies/pkg/catalog/core/entities/category"
	"gomies/pkg/sdk/fault"
	"gomies/pkg/sdk/tests"
	"gomies/pkg/sdk/types"
	"testing"
)

func TestWorkflow_GetCategory(t *testing.T) {

	const operation = "Workflows.Category.GetCategory"
	t.Parallel()

	ctx := tests.WorkflowContext(idExample1, idExample2)
	managers := tests.Managers()

	type (
		args struct {
			key category.Key
		}

		opts struct {
			categories *category.ActionsMock
		}

		test struct {
			name    string
			args    args
			opts    opts
			want    category.Category
			wantKey category.Key
			wantErr error
		}
	)

	cases := []test{
		{
			name: "should return category",
			args: args{
				key: category.Key{ID: idExample1},
			},
			want: category.Category{
				Code: "PRD",
				Name: "PRD",
			},
			wantKey: category.Key{ID: idExample1, Store: types.Store{StoreID: idExample2}},
			opts: opts{
				categories: &category.ActionsMock{
					GetFunc: func(ctx context.Context, key category.Key) (category.Category, error) {
						return category.Category{
							Code: "PRD",
							Name: "PRD",
						}, nil
					},
				},
			},
		},
		{
			name: "should return category not found",
			args: args{
				key: category.Key{ID: idExample1},
			},
			wantErr: fault.ErrNotFound,
			wantKey: category.Key{ID: idExample1, Store: types.Store{StoreID: idExample2}},
			opts: opts{
				categories: &category.ActionsMock{
					GetFunc: func(ctx context.Context, key category.Key) (category.Category, error) {
						return category.Category{}, fault.ErrNotFound
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
			got, err := wf.GetCategory(ctx, c.args.key)

			assert.ErrorIs(t, err, c.wantErr)
			assert.Equal(t, c.want, got)

			if err == nil && c.wantErr == nil {
				assert.Equal(t, c.wantKey, c.opts.categories.GetCalls()[0].CategoryKey)
			}

		})
	}
}
