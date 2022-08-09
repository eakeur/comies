package tests

import (
	"comies/app/core/entities/movement"
	"comies/app/core/entities/product"
	"comies/app/core/types"
	"comies/app/gateway/api/handler"
	"comies/app/gateway/api/menu"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
)

func TestMovementsIntegration(t *testing.T) {
	t.Parallel()

	var (
		productID string
		app       = NewTestApp(t, "/menu/products")
		movements = []menu.ListMovementsResponse{
			{
				Type:      movement.InputType,
				Date:      time.Now().UTC(),
				Quantity:  100,
				PaidValue: 800,
				AgentID:   "0",
			},
			{
				Type:      movement.InputType,
				Date:      time.Now().UTC(),
				Quantity:  50,
				PaidValue: 800,
				AgentID:   "0",
			},
			{
				Type:      movement.OutputType,
				Date:      time.Now().UTC(),
				Quantity:  50,
				PaidValue: 1000,
				AgentID:   "5221",
			},
			{
				Type:      movement.OutputType,
				Date:      time.Now().UTC(),
				Quantity:  50,
				PaidValue: 1000,
				AgentID:   "5222",
			},
			{
				Type:      movement.InputType,
				Date:      time.Now().UTC(),
				Quantity:  50,
				PaidValue: 800,
				AgentID:   "0",
			},
		}
	)

	t.Run("should create movement product successfully", func(t *testing.T) {
		var (
			data menu.CreateProductResponse

			payload = handler.Response{
				Data: &data,
			}
			input = RequestInput{
				body: menu.CreateProductRequest{
					Code:            "COKC",
					Name:            "Coke Can",
					Type:            product.OutputType,
					CostPrice:       800,
					SalePrice:       1000,
					SaleUnit:        "un",
					MinimumSale:     1,
					MaximumQuantity: 150,
					MinimumQuantity: 0,
					Location:        "Fridge",
				},
			}
		)

		response := app.Request(t, http.MethodPost, "", input).To(&payload).Run()
		if response.data.StatusCode != http.StatusCreated {
			t.Errorf("Failed to create movement product: %v", payload.Error)
			return
		}

		productID = data.ID
		for i := range movements {
			movements[i].ProductID = data.ID
		}
	})

	path := fmt.Sprintf("/%s/movements", productID)
	t.Run("should create movements", func(t *testing.T) {
		var (
			data menu.MovementAdditionResult

			payload = handler.Response{
				Data: &data,
			}
		)

		for i, m := range movements {
			response := app.Request(t, http.MethodPost, path, RequestInput{
				body: menu.CreateMovementRequest{
					Type:      m.Type,
					Date:      m.Date,
					Quantity:  m.Quantity,
					PaidValue: m.PaidValue,
					AgentID:   m.AgentID,
				},
			}).To(&payload).Run()
			if response.data.StatusCode != http.StatusCreated {
				t.Errorf("Failed to create movement: %v", payload.Error)
				return
			}

			movements[i].ID = data.ID
		}
	})

	t.Run("should list movements", func(t *testing.T) {
		var (
			data []menu.ListMovementsResponse

			payload = handler.Response{
				Data: &data,
			}
		)

		response := app.Request(t, http.MethodGet, path).To(&payload).Run()
		if response.data.StatusCode != http.StatusOK {
			t.Errorf("Failed to list movements: %v", payload.Error)
			return
		}

		for i := range movements {
			movements[i].Date = data[i].Date
		}

		assert.Equal(t, movements, data)
	})

	t.Run("should remove one movement", func(t *testing.T) {
		lastMov := movements[len(movements)-1]
		movements = movements[0 : len(movements)-1]
		response := app.Request(t, http.MethodDelete, fmt.Sprintf("%s/%s", path, lastMov.ID)).Run()
		if response.data.StatusCode != http.StatusNoContent {
			t.Errorf("Failed to delete movements: %v", response.dump)
			return
		}
	})

	t.Run("should fetch total balance", func(t *testing.T) {
		var (
			data menu.GetProductBalanceResponse

			payload = handler.Response{
				Data: &data,
			}
		)

		response := app.Request(t, http.MethodGet, fmt.Sprintf("/%s/stock-balance", productID)).To(&payload).Run()
		if response.data.StatusCode != http.StatusOK {
			t.Errorf("Failed to create movement product: %v", payload.Error)
			return
		}

		var sum types.Quantity
		for _, m := range movements {
			if m.Type == movement.InputType {
				sum += m.Quantity
			} else {
				sum -= m.Quantity
			}
		}

		assert.Equal(t, sum, data.Balance)
	})
}
