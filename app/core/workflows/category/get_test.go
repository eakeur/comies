package category

import (
	"context"
	"gomies/app/core/entities/category"
	"gomies/app/core/types/id"
	"gomies/app/shared/tests"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWorkflow_Get(t *testing.T) {
	t.Parallel()

	type args struct {
		ctx context.Context
		id  id.External
	}

	type test struct {
		name    string
		opts    workflow
		args    args
		want    category.Category
		wantErr error
	}

	cases := []test{
		{
			name: "should return category",
			args: args{
				ctx: context.Background(),
				id:  id.ExternalFrom("bf593576-69eb-431b-a956-aaf1337e1f66"),
			},
			want: category.Category{
				Code: "PROD1",
				Name: "Product 1",
			},
			opts: workflow{
				categories: &category.ActionsMock{
					GetFunc: func(contextMoqParam context.Context, external id.External) (category.Category, error) {
						return category.Category{
							Code: "PROD1",
							Name: "Product 1",
						}, nil
					},
				},
			},
		},
		{
			name: "should fail because category is not found",
			args: args{
				ctx: context.Background(),
				id:  id.ExternalFrom("bf593576-69eb-431b-a956-aaf1337e1f66"),
			},
			wantErr: category.ErrNotFound,
			opts: workflow{
				categories: &category.ActionsMock{
					GetFunc: func(contextMoqParam context.Context, external id.External) (category.Category, error) {
						return category.Category{}, category.ErrNotFound
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
			prod, err := wf.Get(tc.args.ctx, tc.args.id)

			assert.ErrorIs(t, err, tc.wantErr)
			assert.Equal(t, tc.want, prod)
		})
	}
}
