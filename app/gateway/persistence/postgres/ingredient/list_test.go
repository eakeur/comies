package ingredient

import (
	"comies/app/core/entities/ingredient"
	"comies/app/core/entities/product"
	"comies/app/gateway/persistence/postgres/tests"
	"comies/app/sdk/types"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_actions_List(t *testing.T) {
	t.Parallel()

	type args struct {
		productID types.ID
	}
	cases := []struct {
		name    string
		args    args
		before  tests.Callback
		want    []ingredient.Ingredient
		wantErr error
	}{
		{
			name: "should return empty list of ingredients",
			args: args{
				productID: 1,
			},
			want: []ingredient.Ingredient{},
		},

		{
			name: "should return list with ingredients",
			before: func(ctx context.Context, d *tests.Database, t *testing.T) {
				_, err := d.InsertProducts(ctx, product.Product{
					ID:   1,
					Code: "PZF",
					Name: "Pizza de frango",
					Type: product.OutputType,
				}, product.Product{
					ID:   2,
					Code: "FGO",
					Name: "Frango",
					Type: product.InputType,
				}, product.Product{
					ID:   3,
					Code: "MSS",
					Name: "Massa",
					Type: product.InputType,
				}, product.Product{
					ID:   4,
					Code: "QJO",
					Name: "Queijo",
					Type: product.InputType,
				})
				if err != nil {
					t.Errorf("error inserting products: %v", err)
				}

				_, err = d.InsertIngredients(ctx,
					ingredient.Ingredient{ID: 1, ProductID: 1, IngredientID: 2, Quantity: 100},
					ingredient.Ingredient{ID: 2, ProductID: 1, IngredientID: 3, Quantity: 100},
					ingredient.Ingredient{ID: 3, ProductID: 1, IngredientID: 4, Quantity: 100},
				)
				if err != nil {
					t.Error(err)
				}
			},
			args: args{
				productID: 1,
			},
			want: []ingredient.Ingredient{
				{ID: 1, ProductID: 1, IngredientID: 2, Quantity: 100},
				{ID: 2, ProductID: 1, IngredientID: 3, Quantity: 100},
				{ID: 3, ProductID: 1, IngredientID: 4, Quantity: 100},
			},
		},
	}
	for _, tt := range cases {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctx, db := tests.FetchTestDB(t, tt.before)
			defer db.Drop()

			got, err := actions{db: db.Pool}.List(ctx, tt.args.productID)
			assert.ErrorIs(t, err, tt.wantErr)
			assert.Equal(t, tt.want, got)
		})
	}
}
