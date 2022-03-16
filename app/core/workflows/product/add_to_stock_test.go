package product

import (
	"context"
	"gomies/app/core/entities/product"
	"gomies/app/core/entities/stock"
	"gomies/app/core/types/entity"
	"gomies/app/core/types/history"
	"gomies/app/core/types/id"
	"gomies/app/shared/tests"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestWorkflow_AddToStock(t *testing.T) {
	const operation = "Workflows.Product.AddToStock"

	t.Parallel()

	transactions := tests.GetFakeManagers().Transaction

	type args struct {
		ctx context.Context
		mov stock.Movement
	}

	type test struct {
		name    string
		opts    workflow
		args    args
		want    stock.Movement
		wantErr error
	}

	cases := []test{
		{
			name: "should return stock movement created",
			args: args{
				ctx: context.Background(),
				mov: stock.Movement{
					TargetID:       id.ExternalFrom("bf593576-69eb-431b-a956-aaf1337e1f66"),
					Type:           stock.Output,
					Date:           time.Time{},
					Quantity:       4,
					AdditionalData: "observations",
				},
			},
			want: stock.Movement{
				Entity: entity.Entity{
					StoreExternalID: id.ExternalFrom("7a4ad106-f91d-4898-955d-91f0e7e93972"),
					StoreID:         1,
					Active:          true,
					History: history.History{
						Operation: operation,
						By:        id.ExternalFrom("7a4ad106-f91d-4898-955d-91f0e7e93971"),
					},
				},
				TargetID:       id.ExternalFrom("bf593576-69eb-431b-a956-aaf1337e1f66"),
				Type:           stock.Output,
				Date:           time.Time{},
				Quantity:       4,
				AdditionalData: "observations",
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
							Stock: product.StockInformation{
								MaximumQuantity: 100,
								MinimumQuantity: 10,
							},
						}, nil
					},
				},
				stocks: &stock.ActionsMock{
					ComputeStockFunc: func(contextMoqParam context.Context, filter stock.Filter) (stock.Actual, error) {
						return stock.Actual{
							TargetID: id.ExternalFrom("bf593576-69eb-431b-a956-aaf1337e1f66"),
							Actual:   56,
						}, nil
					},
					AddToStockFunc: func(ctx context.Context, movement stock.Movement) (stock.Movement, error) {
						return movement, nil
					},
				},
			},
		},
		{
			name: "should fail because target id is nil",
			args: args{
				ctx: context.Background(),
				mov: stock.Movement{
					TargetID: id.Nil,
				},
			},
			wantErr: stock.ErrMustHaveTargetID,
		},
		{
			name: "should fail because product is not found",
			args: args{
				ctx: context.Background(),
				mov: stock.Movement{
					TargetID: id.ExternalFrom("bf593576-69eb-431b-a956-aaf1337e1f66"),
				},
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

			transactions := transactions
			if tc.opts.transactions != nil {
				transactions = tc.opts.transactions
			}

			tc.args.ctx = ctx
			wf := NewWorkflow(tc.opts.products, tc.opts.stocks, tc.opts.categories, transactions)
			mov, err := wf.AddToStock(tc.args.ctx, tc.args.mov)

			assert.ErrorIs(t, err, tc.wantErr)
			assert.Equal(t, tc.want, mov)
		})
	}
}
