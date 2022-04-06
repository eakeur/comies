package category

import (
	"context"
	"gomies/app/core/entities/catalog/category"
	"gomies/pkg/sdk/fault"
	"gomies/pkg/sdk/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWorkflow_GetCategory(t *testing.T) {
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
			want    category.Category
			wantErr error
		}
	)

	cases := []test{
		{
			name: "should return category",
			args: args{
				key: category.Key{ID: fakeID},
			},
			want: category.Category{
				Code: "PRD",
				Name: "PRD",
			},
			opts: opts{
				categories: &category.ActionsMock{
					GetCategoryFunc: func(ctx context.Context, key category.Key) (category.Category, error) {
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
				key: category.Key{ID: fakeID},
			},
			wantErr: fault.ErrNotFound,
			opts: opts{
				categories: &category.ActionsMock{
					GetCategoryFunc: func(ctx context.Context, key category.Key) (category.Category, error) {
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
