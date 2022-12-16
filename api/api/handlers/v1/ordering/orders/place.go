package orders

import (
	"comies/api/request"
	"comies/api/send"
	"comies/core/types"
	"comies/jobs/ordering"
	"context"
	"net/http"
	"time"
)

type Ticket struct {
	CustomerName    string `json:"customer_name"`
	CustomerPhone   string `json:"customer_phone"`
	CustomerAddress string `json:"customer_address"`

	DeliveryType types.Type `json:"delivery_type"`
	Observations string     `json:"observations"`

	Items []TicketItem `json:"items"`
}

type TicketItem struct {
	ProductID    types.ID       `json:"product_id"`
	Quantity     types.Quantity `json:"quantity"`
	Observations string         `json:"observations"`
}

func (h Handler) Place(ctx context.Context, r request.Request) send.Response {
	var order Ticket
	err := r.JSONBody(&order)
	if err != nil {
		return send.JSONError(err)
	}

	items := make([]ordering.TicketItem, len(order.Items))
	for i, it := range order.Items {
		items[i] = ordering.TicketItem{
			ProductID:    it.ProductID,
			Quantity:     it.Quantity,
			Observations: it.Observations,
		}
	}

	summ, err := h.orders.PlaceOrder(ctx, ordering.Ticket{
		CustomerName:    order.CustomerName,
		CustomerPhone:   order.CustomerPhone,
		CustomerAddress: order.CustomerAddress,
		DeliveryType:    order.DeliveryType,
		Observations:    order.Observations,
		Date:            time.Now(),
		Items:           items,
	})

	if err != nil {
		return send.FromError(err)
	}

	r.Commit(ctx)

	return send.Data(http.StatusCreated, summ, send.WithHeaders(map[string]string{"Location": summ.ID.String()}))
}
