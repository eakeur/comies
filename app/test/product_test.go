package test

import (
	"comies/app/gateway/api/gen/menu"
	"comies/app/test/tests"
	"context"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/timestamppb"
	"testing"
)

func TestMenuFunctions(t *testing.T) {
	t.Parallel()
	srv := tests.NewTestApp(t)
	ctx, client := context.Background(), menu.NewMenuClient(srv)

	var (
		product = &menu.Product{
			Code:         "COCZ",
			Name:         "Coca Cola zero 2L",
			Type:         menu.ProductType_PRODUCT_TYPE_OUTPUT,
			Cost:         550,
			Price:        800,
			Unit:         "un",
			Minimum:      1,
			StockMinimum: 10,
			StockMaximum: 100,
			Location:     "Fridge",
		}

		waterInputProduct = &menu.Product{
			Code:         "WAT",
			Name:         "Water",
			Type:         menu.ProductType_PRODUCT_TYPE_INPUT,
			Cost:         100,
			Unit:         "L",
			StockMinimum: 10,
			StockMaximum: 1000,
			Location:     "Fridge",
		}

		sugarInputProduct = &menu.Product{
			Code:         "SUG",
			Name:         "Sugar",
			Type:         menu.ProductType_PRODUCT_TYPE_INPUT,
			Cost:         250,
			Unit:         "mg",
			StockMinimum: 10,
			StockMaximum: 1000,
			Location:     "Depot",
		}

		caramelInputProduct = &menu.Product{
			Code:         "CARSUG",
			Name:         "Caramel Sugar",
			Type:         menu.ProductType_PRODUCT_TYPE_INPUT,
			Cost:         500,
			Unit:         "ml",
			StockMinimum: 10,
			StockMaximum: 1000,
			Location:     "Depot",
		}
	)
	t.Run("should create products successfully", func(t *testing.T) {
		for _, prod := range []*menu.Product{product, waterInputProduct, sugarInputProduct, caramelInputProduct} {
			prd, err := client.CreateProduct(ctx, &menu.CreateProductRequest{
				Product: prod,
			})
			if err != nil {
				t.Error(err)
			}

			prod.Id = prd.Id
		}
	})

	t.Run("should create ingredient relations successfully", func(t *testing.T) {
		for _, ing := range []*menu.Ingredient{
			{
				ProductID:    product.Id,
				IngredientID: sugarInputProduct.Id,
				Quantity:     250,
			}, {
				ProductID:    product.Id,
				IngredientID: caramelInputProduct.Id,
				Quantity:     100,
			}, {
				ProductID:    product.Id,
				IngredientID: waterInputProduct.Id,
				Quantity:     500,
			},
		} {
			_, err := client.CreateIngredient(ctx, &menu.CreateIngredientRequest{
				Ingredient: ing,
			})
			if err != nil {
				t.Error(err)
			}
		}
	})

	t.Run("should create movements for products", func(t *testing.T) {
		for i, mv := range []*menu.Movement{
			{
				ProductID: waterInputProduct.Id,
				Type:      menu.MovementType_MOVEMENT_TYPE_INPUT_MOVEMENT_TYPE,
				Date:      timestamppb.Now(),
				Quantity:  100,
				Value:     100,
			},
			{
				ProductID: waterInputProduct.Id,
				Type:      menu.MovementType_MOVEMENT_TYPE_INPUT_MOVEMENT_TYPE,
				Date:      timestamppb.Now(),
				Quantity:  200,
				Value:     100,
			},
			{
				ProductID: waterInputProduct.Id,
				Type:      menu.MovementType_MOVEMENT_TYPE_INPUT_MOVEMENT_TYPE,
				Date:      timestamppb.Now(),
				Quantity:  300,
				Value:     100,
			},
			{
				ProductID: waterInputProduct.Id,
				Type:      menu.MovementType_MOVEMENT_TYPE_INPUT_MOVEMENT_TYPE,
				Date:      timestamppb.Now(),
				Quantity:  400,
				Value:     100,
			}, // Should sum 1000
			{
				ProductID: sugarInputProduct.Id,
				Type:      menu.MovementType_MOVEMENT_TYPE_INPUT_MOVEMENT_TYPE,
				Date:      timestamppb.Now(),
				Quantity:  10,
				Value:     250,
			},
			{
				ProductID: sugarInputProduct.Id,
				Type:      menu.MovementType_MOVEMENT_TYPE_INPUT_MOVEMENT_TYPE,
				Date:      timestamppb.Now(),
				Quantity:  44,
				Value:     250,
			},
			{
				ProductID: sugarInputProduct.Id,
				Type:      menu.MovementType_MOVEMENT_TYPE_INPUT_MOVEMENT_TYPE,
				Date:      timestamppb.Now(),
				Quantity:  100,
				Value:     250,
			}, // Should sum 154
			{
				ProductID: caramelInputProduct.Id,
				Type:      menu.MovementType_MOVEMENT_TYPE_INPUT_MOVEMENT_TYPE,
				Date:      timestamppb.Now(),
				Quantity:  500,
				Value:     500,
			},
		} {
			_, err := client.CreateMovement(ctx, &menu.CreateMovementRequest{
				Movement: mv,
			})
			if err != nil {
				t.Errorf("error creating movements for product[%d] %d:  %v", i, mv.ProductID, err)
			}
		}
	})

	t.Run("should get products stock balance", func(t *testing.T) {
		bal, err := client.GetProductBalance(ctx, &menu.GetProductBalanceRequest{Id: waterInputProduct.Id})
		if err != nil {
			t.Errorf("error getting balance for water product %d: %v", waterInputProduct.Id, err)
		}
		assert.EqualValues(t, 1000, bal.Balance)

		bal, err = client.GetProductBalance(ctx, &menu.GetProductBalanceRequest{Id: sugarInputProduct.Id})
		if err != nil {
			t.Errorf("error getting balance for sugar product %d: %v", sugarInputProduct.Id, err)
		}
		assert.EqualValues(t, 154, bal.Balance)

		bal, err = client.GetProductBalance(ctx, &menu.GetProductBalanceRequest{Id: caramelInputProduct.Id})
		if err != nil {
			t.Errorf("error getting balance for caramel product %d: %v", caramelInputProduct.Id, err)
		}
		assert.EqualValues(t, 500, bal.Balance)
	})

	t.Run("should fetch product by id", func(t *testing.T) {
		prd, err := client.GetProductById(ctx, &menu.GetProductByIdRequest{Id: product.Id})
		if err != nil {
			t.Error(err)
		}

		assert.EqualValues(t, (&menu.Product{
			Id:      product.Id,
			Code:    product.Code,
			Name:    product.Name,
			Type:    product.Type,
			Cost:    product.Cost,
			Price:   product.Price,
			Minimum: product.Minimum,
		}).String(), (&menu.Product{
			Id:      prd.Product.Id,
			Code:    prd.Product.Code,
			Name:    prd.Product.Name,
			Type:    prd.Product.Type,
			Cost:    prd.Product.Cost,
			Price:   prd.Product.Price,
			Minimum: prd.Product.Minimum,
		}).String())
	})

	t.Run("should update product by id", func(t *testing.T) {
		product.Code = "COCXT"
		product.Name = "Coca Cola zero 2XL"
		product.Type = menu.ProductType_PRODUCT_TYPE_OUTPUT
		product.Cost = 600
		product.Price = 850
		product.Minimum = 1
		_, err := client.UpdateProduct(ctx, &menu.UpdateProductRequest{
			Product: product,
		})
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("should remove product by id", func(t *testing.T) {
		_, err := client.RemoveProduct(ctx, &menu.RemoveProductRequest{Id: product.Id})
		if err != nil {
			t.Error(err)
		}
	})
}
