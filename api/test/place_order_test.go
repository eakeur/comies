package test

import (
	"bytes"
	"comies/api/handlers/menu"
	"comies/api/handlers/ordering/orders"
	"comies/core/types"
	"comies/jobs/ordering"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrderingAPI_PlaceOrder(t *testing.T) {
	t.Parallel()

	addr := createAPI(t)

	var productID types.ID
	t.Run("should create product", func(t *testing.T) {
		var route = fmt.Sprintf("%s/api/v1/menu/items", addr)

		prod, _ := json.Marshal(menu.Item{
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

		assert.Equal(t, http.StatusCreated, res.StatusCode)

		productID, err = types.FromString(res.Header.Get("Location"))
		if err != nil {
			t.Fatalf("could not retrieve product id from header: %v", err)
		}
	})

	t.Run("should create order", func(t *testing.T) {
		var route = fmt.Sprintf("%s/api/v1/ordering/orders", addr)

		ord, _ := json.Marshal(orders.Ticket{
			DeliveryType:    20,
			CustomerName:    "Ashumundum Vissam",
			CustomerPhone:   "991222212",
			CustomerAddress: "My Home, 2022",
			Items: []orders.TicketItem{
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

		assert.Equal(t, http.StatusCreated, res.StatusCode)

		data := ordering.OrderSummary{}
		err = json.NewDecoder(res.Body).Decode(&data)
		if err != nil {
			t.Fatalf("could not parse place order response: %s", err)
		}

		if data.BillAmountDue.Net != 15 {
			t.Errorf("order bill amount due is different from 15: got %d", data.BillAmountDue.Net)
		}
	})

	t.Run("should check if order is being prepared", func(t *testing.T) {
		const preparingStatus = 30.0

		res, err := http.Get(fmt.Sprintf("%s/api/v1/ordering/orders/991222212?phone=true", addr))
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, http.StatusOK, res.StatusCode)

		data := map[string]interface{}{}
		err = json.NewDecoder(res.Body).Decode(&data)
		if err != nil {
			t.Fatalf("could not parse order check response: %s", err)
		}

		assert.Equal(t, preparingStatus, data["Value"], "order status is not as the expected: %v", data)
	})

	t.Run("should check if product left stock", func(t *testing.T) {
		const expectedStock = -3.0

		res, err := http.Get(fmt.Sprintf("%s/api/v1/menu/items/%s/stock", addr, productID))
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, http.StatusOK, res.StatusCode)

		data := map[string]interface{}{}
		err = json.NewDecoder(res.Body).Decode(&data)
		if err != nil {
			t.Fatalf("could not parse balance response: %s", err)
		}

		assert.Equal(t, expectedStock, data["balance"], "product stock is not as the expected: %v", data)
	})

}
