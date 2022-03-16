package product

import (
	"context"
	"gomies/app/core/entities/product"
	"gomies/app/core/types/entity"
	"gomies/app/core/types/id"
	"gomies/app/shared/tests"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWorkflow_List(t *testing.T) {
	t.Parallel()

	type args struct {
		ctx    context.Context
		filter product.Filter
	}

	type test struct {
		name    string
		opts    workflow
		args    args
		want    []product.Product
		wantErr error
	}

	cases := []test{
		{
			name: "should return list",
			args: args{
				ctx: context.Background(),
			},
			want: []product.Product{
				{
					Entity: entity.Entity{
						ID:         1,
						ExternalID: id.ExternalFrom("bf593576-69eb-431b-a956-aaf1337e1f66"),
					},
					Code: "PROD1",
					Name: "Product 1",
				},
			},
			opts: workflow{
				products: &product.ActionsMock{
					ListFunc: func(contextMoqParam context.Context, filter product.Filter) ([]product.Product, error) {
						return []product.Product{
							{
								Entity: entity.Entity{
									ID:         1,
									ExternalID: id.ExternalFrom("bf593576-69eb-431b-a956-aaf1337e1f66"),
								},
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
			want: []product.Product{},
			opts: workflow{
				products: &product.ActionsMock{
					ListFunc: func(contextMoqParam context.Context, filter product.Filter) ([]product.Product, error) {
						return []product.Product{}, nil
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
			prod, err := wf.List(tc.args.ctx, tc.args.filter)

			assert.ErrorIs(t, err, tc.wantErr)
			assert.Equal(t, tc.want, prod)
		})
	}
}
