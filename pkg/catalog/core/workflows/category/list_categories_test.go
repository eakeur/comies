package category

import (
	"context"
	"github.com/stretchr/testify/assert"
	"gomies/pkg/catalog/core/entities/category"
	"testing"
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

			wf := NewWorkflow(c.opts.categories)
			got, err := wf.ListCategories(ctx, c.args.filter)

			assert.ErrorIs(t, err, c.wantErr)
			assert.Equal(t, c.want, got)
		})
	}
}
