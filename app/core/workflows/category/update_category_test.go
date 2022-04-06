package category

import (
	"context"
	"gomies/app/core/entities/catalog/category"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWorkflow_UpdateCategory(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	type (
		args struct {
			category category.Category
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
			name: "should return nil for successful update",
			args: args{
				category: category.Category{
					Code: "AAA",
					Name: "AAAAAAAA",
				},
			},
			opts: opts{
				categories: &category.ActionsMock{
					UpdateCategoryFunc: func(ctx context.Context, category category.Category) error {
						return nil
					},
				},
			},
		},
		{
			name: "should return error for invalid code",
			args: args{
				category: category.Category{
					Code: "",
					Name: "AAAAAAAA",
				},
			},
			wantErr: category.ErrInvalidCode,
			opts: opts{
				categories: &category.ActionsMock{
					UpdateCategoryFunc: func(ctx context.Context, category category.Category) error {
						return nil
					},
				},
			},
		},
		{
			name: "should return error for invalid name",
			args: args{
				category: category.Category{
					Code: "AAA",
					Name: "",
				},
			},
			wantErr: category.ErrInvalidName,
			opts: opts{
				categories: &category.ActionsMock{
					UpdateCategoryFunc: func(ctx context.Context, category category.Category) error {
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

			wf := workflow{
				categories: c.opts.categories,
			}
			err := wf.UpdateCategory(ctx, c.args.category)
			assert.ErrorIs(t, err, c.wantErr)
		})
	}

}
