package ordering

import (
	"comies/app/core/id"
	"comies/app/core/order"
	"comies/app/core/types"
	"comies/app/handler/rest"
	"comies/app/workflows/ordering"
	"context"
	"net/http"
	"time"
)

type GetOrderByIDResponse struct {
	ID             id.ID              `json:"id"`
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

// GetOrderByID fetches a specific order
//
// @Summary     Get order by ID
// @Description Fetches an order looking for its ID
// @Tags        Ordering
// @Param       order_id path     string                  false "The order ID"
// @Success     200         {object} rest.Response{data=PlaceOrder{}}
// @Failure     400         {object} rest.Response{error=rest.Error{}} "INVALID_ID"
// @Failure     404         {object} rest.Response{data=[]Failure{}} "ORDER_NOT_FOUND"
// @Failure     500         {object} rest.Response{error=rest.Error{}} "ERR_INTERNAL_SERVER_ERROR"
// @Router      /ordering/orders/{order_id}/items [POST]
func GetOrderByID(ctx context.Context, r *http.Request) rest.Response {
	id, err := rest.GetResourceIDFromURL(r, "order_id")
	if err != nil {
		return rest.IDParsingErrorResponse(err)
	}

	o, err := ordering.GetOrderByID(ctx, id)
	if err != nil {
		return rest.Fail(err)
	}

	return rest.ResponseWithData(http.StatusOK, NewGetOrderByIDResponse(o))
}

func NewGetOrderByIDResponse(o order.Order) GetOrderByIDResponse {
	return GetOrderByIDResponse{
		ID:             o.ID,
		Identification: o.Identification,
		PlacedAt:       o.PlacedAt,
		Status:         o.Status,
		DeliveryMode:   o.DeliveryMode,
		Observations:   o.Observations,
		FinalPrice:     o.FinalPrice,
		Address:        o.Address,
		Phone:          o.Phone,
	}
}
