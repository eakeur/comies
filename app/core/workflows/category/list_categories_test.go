package category

import (
	"context"
	"gomies/app/core/entities/catalog/category"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWorkflow_ListCategories(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	type (
		args struct {
			filter category.Filter
		}

		opts struct {
			categories *category.ActionsMock
		}

		test struct {
			name    string
			args    args
			opts    opts
			want    []category.Category
			wantErr error
		}
	)

	cases := []test{
		{
			name: "should return product",
			want: []category.Category{
				{}, {},
			},
			opts: opts{
				categories: &category.ActionsMock{
					ListCategoriesFunc: func(ctx context.Context, productFilter category.Filter) ([]category.Category, int, error) {
						return []category.Category{
							{}, {},
						}, 0, nil
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
			got, _, err := wf.ListCategories(ctx, c.args.filter)

			assert.ErrorIs(t, err, c.wantErr)
			assert.Equal(t, c.want, got)
		})
	}
}
