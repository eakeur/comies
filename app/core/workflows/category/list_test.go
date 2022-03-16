package category

import (
	"context"
	"gomies/app/core/entities/category"
	"gomies/app/shared/tests"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWorkflow_List(t *testing.T) {
	t.Parallel()

	type args struct {
		ctx    context.Context
		filter category.Filter
	}

	type test struct {
		name    string
		opts    workflow
		args    args
		want    []category.Category
		wantErr error
	}

	cases := []test{
		{
			name: "should return list",
			args: args{
				ctx: context.Background(),
			},
			want: []category.Category{
				{
					Code: "PROD1",
					Name: "Product 1",
				},
			},
			opts: workflow{
				categories: &category.ActionsMock{
					ListFunc: func(contextMoqParam context.Context, filter category.Filter) ([]category.Category, error) {
						return []category.Category{
							{
								Code: "PROD1",
								Name: "Product 1",
							},
						}, nil
					},
				},
			},
		},
		{
			name: "should return empty list",
			args: args{
				ctx: context.Background(),
			},
			want: []category.Category{},
			opts: workflow{
				categories: &category.ActionsMock{
					ListFunc: func(contextMoqParam context.Context, filter category.Filter) ([]category.Category, error) {
						return []category.Category{}, nil
					},
				},
			},
		},
	}

	ctx := tests.GetAuthorizedContext()
	for _, tc := range cases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			tc.args.ctx = ctx
			wf := NewWorkflow(tc.opts.categories, nil)
			prod, err := wf.List(tc.args.ctx, tc.args.filter)

			assert.ErrorIs(t, err, tc.wantErr)
			assert.Equal(t, tc.want, prod)
		})
	}
}
