package product

import (
	"context"
	"github.com/stretchr/testify/assert"
	"gomies/app/core/entities/catalog/product"
	"testing"
)

func TestWorkflow_AddIngredient(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	type (
		args struct {
			key        product.Key
			ingredient product.Ingredient
		}

		opts struct {
			products *product.ActionsMock
		}

		test struct {
			name    string
			args    args
			opts    opts
			want    product.Ingredient
			wantErr error
		}
	)

	cases := []test{
		{
			name: "should return ingredient created",
			args: args{
				key: product.Key{ID: idExample1},
				ingredient: product.Ingredient{
					Quantity:             1,
					ProductID:            1,
					ProductExternalID:    idExample1,
					IngredientID:         2,
					IngredientExternalID: idExample2,
				},
			},
			want: product.Ingredient{
				Quantity:             1,
				ProductID:            1,
				ProductExternalID:    idExample1,
				IngredientID:         2,
				IngredientExternalID: idExample2,
			},
			opts: opts{
				products: &product.ActionsMock{
					SaveIngredientsFunc: func(ctx context.Context, productKey product.Key, ingredients ...product.Ingredient) ([]product.Ingredient, error) {
						return ingredients, nil
					},
				},
			},
		},
		{
			name: "should return ingredient created with product id not set",
			args: args{
				key: product.Key{ID: idExample1},
				ingredient: product.Ingredient{
					Quantity:             1,
					IngredientID:         2,
					IngredientExternalID: idExample2,
				},
			},
			want: product.Ingredient{
				Quantity:             1,
				ProductExternalID:    idExample1,
				IngredientID:         2,
				IngredientExternalID: idExample2,
			},
			opts: opts{
				products: &product.ActionsMock{
					SaveIngredientsFunc: func(ctx context.Context, productKey product.Key, ingredients ...product.Ingredient) ([]product.Ingredient, error) {
						return ingredients, nil
					},
				},
			},
		},
		{
			name: "should fail ErrInvalidIngredient because ingredient was not set",
			args: args{
				key: product.Key{ID: idExample1},
				ingredient: product.Ingredient{
					Quantity:          1,
					ProductExternalID: idExample1,
				},
			},
			want:    product.Ingredient{},
			wantErr: product.ErrInvalidIngredient,
		},
		{
			name: "should fail ErrInvalidIngredientQuantity because quantity was not set",
			args: args{
				key: product.Key{ID: idExample1},
				ingredient: product.Ingredient{
					IngredientExternalID: idExample2,
				},
			},
			want:    product.Ingredient{},
			wantErr: product.ErrInvalidIngredientQuantity,
		},
	}

	for _, c := range cases {
		c := c

		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			wf := NewWorkflow(c.opts.products, nil, nil)
			ingredient, err := wf.AddIngredient(ctx, c.args.key, c.args.ingredient)

			assert.ErrorIs(t, err, c.wantErr)
			assert.Equal(t, c.want, ingredient)

		})
	}

}
