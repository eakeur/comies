package menu

import (
	"comies/app/gateway/api/gen/menu"
	"comies/app/sdk/throw"
	"errors"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

		if err == nil {
			t.Errorf("test failed because it was expected to receive error")
		}

		st, _ := status.FromError(err)
		c := st.Code()
		msg := st.Message()
		if c != codes.AlreadyExists && msg != "The code provided is already assigned to another product" {
			t.Errorf("test failed because it was expected to receive AlreadyExists error, but received: %v - %v", c, msg)
		}
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
			Active:       prd.Active,
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
		if err == nil || !errors.Is(err, throw.ErrNotFound) {
			t.Errorf("test failed because it was expected to receive ErrNotFound: %v", err)
		}
	})
}
