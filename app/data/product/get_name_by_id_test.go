package product

import (
	"comies/app/core/entities/product"
	"comies/app/core/types"
	"comies/app/gateway/persistence/postgres/tests"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_actions_GetNameByID(t *testing.T) {
	t.Parallel()

	type args struct {
		id types.ID
	}
	cases := []struct {
		name    string
		args    args
		want    string
		wantErr error
		before  tests.Callback
		after   tests.Callback
	}{
		{
			name: "should return name with id",
			args: args{
				id: 1,
			},
			want: "Product X",
			before: func(ctx context.Context, db *tests.Database, t *testing.T) {
				if _, err := db.InsertProducts(ctx, product.Product{
					ID:   1,
					Code: "PRDX",
					Name: "Product X",
					Type: product.OutputType,
					Sale: product.Sale{
						CostPrice:   10,
						SalePrice:   20,
						SaleUnit:    types.Unit,
						MinimumSale: 1,
					},
				}); err != nil {
					t.Error(err)
				}
			},
		},
		{
			name: "should return ErrNotFound error for nonexistent product",
			args: args{
				id: 1,
			},
			wantErr: product.ErrNotFound,
		},
	}
	for _, tt := range cases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctx, db := tests.FetchTestDB(t, tt.before)

			a := actions{db: db.Pool}
			got, err := a.GetNameByID(ctx, tt.args.id)
			assert.ErrorIs(t, err, tt.wantErr)
			assert.Equalf(t, tt.want, got, "GetByID(%v)", tt.args.id)
		})
	}
}
