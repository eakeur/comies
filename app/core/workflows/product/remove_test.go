package product

import (
	"context"
	"gomies/app/core/entities/product"
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
			name: "should return nil for deleted product",
			args: args{
				ctx: context.Background(),
				id:  id.ExternalFrom("bf593576-69eb-431b-a956-aaf1337e1f66"),
			},
			opts: workflow{
				products: &product.ActionsMock{
					RemoveFunc: func(contextMoqParam context.Context, external id.External) error {
						return nil
					},
				},
			},
		},
	}

	ctx := tests.GetAuthorizedContext()

	for _, tc := range cases {
		tc := tc
		tc.args.ctx = ctx
		transactions := transactions
		if tc.opts.transactions != nil {
			transactions = tc.opts.transactions
		}

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			wf := NewWorkflow(tc.opts.products, tc.opts.stocks, tc.opts.categories, transactions)
			err := wf.Remove(tc.args.ctx, tc.args.id)

			assert.ErrorIs(t, err, tc.wantErr)
		})
	}
}
