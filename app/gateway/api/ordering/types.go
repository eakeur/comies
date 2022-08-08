package ordering

import (
	"comies/app/core/entities/item"
	"comies/app/core/entities/order"
	"comies/app/core/types"
	"time"
)

type (
	ItemAdditionRequest struct {
		ProductID          types.ID          `json:"product_id"`
		Quantity           types.Quantity    `json:"quantity"`
		Observations       string            `json:"observations"`
		IgnoreIngredients  []string          `json:"ignore_ingredients"`
		ReplaceIngredients map[string]string `json:"replace_ingredients"`
	}

	ConfirmOrderRequest struct {
		DeliveryMode order.DeliveryMode `json:"delivery_mode"`
	}

	Order struct {
		ID             string             `json:"id"`
		Identification string             `json:"identification"`
		PlacedAt       time.Time          `json:"placed_at"`
		Status         order.Status       `json:"status"`
		DeliveryMode   order.DeliveryMode `json:"delivery_mode"`
		Observations   string             `json:"observations"`
		FinalPrice     types.Currency     `json:"final_price"`
		Address        string             `json:"address"`
		Phone          string             `json:"phone"`
		Items          []Item             `json:"items,omitempty"`
	}

	Item struct {
		ID                 string            `json:"id,omitempty"`
		OrderID            string            `json:"order_id"`
		Status             item.Status       `json:"status"`
		Price              types.Currency    `json:"price"`
		ProductID          types.ID          `json:"product_id"`
		Quantity           types.Quantity    `json:"quantity"`
		Observations       string            `json:"observations"`
		IgnoreIngredients  []string          `json:"ignore_ingredients"`
		ReplaceIngredients map[string]string `json:"replace_ingredients"`
	}

	Failure struct {
		For       string `json:"for"`
		ProductID string `json:"product_id"`
		Error     error  `json:"error"`
	}

	OrderConfirmation struct {
		DeliveryMode order.DeliveryMode `json:"delivery_mode"`
	}

	SetItemStatusRequest struct {
		Status item.Status `json:"status"`
	}

	SetOrderStatusRequest struct {
		Status order.Status `json:"status"`
	}

	SetOrderDeliveryModeRequest struct {
		Mode order.DeliveryMode `json:"mode"`
	}

	OrderRequestResponse struct {
		ID string `json:"id"`
	}

	ItemAdditionResponse struct {
		ID string `json:"id"`
	}
)
