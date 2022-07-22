package ordering

import (
	"comies/app/core/entities/item"
	"comies/app/core/entities/order"
	"comies/app/gateway/api/response"
	"comies/app/sdk/throw"
	"comies/app/sdk/types"
	"net/http"
	"strconv"
	"time"
)

type (
	Order struct {
		ID             types.ID           `json:"id"`
		Identification string             `json:"identification,omitempty"`
		PlacedAt       time.Time          `json:"placed_at"`
		Status         order.Status       `json:"status,omitempty"`
		DeliveryMode   order.DeliveryMode `json:"delivery_mode,omitempty"`
		Observations   string             `json:"observations,omitempty"`
		FinalPrice     types.Currency     `json:"final_price,omitempty"`
		Address        string             `json:"address,omitempty"`
		Phone          string             `json:"phone,omitempty"`
		Items          []Item             `json:"items,omitempty"`
	}

	Item struct {
		ID                 types.ID              `json:"id,omitempty"`
		OrderID            types.ID              `json:"order_id"`
		Status             item.Status           `json:"status"`
		Price              types.Currency        `json:"price"`
		ProductID          types.ID              `json:"product_id"`
		Quantity           types.Quantity        `json:"quantity"`
		Observations       string                `json:"observations"`
		IgnoreIngredients  []types.ID            `json:"ignore_ingredients"`
		ReplaceIngredients map[types.ID]types.ID `json:"replace_ingredients"`
	}

	Failure struct {
		For       types.ID `json:"for"`
		ProductID types.ID `json:"product_id"`
		Error     error    `json:"error"`
	}

	OrderConfirmation struct {
		OrderID      types.ID           `json:"order_id"`
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

	AdditionResult struct {
		ID types.ID `json:"id"`
	}
)

func convertToID(in string) (types.ID, error, response.Response) {
	id, err := strconv.Atoi(in)
	if err != nil {
		return 0, err, response.WithError(http.StatusBadRequest, response.Error{
			Code: "INVALID_ID", Message: "The id provided is invalid",
		}).Err(throw.Error(err))
	}

	return types.ID(id), nil, response.Response{}
}
