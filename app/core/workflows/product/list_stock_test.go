package product

import (
	"context"
	"gomies/app/core/entities/stock"
	"gomies/app/core/types/id"
	"gomies/app/shared/tests"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestWorkflow_ListStock(t *testing.T) {
	t.Parallel()

	type args struct {
		ctx    context.Context
		filter stock.Filter
	}

	type test struct {
		name    string
		opts    workflow
		args    args
		want    []stock.Movement
		wantErr error
	}

	cases := []test{
		{
			name: "should return list",
			args: args{
				ctx: context.Background(),
			},
			want: []stock.Movement{
				{
					TargetID:       id.ExternalFrom("bf593576-69eb-431b-a956-aaf1337e1f66"),
					Type:           stock.Output,
					Date:           time.Time{},
					Quantity:       4,
					AdditionalData: "observations",
				},
				{
					TargetID:       id.ExternalFrom("bf593576-69eb-431b-a956-aaf1337e1f63"),
					Type:           stock.Output,
					Date:           time.Time{},
					Quantity:       2,
					AdditionalData: "observations",
				},
			},
			opts: workflow{
				stocks: &stock.ActionsMock{
					ListMovementsFunc: func(contextMoqParam context.Context, filter stock.Filter) ([]stock.Movement, error) {
						return []stock.Movement{
							{
								TargetID:       id.ExternalFrom("bf593576-69eb-431b-a956-aaf1337e1f66"),
								Type:           stock.Output,
								Date:           time.Time{},
								Quantity:       4,
								AdditionalData: "observations",
							},
							{
								TargetID:       id.ExternalFrom("bf593576-69eb-431b-a956-aaf1337e1f63"),
								Type:           stock.Output,
								Date:           time.Time{},
								Quantity:       2,
								AdditionalData: "observations",
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
			want: []stock.Movement{},
			opts: workflow{
				stocks: &stock.ActionsMock{
					ListMovementsFunc: func(contextMoqParam context.Context, filter stock.Filter) ([]stock.Movement, error) {
						return []stock.Movement{}, nil
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
			wf := NewWorkflow(tc.opts.products, tc.opts.stocks, tc.opts.categories, nil)
			prod, err := wf.ListStock(tc.args.ctx, tc.args.filter)

			assert.ErrorIs(t, err, tc.wantErr)
			assert.Equal(t, tc.want, prod)
		})
	}
}
