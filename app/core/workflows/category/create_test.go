package category

import (
	"context"
	"gomies/app/core/entities/category"
	"gomies/app/core/types/entity"
	"gomies/app/core/types/history"
	"gomies/app/core/types/id"
	"gomies/app/shared/tests"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWorkflow_Create(t *testing.T) {
	const operation = "Workflows.Category.Create"
	t.Parallel()

	transactions := tests.GetFakeManagers().Transaction

	type args struct {
		ctx context.Context
		cat category.Category
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
			name: "should return category created",
			args: args{
				ctx: context.Background(),
				cat: category.Category{
					Name: "CAT",
					Code: "Category 1",
				},
			},
			want: category.Category{
				Entity: entity.Entity{
					StoreExternalID: id.ExternalFrom("7a4ad106-f91d-4898-955d-91f0e7e93972"),
					StoreID:         1,
					Active:          true,
					History: history.History{
						Operation: operation,
						By:        id.ExternalFrom("7a4ad106-f91d-4898-955d-91f0e7e93971"),
					},
				},
				Name: "CAT",
				Code: "Category 1",
			},
			opts: workflow{
				categories: &category.ActionsMock{
					CreateFunc: func(ctx context.Context, c category.Category) (category.Category, error) {
						return c, nil
					},
				},
			},
		},
		{
			name: "should fail because product has no code",
			args: args{
				ctx: context.Background(),
				cat: category.Category{Name: "Category"},
			},
			wantErr: category.ErrInvalidCode,
		},
		{
			name: "should fail because product has no name",
			args: args{
				ctx: context.Background(),
				cat: category.Category{Code: "CAT1"},
			},
			wantErr: category.ErrInvalidName,
		},
		{
			name: "should fail because category with code already exists",
			args: args{
				ctx: context.Background(),
				cat: category.Category{
					Name: "CAT",
					Code: "Category 1",
				},
			},
			wantErr: category.ErrAlreadyExists,
			opts: workflow{
				categories: &category.ActionsMock{
					CreateFunc: func(contextMoqParam context.Context, c category.Category) (category.Category, error) {
						return c, category.ErrAlreadyExists
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
			cat, err := wf.Create(tc.args.ctx, tc.args.cat)

			assert.ErrorIs(t, err, tc.wantErr)
			assert.Equal(t, tc.want, cat)
		})
	}
}
