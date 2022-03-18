package category

import (
	"context"
	"github.com/stretchr/testify/assert"
	"gomies/pkg/catalog/core/entities/category"
	"gomies/pkg/sdk/tests"
	"gomies/pkg/sdk/types"
	"testing"
)

func TestWorkflow_RemoveCategory(t *testing.T) {

	const operation = "Workflows.Category.RemoveCategory"
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
			wantKey: category.Key{ID: idExample1, Store: types.Store{StoreID: idExample2}},
			opts: opts{
				categories: &category.ActionsMock{
					RemoveFunc: func(ctx context.Context, key category.Key) error {
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

			wf := NewWorkflow(c.opts.categories, managers.Transactions)
			err := wf.RemoveCategory(ctx, c.args.key)

			assert.ErrorIs(t, err, c.wantErr)

			if err == nil && c.wantErr == nil {
				assert.Equal(t, c.wantKey, c.opts.categories.RemoveCalls()[0].CategoryID)
			}

		})
	}
}
