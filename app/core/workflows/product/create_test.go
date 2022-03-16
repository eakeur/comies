package product

import (
	"context"
	"gomies/app/core/entities/category"
	"gomies/app/core/entities/product"
	"gomies/app/core/types/entity"
	"gomies/app/core/types/history"
	"gomies/app/core/types/id"
	"gomies/app/core/types/units"
	"gomies/app/shared/tests"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWorkflow_Create(t *testing.T) {
	const operation = "Workflows.Product.Create"
	t.Parallel()

	transactions := tests.GetFakeManagers().Transaction

	type args struct {
		ctx context.Context
		prd product.Product
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
			name: "should fail because category is not found",
			args: args{
				ctx: context.Background(),
				prd: product.Product{
					Code:               "PRD1",
					Name:               "Product 1",
					CategoryExternalID: id.ExternalFrom("c7e0b22e-95e5-4ca9-8820-d0e83553d982"),
					Sale: product.SaleInformation{
						Display:         "",
						Description:     "",
						Price:           20,
						Unit:            units.Unit,
						MinimumSale:     1,
						MaximumDiscount: 10,
					},
				},
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
		{
			name: "should return product created",
			args: args{
				ctx: context.Background(),
				prd: product.Product{
					Code:               "PRD1",
					Name:               "Product 1",
					CategoryExternalID: id.ExternalFrom("c7e0b22e-95e5-4ca9-8820-d0e83553d982"),
					Sale: product.SaleInformation{
						Display:         "",
						Description:     "",
						Price:           20,
						Unit:            units.Unit,
						MinimumSale:     1,
						MaximumDiscount: 10,
					},
				},
			},
			want: product.Product{
				Entity: entity.Entity{
					StoreExternalID: id.ExternalFrom("7a4ad106-f91d-4898-955d-91f0e7e93972"),
					StoreID:         1,
					Active:          true,
					History: history.History{
						Operation: operation,
						By:        id.ExternalFrom("7a4ad106-f91d-4898-955d-91f0e7e93971"),
					},
				},
				Code:               "PRD1",
				Name:               "Product 1",
				CategoryID:         1,
				CategoryExternalID: id.ExternalFrom("c7e0b22e-95e5-4ca9-8820-d0e83553d982"),
				Sale: product.SaleInformation{
					Display:         "",
					Description:     "",
					Price:           20,
					Unit:            units.Unit,
					MinimumSale:     1,
					MaximumDiscount: 10,
				},
			},
			opts: workflow{
				products: &product.ActionsMock{
					CreateFunc: func(_ context.Context, p product.Product) (product.Product, error) {
						return p, nil
					},
				},
				categories: &category.ActionsMock{
					GetFunc: func(contextMoqParam context.Context, external id.External) (category.Category, error) {
						return category.Category{
							Entity:      entity.Entity{ID: 1, ExternalID: id.ExternalFrom("c7e0b22e-95e5-4ca9-8820-d0e83553d982")},
							Name:        "Category 1",
							Code:        "CT1",
							Description: "... .... ... ...",
						}, nil
					},
				},
			},
		},
		{
			name: "should fail because product has no price",
			args: args{
				ctx: context.Background(),
				prd: product.Product{
					Code:               "PRD1",
					Name:               "Product 1",
					CategoryExternalID: id.ExternalFrom("c7e0b22e-95e5-4ca9-8820-d0e83553d982"),
					Sale: product.SaleInformation{
						Display:         "",
						Description:     "",
						Unit:            units.Unit,
						MinimumSale:     1,
						MaximumDiscount: 10,
					},
				},
			},
			wantErr: product.ErrInvalidSalePrice,
		},
		{
			name: "should fail because product with code already exists",
			args: args{
				ctx: context.Background(),
				prd: product.Product{
					Code:               "PRD1",
					Name:               "Product 1",
					CategoryExternalID: id.ExternalFrom("c7e0b22e-95e5-4ca9-8820-d0e83553d982"),
					Sale: product.SaleInformation{
						Display:         "",
						Description:     "",
						Price:           20,
						Unit:            units.Unit,
						MinimumSale:     1,
						MaximumDiscount: 10,
					},
				},
			},
			wantErr: product.ErrAlreadyExists,
			opts: workflow{
				products: &product.ActionsMock{
					CreateFunc: func(contextMoqParam context.Context, p product.Product) (product.Product, error) {
						return p, product.ErrAlreadyExists
					},
				},
				categories: &category.ActionsMock{
					GetFunc: func(contextMoqParam context.Context, external id.External) (category.Category, error) {
						return category.Category{
							Entity:      entity.Entity{ID: 1, ExternalID: id.ExternalFrom("c7e0b22e-95e5-4ca9-8820-d0e83553d982")},
							Name:        "Category 1",
							Code:        "CT1",
							Description: "... .... ... ...",
						}, nil
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

			wf := NewWorkflow(tc.opts.products, tc.opts.stocks, tc.opts.categories, transactions)
			mov, err := wf.Create(tc.args.ctx, tc.args.prd)

			assert.ErrorIs(t, err, tc.wantErr)
			assert.Equal(t, tc.want, mov)
		})
	}
}
