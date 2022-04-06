package product

import (
	"context"
	"gomies/app/core/entities/catalog/product"
	"gomies/pkg/sdk/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWorkflow_AddIngredient(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	fakeID := types.UID("1bdcafba-9deb-48b4-8a0e-ecea4c99b0e3")

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
				key: product.Key{ID: fakeID},
				ingredient: product.Ingredient{
					Quantity:             1,
					ProductID:            1,
					ProductExternalID:    fakeID,
					IngredientID:         2,
					IngredientExternalID: fakeID,
				},
			},
			want: product.Ingredient{
				Quantity:             1,
				ProductID:            1,
				ProductExternalID:    fakeID,
				IngredientID:         2,
				IngredientExternalID: fakeID,
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
				key: product.Key{ID: fakeID},
				ingredient: product.Ingredient{
					Quantity:             1,
					IngredientID:         2,
					IngredientExternalID: fakeID,
				},
			},
			want: product.Ingredient{
				Quantity:             1,
				ProductExternalID:    fakeID,
				IngredientID:         2,
				IngredientExternalID: fakeID,
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
				key: product.Key{ID: fakeID},
				ingredient: product.Ingredient{
					Quantity:          1,
					ProductExternalID: fakeID,
				},
			},
			want:    product.Ingredient{},
			wantErr: product.ErrInvalidIngredient,
		},
		{
			name: "should fail ErrInvalidIngredientQuantity because quantity was not set",
			args: args{
				key: product.Key{ID: fakeID},
				ingredient: product.Ingredient{
					IngredientExternalID: fakeID,
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

			wf := workflow{
				products: c.opts.products,
			}

			ingredient, err := wf.CreateIngredient(ctx, c.args.key, c.args.ingredient)

			assert.ErrorIs(t, err, c.wantErr)
			assert.Equal(t, c.want, ingredient)

		})
	}

}
