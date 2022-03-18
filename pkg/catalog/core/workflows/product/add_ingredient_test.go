package product

import (
	"context"
	"github.com/stretchr/testify/assert"
	"gomies/pkg/catalog/core/entities/product"
	"gomies/pkg/sdk/tests"
	"gomies/pkg/sdk/types"
	"testing"
)

func TestWorkflow_AddIngredient(t *testing.T) {
	const operation = "Workflows.Product.AddIngredient"
	t.Parallel()

	ctx := tests.WorkflowContext(idExample1, idExample2)
	managers := tests.Managers()

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
			wantKey product.Key
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
				Entity: types.Entity{
					History: types.History{
						By:        idExample1,
						Operation: operation,
					},
				},
				Quantity:             1,
				ProductID:            1,
				ProductExternalID:    idExample1,
				IngredientID:         2,
				IngredientExternalID: idExample2,
				Store:                types.Store{StoreID: idExample2},
			},
			wantKey: product.Key{
				ID:    idExample1,
				Store: types.Store{StoreID: idExample2},
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
				Entity: types.Entity{
					History: types.History{
						By:        idExample1,
						Operation: operation,
					},
				},
				Quantity:             1,
				ProductExternalID:    idExample1,
				IngredientID:         2,
				IngredientExternalID: idExample2,
				Store:                types.Store{StoreID: idExample2},
			},
			wantKey: product.Key{
				ID:    idExample1,
				Store: types.Store{StoreID: idExample2},
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

			wf := NewWorkflow(c.opts.products, nil, nil, managers.Transactions)
			ingredient, err := wf.AddIngredient(ctx, c.args.key, c.args.ingredient)

			assert.ErrorIs(t, err, c.wantErr)
			assert.Equal(t, c.want, ingredient)

			if err == nil && c.wantErr == nil {
				assert.Equal(t, c.wantKey, c.opts.products.SaveIngredientsCalls()[0].ProductKey)
			}

		})
	}

}
