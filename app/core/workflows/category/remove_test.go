package category

import (
	"context"
	"gomies/app/core/entities/category"
	"gomies/app/core/types/id"
	"gomies/app/shared/tests"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWorkflow_Remove(t *testing.T) {
	t.Parallel()

	transactions := tests.GetFakeManagers().Transaction

	type args struct {
		ctx context.Context
		id  id.External
	}

	type test struct {
		name    string
		opts    workflow
		args    args
		wantErr error
	}

	cases := []test{
		{
			name: "should return nil for deleted category",
			args: args{
				ctx: context.Background(),
				id:  id.ExternalFrom("bf593576-69eb-431b-a956-aaf1337e1f66"),
			},
			opts: workflow{
				categories: &category.ActionsMock{
					RemoveFunc: func(contextMoqParam context.Context, external id.External) error {
						return nil
					},
				},
			},
		},
		{
			name: "should fail because category has products",
			args: args{
				ctx: context.Background(),
				id:  id.ExternalFrom("bf593576-69eb-431b-a956-aaf1337e1f66"),
			},
			wantErr: category.ErrHasDependants,
			opts: workflow{
				categories: &category.ActionsMock{
					RemoveFunc: func(contextMoqParam context.Context, external id.External) error {
						return category.ErrHasDependants
					},
				},
			},
		},
	}

	ctx := tests.GetAuthorizedContext()
	for _, tc := range cases {
		tc := tc

		transactions := transactions
		if tc.opts.transactions != nil {
			transactions = tc.opts.transactions
		}

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			tc.args.ctx = ctx
			wf := NewWorkflow(tc.opts.categories, transactions)
			err := wf.Remove(tc.args.ctx, tc.args.id)

			assert.ErrorIs(t, err, tc.wantErr)
		})
	}
}
