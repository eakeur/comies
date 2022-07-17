package menu

import (
	"comies/app/gateway/api/gen/menu"
	"comies/app/gateway/api/tests"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"testing"
)

func TestProductCrud(t *testing.T) {
	t.Parallel()
	ctx, client := NewClient(t)

	var id int64
	t.Run("should create a product successfully", func(t *testing.T) {
		prd, err := client.CreateProduct(ctx, &menu.CreateProductRequest{
			Code:         "COCZ",
			Name:         "Coca Cola zero 2L",
			Type:         menu.ProductType_OUTPUT,
			Cost:         550,
			Price:        800,
			Unit:         "un",
			Minimum:      1,
			StockMinimum: 10,
			StockMaximum: 100,
			Location:     "Fridge",
		})
		if err != nil {
			t.Error(err)
		}

		id = prd.Id
	})

	t.Run("should fail creating a product for repeated code", func(t *testing.T) {
		_, err := client.CreateProduct(ctx, &menu.CreateProductRequest{
			Code:         "COCZ",
			Name:         "Coca Cola zero 3L",
			Type:         menu.ProductType_OUTPUT,
			Cost:         750,
			Price:        1000,
			Unit:         "un",
			Minimum:      1,
			StockMinimum: 5,
			StockMaximum: 20,
			Location:     "Fridge",
		})

		tests.ExpectError(t, err, codes.AlreadyExists, "Ops! The code assigned to this product seems to belong to another product already")
	})

	t.Run("should fetch product by id", func(t *testing.T) {
		prd, err := client.GetProductByID(ctx, &menu.GetProductByIDRequest{Id: id})
		if err != nil {
			t.Error(err)
		}

		assert.EqualValues(t, &menu.GetProductByIDResponse{
			Id:      id,
			Code:    "COCZ",
			Name:    "Coca Cola zero 2L",
			Type:    menu.ProductType_OUTPUT,
			Cost:    550,
			Price:   800,
			Minimum: 1,
		}, &menu.GetProductByIDResponse{
			Id:           prd.Id,
			Code:         prd.Code,
			Name:         prd.Name,
			Type:         prd.Type,
			Cost:         prd.Cost,
			Price:        prd.Price,
			Unit:         prd.Unit,
			Minimum:      prd.Minimum,
			StockMinimum: prd.StockMinimum,
			StockMaximum: prd.StockMaximum,
			Location:     prd.Location,
		})
	})

	t.Run("should fail fetching product by id for nonexistent product", func(t *testing.T) {
		_, err := client.GetProductByID(ctx, &menu.GetProductByIDRequest{Id: 0})
		tests.ExpectError(t, err, codes.NotFound, "Ops! This product does not exist or could not be found")
	})

	t.Run("should update product by id", func(t *testing.T) {
		_, err := client.UpdateProduct(ctx, &menu.UpdateProductRequest{
			Id:      id,
			Code:    "COCXT",
			Name:    "Coca Cola zero 2XL",
			Type:    menu.ProductType_OUTPUT,
			Cost:    600,
			Price:   850,
			Minimum: 1,
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

		assert.EqualValues(t, &menu.GetProductByIDResponse{
			Id:      id,
			Code:    "COCXT",
			Name:    "Coca Cola zero 2XL",
			Type:    menu.ProductType_OUTPUT,
			Cost:    600,
			Price:   850,
			Minimum: 1,
		}, &menu.GetProductByCodeResponse{
			Id:           prd.Id,
			Code:         prd.Code,
			Name:         prd.Name,
			Type:         prd.Type,
			Cost:         prd.Cost,
			Price:        prd.Price,
			Unit:         prd.Unit,
			Minimum:      prd.Minimum,
			StockMinimum: prd.StockMinimum,
			StockMaximum: prd.StockMaximum,
			Location:     prd.Location,
		})
	})

	t.Run("should remove product by id", func(t *testing.T) {
		_, err := client.RemoveProduct(ctx, &menu.RemoveProductRequest{Id: id})
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("should fail deletion of nonexistent product", func(t *testing.T) {
		_, err := client.RemoveProduct(ctx, &menu.RemoveProductRequest{Id: id})
		tests.ExpectError(t, err, codes.NotFound, "Ops! This product does not exist or could not be found")
	})
}
