package ordering

import (
	"comies/app/core/workflows/ordering"
	"comies/app/gateway/api/handler"
)

type Service struct {
	RequestOrderTicketRoute handler.Route `method:"POST" path:"/ordering/request" middleware:"tx"`
	OrderRoute              handler.Route `method:"POST" path:"/ordering/orders/{order_id}/confirm" middleware:"tx"`
	AddToOrderRoute         handler.Route `method:"POST" path:"/ordering/orders/{order_id}/items" body:"Item" middleware:"tx"`

	CancelOrderRoute handler.Route `method:"DELETE" path:"/ordering/orders/{order_id}" url:"order_id" middleware:"tx"`

	ListOrdersInFlowRoute handler.Route `method:"GET" path:"/ordering/flow"`
	ListOrdersRoute       handler.Route `method:"GET" path:"/ordering/orders"`
	GetOrderByIDRoute     handler.Route `method:"GET" path:"/ordering/orders/{order_id}" url:"order_id"`
	ListItemsRoute        handler.Route `method:"GET" path:"/ordering/orders/{order_id}/items" url:"order_id"`

	SetOrderStatusRoute       handler.Route `method:"PUT" path:"/ordering/orders/{order_id}/status" url:"order_id" middleware:"tx"`
	SetOrderDeliveryModeRoute handler.Route `method:"PUT" path:"/ordering/orders/{order_id}/delivery-mode" url:"order_id" middleware:"tx"`
	SetItemStatusRoute        handler.Route `method:"PUT" path:"/ordering/orders/{order_id}/items/{item_id}" url:"item_id" middleware:"tx"`

	ordering ordering.Workflow
}

func NewService(ordering ordering.Workflow) *Service {
	return &Service{
		ordering: ordering,
	}
}
