package category

import (
	"context"
	"gomies/app/core/entities/catalog/category"
	"gomies/pkg/sdk/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWorkflow_RemoveCategory(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	fakeID := types.UID("1bdcafba-9deb-48b4-8a0e-ecea4c99b0e3")

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
			wantErr error
		}
	)

	cases := []test{
		{
			name: "should return category",
			args: args{
				key: category.Key{ID: fakeID},
			},
			opts: opts{
				categories: &category.ActionsMock{
					RemoveCategoryFunc: func(ctx context.Context, key category.Key) error {
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

			wf := NewWorkflow(c.opts.categories)
			err := wf.RemoveCategory(ctx, c.args.key)

			assert.ErrorIs(t, err, c.wantErr)

		})
	}
}
