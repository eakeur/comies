package category

import (
	"context"
	"gomies/app/core/entities/category"
	"gomies/app/shared/tests"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWorkflow_Update(t *testing.T) {
	t.Parallel()

	transactions := tests.GetFakeManagers().Transaction

	type args struct {
		ctx context.Context
		prd category.Category
	}

	type test struct {
		name    string
		opts    workflow
		args    args
		wantErr error
	}

	cases := []test{
		{
			name: "should update category",
			args: args{
				ctx: context.Background(),
				prd: category.Category{
					Code: "PRD1",
					Name: "Product 1",
				},
			},
			opts: workflow{
				categories: &category.ActionsMock{
					UpdateFunc: func(contextMoqParam context.Context, category category.Category) error {
						return nil
					},
				},
			},
		},
		{
			name: "should fail because category has no code",
			args: args{
				ctx: context.Background(),
				prd: category.Category{
					Code: "",
					Name: "Product 1",
				},
			},
			wantErr: category.ErrInvalidCode,
		},
		{
			name: "should fail because category with code already exists",
			args: args{
				ctx: context.Background(),
				prd: category.Category{
					Code: "PRD1",
					Name: "Product 1",
				},
			},
			wantErr: category.ErrAlreadyExists,
			opts: workflow{
				categories: &category.ActionsMock{
					UpdateFunc: func(contextMoqParam context.Context, c category.Category) error {
						return category.ErrAlreadyExists
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
			transactions := transactions
			if tc.opts.transactions != nil {
				transactions = tc.opts.transactions
			}

			wf := NewWorkflow(tc.opts.categories, transactions)
			err := wf.Update(tc.args.ctx, tc.args.prd)

			assert.ErrorIs(t, err, tc.wantErr)
		})
	}
}
