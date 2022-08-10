package ingredient

import (
	"comies/app/core/entities/ingredient"
	"comies/app/core/entities/product"
	"comies/app/core/throw"
	"comies/app/gateway/persistence/postgres/tests"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_actions_Create(t *testing.T) {
	t.Parallel()

	type args struct {
		i ingredient.Ingredient
	}
	cases := []struct {
		name    string
		args    args
		want    ingredient.Ingredient
		wantErr error
		before  tests.Callback
		after   tests.Callback
	}{
		{
			name: "should return ingredient",
			args: args{i: ingredient.Ingredient{
				ID:           1,
				ProductID:    1,
				IngredientID: 2,
				Quantity:     100,
				Optional:     false,
			}},
			want: ingredient.Ingredient{
				ID:           1,
				ProductID:    1,
				IngredientID: 2,
				Quantity:     100,
				Optional:     false,
			},
			before: func(ctx context.Context, db *tests.Database, t *testing.T) {
				_, err := db.InsertProducts(ctx, product.Product{
					ID:   1,
					Code: "PZF",
					Name: "Pizza de frango",
					Type: product.OutputType,
				}, product.Product{
					ID:   2,
					Code: "FGO",
					Name: "Frango",
					Type: product.InputType,
				})
				if err != nil {
					t.Errorf("error inserting products: %v", err)
				}
			},
			after: func(ctx context.Context, db *tests.Database, t *testing.T) {
				db.CheckValue(t, ctx, "select count(id) from ingredients", int64(1))
			},
		},
		{
			name: "should fail for nonexistent product in ingredient id field",
			args: args{i: ingredient.Ingredient{
				ID:           1,
				ProductID:    1,
				IngredientID: 3,
				Quantity:     100,
				Optional:     false,
			}},
			wantErr: throw.ErrNotFound,
			before: func(ctx context.Context, db *tests.Database, t *testing.T) {
				_, err := db.InsertProducts(ctx, product.Product{
					ID:   1,
					Code: "PZF",
					Name: "Pizza de frango",
					Type: product.OutputType,
				})
				if err != nil {
					t.Errorf("error inserting products: %v", err)
				}
			},
		},
		{
			name: "should fail for nonexistent product in product id field",
			args: args{i: ingredient.Ingredient{
				ID:           1,
				ProductID:    3,
				IngredientID: 1,
				Quantity:     100,
				Optional:     false,
			}},
			wantErr: throw.ErrNotFound,
			before: func(ctx context.Context, db *tests.Database, t *testing.T) {
				_, err := db.InsertProducts(ctx, product.Product{
					ID:   1,
					Code: "FGO",
					Name: "Frango",
					Type: product.InputType,
				})
				if err != nil {
					t.Errorf("error inserting products: %v", err)
				}
			},
		},
		{
			name: "should fail for already existent ingredient relation",
			args: args{i: ingredient.Ingredient{
				ID:           2,
				ProductID:    1,
				IngredientID: 2,
				Quantity:     100,
				Optional:     false,
			}},
			wantErr: throw.ErrAlreadyExists,
			before: func(ctx context.Context, db *tests.Database, t *testing.T) {
				_, err := db.InsertProducts(ctx, product.Product{
					ID:   1,
					Code: "PZF",
					Name: "Pizza de frango",
					Type: product.OutputType,
				}, product.Product{
					ID:   2,
					Code: "FGO",
					Name: "Frango",
					Type: product.InputType,
				})
				if err != nil {
					t.Errorf("error inserting products: %v", err)
				}

				_, err = db.InsertIngredients(ctx, ingredient.Ingredient{
					ID:           1,
					ProductID:    1,
					IngredientID: 2,
					Quantity:     100,
					Optional:     false,
				})
				if err != nil {
					t.Errorf("error inserting ingredients: %v", err)
				}

			},
		},
	}
	for _, tt := range cases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctx, _ := tests.FetchTestTX(t, tt.before, tt.after)

			a := actions{}
			got, err := a.Create(ctx, tt.args.i)
			assert.ErrorIs(t, err, tt.wantErr)
			assert.Equalf(t, tt.want, got, "Create(%v)", tt.args.i)
		})
	}
}
