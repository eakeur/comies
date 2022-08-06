package tests

import (
	"comies/app/core/entities/product"
	"comies/app/gateway/api/handler"
	"comies/app/gateway/api/menu"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestProductIntegration(t *testing.T) {
	t.Parallel()

	var (
		app      = NewTestApp(t, "/menu/products")
		products []menu.GetProductByKeyResponse
	)

	t.Run("should create all products successfully", func(t *testing.T) {
		for _, prod := range []menu.CreateProductRequest{
			{
				Code:            "COKL",
				Name:            "Coke 1L",
				Type:            product.OutputType,
				CostPrice:       800,
				SalePrice:       1000,
				SaleUnit:        "un",
				MinimumSale:     1,
				MaximumQuantity: 100,
				MinimumQuantity: 0,
				Location:        "Fridge",
			},
			{
				Code:            "COKC",
				Name:            "Coke Can",
				Type:            product.OutputType,
				CostPrice:       500,
				SalePrice:       800,
				SaleUnit:        "un",
				MinimumSale:     1,
				MaximumQuantity: 100,
				MinimumQuantity: 0,
				Location:        "Fridge",
			},
			{
				Code:            "CHIKP",
				Name:            "Chicken pieces",
				Type:            product.InputType,
				CostPrice:       60,
				SaleUnit:        "g",
				MaximumQuantity: 10000,
				MinimumQuantity: 0,
				Location:        "Fridge",
			},
		} {
			var payload handler.Response
			var data menu.CreateProductResponse
			payload.Data = &data

			response := app.Request(t, http.MethodPost, "", RequestInput{body: prod}).To(&payload).Run()
			if response.data.StatusCode != http.StatusCreated {
				t.Errorf("Failed to create product %s: %v", prod.Code, payload.Error)
				return
			}

			products = append(products, menu.GetProductByKeyResponse{
				ID:              data.ID,
				Code:            prod.Code,
				Name:            prod.Name,
				Type:            prod.Type,
				CostPrice:       prod.CostPrice,
				SalePrice:       prod.SalePrice,
				SaleUnit:        prod.SaleUnit,
				MinimumSale:     prod.MinimumSale,
				MaximumQuantity: prod.MaximumQuantity,
				MinimumQuantity: prod.MinimumQuantity,
				Location:        prod.Location,
			})
		}
	})

	t.Run("should fetch products by id successfully", func(t *testing.T) {
		for idx, id := range []string{
			products[0].ID, products[1].ID, products[2].ID,
		} {

			var payload handler.Response
			var data menu.GetProductByKeyResponse
			payload.Data = &data

			response := app.Request(t, http.MethodGet, fmt.Sprintf("/%s", id)).To(&payload).Run()
			if response.data.StatusCode != http.StatusOK {
				t.Errorf("Failed to fetch product %s: %v", id, payload.Error)
				return
			}

			assert.Equal(t, products[idx], data, "the product received was not as the expected for %s", idx)
		}
	})

	t.Run("should list products with filters successfully", func(t *testing.T) {
		for _, r := range []struct {
			filter map[string]string
			want   []menu.ListProductsResponse
		}{
			{
				filter: map[string]string{
					"code": "COK",
				},
				want: []menu.ListProductsResponse{
					{
						Code: "COKC",
						Name: "Coke Can",
						Type: product.OutputType,
					},
					{
						Code: "COKL",
						Name: "Coke 1L",
						Type: product.OutputType,
					},
				},
			},
			{
				filter: map[string]string{
					"name": "pieces",
				},
				want: []menu.ListProductsResponse{
					{
						Code: "CHIKP",
						Name: "Chicken pieces",
						Type: product.InputType,
					},
				},
			},
			{
				filter: map[string]string{
					"type": "30",
				},
				want: []menu.ListProductsResponse{
					{
						Code: "CHIKP",
						Name: "Chicken pieces",
						Type: product.InputType,
					},
				},
			},
		} {
			var payload handler.Response
			var data []menu.ListProductsResponse
			payload.Data = &data

			response := app.Request(t, http.MethodGet, "", RequestInput{query: r.filter}).To(&payload).Run()
			if response.data.StatusCode != http.StatusOK {
				t.Errorf("Failed to fetch products: %v", payload.Error)
				return
			}

			for i, g := range data {
				r.want[i].ID = g.ID
			}

			assert.Equal(t, r.want, data, "the list received was not the same as expected")
		}
	})

	t.Run("should update product successfully", func(t *testing.T) {
		var (
			prod = products[0]
			path = fmt.Sprintf("/%s", products[0].ID)
		)

		prod.Location = "In the supermarket :)"
		prod.Code = "ANY"
		prod.Name = "Any product"
		response := app.Request(t, http.MethodPut, path, RequestInput{
			body: menu.UpdateProductRequest{
				Code:            prod.Code,
				Name:            prod.Name,
				Type:            prod.Type,
				CostPrice:       prod.CostPrice,
				SalePrice:       prod.SalePrice,
				SaleUnit:        prod.SaleUnit,
				MinimumSale:     prod.MinimumSale,
				MaximumQuantity: prod.MaximumQuantity,
				MinimumQuantity: prod.MinimumQuantity,
				Location:        prod.Location,
			},
		}).Run()
		if response.data.StatusCode != http.StatusNoContent {
			t.Errorf("Failed to update product %s: %v", products[0].Code, response.dump)
			return
		}

		var payload handler.Response
		var data menu.GetProductByKeyResponse
		payload.Data = &data

		response = app.Request(t, http.MethodGet, path).To(&payload).Run()
		if response.data.StatusCode != http.StatusOK {
			t.Errorf("Failed to fetch updated product %s: %v", products[0].Code, payload.Error)
			return
		}

		products[0] = prod
		assert.Equal(t, products[0], data, "the product received was not as the expected for %s", prod.Code)
	})

	t.Run("should delete product successfully", func(t *testing.T) {
		var path = fmt.Sprintf("/%s", products[0].ID)

		response := app.Request(t, http.MethodDelete, path).Run()
		if response.data.StatusCode != http.StatusNoContent {
			t.Errorf("Failed to delete product %s: %v", products[0].Code, response.dump)
			return
		}
	})

	t.Run("should fail fetching deleted product", func(t *testing.T) {
		response := app.Request(t, http.MethodGet, fmt.Sprintf("/%s", products[0].ID)).Run()
		if response.data.StatusCode != http.StatusNotFound {
			t.Errorf("Dleted product remains!: %v", response.dump)
			return
		}
	})
}
