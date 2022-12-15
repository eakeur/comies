package test

import (
	"bytes"
	"comies/api/handlers/v1/menu/products"
	"comies/core/ordering/item"
	"comies/core/ordering/status"
	"comies/core/types"
	"comies/jobs/ordering"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

func TestOrderingAPI_PlaceOrder(t *testing.T) {
	t.Parallel()

	addr := createAPI(t)

	var productID types.ID
	t.Run("should create product", func(t *testing.T) {
		var route = fmt.Sprintf("%s/api/v1/menu/products", addr)

		prod, _ := json.Marshal(products.Product{
			Code:            "COCA",
			Name:            "Lata de Coca-Cola",
			Type:            10,
			SalePrice:       5,
			CostPrice:       3,
			MinimumSale:     1,
			SaleUnit:        "un",
			MaximumQuantity: 100,
			MinimumQuantity: 1,
		})

		res, err := http.Post(route, "application/json", bytes.NewReader(prod))
		if err != nil {
			t.Fatal(err)
		}

		if res.StatusCode != http.StatusCreated {
			t.Fatalf("could not create product: status(%v)", res.StatusCode)
		}

		productID, err = types.FromString(res.Header.Get("Location"))
		if err != nil {
			t.Fatalf("could not retrieve product id from header: %v", err)
		}
	})

	t.Run("should create order", func(t *testing.T) {
		var route = fmt.Sprintf("%s/api/v1/ordering/orders", addr)

		ord, _ := json.Marshal(ordering.Order{
			DeliveryType:    20,
			CustomerName:    "Ashumundum Vissam",
			CustomerPhone:   "991222212",
			CustomerAddress: "My Home, 2022",
			Items: []item.Item{
				{
					ProductID: productID,
					Quantity:  3,
				},
			},
		})

		res, err := http.Post(route, "application/json", bytes.NewReader(ord))
		if err != nil {
			t.Fatal(err)
		}

		if res.StatusCode != http.StatusCreated {
			t.Fatalf("could not create order: status(%v)", res.StatusCode)
		}
	})

	t.Run("should check if order is being prepared", func(t *testing.T) {
		var route = fmt.Sprintf("%s/api/v1/ordering/orders/991222212?phone=true", addr)

		res, err := http.Get(route)
		if err != nil {
			t.Fatal(err)
		}

		if res.StatusCode != http.StatusOK {
			t.Fatalf("could not create order: status(%v)", res.StatusCode)
		}

		data := ordering.Status{}
		err = json.NewDecoder(res.Body).Decode(&data)
		if err != nil {
			t.Fatalf("could not parse order check response: %s", err)
		}

		if data.Status.Value != status.PreparingStatus {
			t.Fatalf("order status is not as the expected: %v", data.Status)
		}
	})

}
