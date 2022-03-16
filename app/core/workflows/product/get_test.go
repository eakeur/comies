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
		want    product.Product
		wantErr error
	}

	cases := []test{
		{
			name: "should return product",
			args: args{
				ctx: context.Background(),
				id:  id.ExternalFrom("bf593576-69eb-431b-a956-aaf1337e1f66"),
			},
			want: product.Product{
				Entity: entity.Entity{
					ID:         1,
					ExternalID: id.ExternalFrom("bf593576-69eb-431b-a956-aaf1337e1f66"),
				},
				Code: "PROD1",
				Name: "Product 1",
			},
			opts: workflow{
				products: &product.ActionsMock{
					GetFunc: func(contextMoqParam context.Context, external id.External, additionalDataToConsiders ...product.AdditionalDataToConsider) (product.Product, error) {
						return product.Product{
							Entity: entity.Entity{
								ID:         1,
								ExternalID: id.ExternalFrom("bf593576-69eb-431b-a956-aaf1337e1f66"),
							},
							Code: "PROD1",
							Name: "Product 1",
						}, nil
					},
				},
			},
		},
		{
			name: "should fail because product is not found",
			args: args{
				ctx: context.Background(),
				id:  id.ExternalFrom("bf593576-69eb-431b-a956-aaf1337e1f66"),
			},
			wantErr: product.ErrNotFound,
			opts: workflow{
				products: &product.ActionsMock{
					GetFunc: func(contextMoqParam context.Context, external id.External, additionalDataToConsiders ...product.AdditionalDataToConsider) (product.Product, error) {
						return product.Product{}, product.ErrNotFound
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
			prod, err := wf.Get(tc.args.ctx, tc.args.id)

			assert.ErrorIs(t, err, tc.wantErr)
			assert.Equal(t, tc.want, prod)
		})
	}
}
