package tests

import (
	"comies/app/core/entities/product"
	"comies/app/gateway/api/handler"
	"comies/app/gateway/api/menu"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"sync"
	"testing"
)

func TestProductIntegration(t *testing.T) {
	t.Parallel()

	var (
		app      = NewTestApp(t, "/menu/products")
		products = []menu.GetProductByKeyResponse{
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
		}
	)

	t.Run("should create all products successfully", func(t *testing.T) {
		wg := sync.WaitGroup{}
		for i, prod := range products {
			wg.Add(1)
			go func(i int, prod menu.GetProductByKeyResponse) {
				defer wg.Done()

				var (
					data menu.CreateProductResponse

					payload = handler.Response{
						Data: &data,
					}
					input = RequestInput{
						body: menu.CreateProductRequest{
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
					}
				)

				response := app.Request(t, http.MethodPost, "", input).To(&payload).Run()
				if response.data.StatusCode != http.StatusCreated {
					t.Errorf("Failed to create product %s: %v", prod.Code, payload.Error)
					return
				}

				products[i].ID = data.ID
			}(i, prod)
		}
		wg.Wait()
	})

	t.Run("should update product successfully", func(t *testing.T) {
		products[0].Name = "Coke 1 Liter"

		var (
			prod = products[0]
			path = fmt.Sprintf("/%s", prod.ID)
		)

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
	})

	t.Run("should fetch products by id successfully", func(t *testing.T) {
		for _, prod := range products {
			var (
				data    menu.GetProductByKeyResponse
				payload = handler.Response{
					Data: &data,
				}
			)

			response := app.Request(t, http.MethodGet, fmt.Sprintf("/%s", prod.ID)).To(&payload).Run()
			if response.data.StatusCode != http.StatusOK {
				t.Errorf("Failed to fetch product %s: %v", prod.ID, payload.Error)
				return
			}

			assert.Equal(t, prod, data, "the product received was not as the expected for %s", prod)
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
						Name: "Coke 1 Liter",
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
