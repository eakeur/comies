package menu

import (
	"comies/app/gateway/api/gen/menu"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProductCrud(t *testing.T) {
	t.Parallel()
	ctx, client := NewClient(t)

	var id int64
	t.Run("should create a product successfully", func(t *testing.T) {
		prd, err := client.CreateProduct(ctx, &menu.CreateProductRequest{
			Product: &menu.Product{
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
			},
		})
		if err != nil {
			t.Error(err)
		}

		id = prd.Id
	})

	t.Run("should fetch product by id", func(t *testing.T) {
		prd, err := client.GetProductById(ctx, &menu.GetProductByIdRequest{Id: id})
		if err != nil {
			t.Error(err)
		}

		assert.EqualValues(t, &menu.GetProductByIdResponse{
			Product: &menu.Product{
				Id:      id,
				Code:    "COCZ",
				Name:    "Coca Cola zero 2L",
				Type:    menu.ProductType_PRODUCT_TYPE_OUTPUT,
				Cost:    550,
				Price:   800,
				Minimum: 1,
			},
		}, &menu.GetProductByIdResponse{
			Product: &menu.Product{
				Id:           prd.Product.Id,
				Code:         prd.Product.Code,
				Name:         prd.Product.Name,
				Type:         prd.Product.Type,
				Cost:         prd.Product.Cost,
				Price:        prd.Product.Price,
				Unit:         prd.Product.Unit,
				Minimum:      prd.Product.Minimum,
				StockMinimum: prd.Product.StockMinimum,
				StockMaximum: prd.Product.StockMaximum,
				Location:     prd.Product.Location,
			},
		})
	})

	t.Run("should update product by id", func(t *testing.T) {
		_, err := client.UpdateProduct(ctx, &menu.UpdateProductRequest{
			Product: &menu.Product{
				Id:      id,
				Code:    "COCXT",
				Name:    "Coca Cola zero 2XL",
				Type:    menu.ProductType_PRODUCT_TYPE_OUTPUT,
				Cost:    600,
				Price:   850,
				Minimum: 1,
			},
		})
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("should fetch updated product by code", func(t *testing.T) {
		prd, err := client.GetProductByCode(ctx, &menu.GetProductByCodeRequest{Code: "COCXT"})
		if err != nil {
			t.Error(err)
		}

		assert.EqualValues(t, &menu.GetProductByIdResponse{
			Product: &menu.Product{
				Id:      id,
				Code:    "COCXT",
				Name:    "Coca Cola zero 2XL",
				Type:    menu.ProductType_PRODUCT_TYPE_OUTPUT,
				Cost:    600,
				Price:   850,
				Minimum: 1,
			},
		}, &menu.GetProductByCodeResponse{
			Product: &menu.Product{
				Id:           prd.Product.Id,
				Code:         prd.Product.Code,
				Name:         prd.Product.Name,
				Type:         prd.Product.Type,
				Cost:         prd.Product.Cost,
				Price:        prd.Product.Price,
				Unit:         prd.Product.Unit,
				Minimum:      prd.Product.Minimum,
				StockMinimum: prd.Product.StockMinimum,
				StockMaximum: prd.Product.StockMaximum,
				Location:     prd.Product.Location,
			},
		})
	})

	t.Run("should remove product by id", func(t *testing.T) {
		_, err := client.RemoveProduct(ctx, &menu.RemoveProductRequest{Id: id})
		if err != nil {
			t.Error(err)
		}
	})
}
