package category

import (
	"context"
	"github.com/stretchr/testify/assert"
	"gomies/pkg/catalog/core/entities/category"
	"gomies/pkg/sdk/fault"
	"testing"
)

func TestWorkflow_GetCategory(t *testing.T) {
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
			want    category.Category
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

			wf := NewWorkflow(c.opts.categories)
			got, err := wf.GetCategory(ctx, c.args.key)

			assert.ErrorIs(t, err, c.wantErr)
			assert.Equal(t, c.want, got)

		})
	}
}
