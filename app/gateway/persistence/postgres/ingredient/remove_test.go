package ingredient

import (
	"context"
	"github.com/stretchr/testify/assert"
	"gomies/app/core/entities/ingredient"
	"gomies/app/core/entities/product"
	"gomies/app/gateway/persistence/postgres/tests"
	"gomies/app/sdk/fault"
	"gomies/app/sdk/types"
	"testing"
)

func Test_actions_Remove(t *testing.T) {
	t.Parallel()

	type args struct {
		itemID types.ID
	}
	cases := []struct {
		name    string
		args    args
		wantErr error
		before  tests.Callback
		after   tests.Callback
	}{
		{
			name: "should delete ingredient successfully",
			args: args{
				itemID: 3,
			},
			before: func(ctx context.Context, d *tests.Database, t *testing.T) {
				_, err := d.InsertProducts(ctx, product.Product{
					ID:     1,
					Active: true,
					Code:   "PZF",
					Name:   "Pizza de frango",
					Type:   product.OutputType,
				}, product.Product{
					ID:     2,
					Active: true,
					Code:   "FGO",
					Name:   "Frango",
					Type:   product.InputType,
				}, product.Product{
					ID:     3,
					Active: true,
					Code:   "MSS",
					Name:   "Massa",
					Type:   product.InputType,
				}, product.Product{
					ID:     4,
					Active: true,
					Code:   "QJO",
					Name:   "Queijo",
					Type:   product.InputType,
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
			after: func(ctx context.Context, d *tests.Database, t *testing.T) {
				d.CheckValue(ctx, "select count(id) from ingredients", int64(2))
			},
		},
		{
			name: "should fail for nonexistent ingredient",
			args: args{
				itemID: 1,
			},
			wantErr: fault.ErrNotFound,
		},
	}
	for _, tt := range cases {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctx, db := tests.FetchTestTX(t, tt.before)
			defer db.Drop(tt.after)

			a := actions{}
			err := a.Remove(ctx, tt.args.itemID)
			assert.ErrorIs(t, err, tt.wantErr)
		})
	}
}
