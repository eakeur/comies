package category

import (
	"context"
	"github.com/stretchr/testify/assert"
	"gomies/app/core/entities/catalog/category"
	"testing"
)

func TestWorkflow_RemoveCategory(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

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
				key: category.Key{ID: idExample1},
			},
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

			wf := NewWorkflow(c.opts.categories)
			err := wf.RemoveCategory(ctx, c.args.key)

			assert.ErrorIs(t, err, c.wantErr)

		})
	}
}
