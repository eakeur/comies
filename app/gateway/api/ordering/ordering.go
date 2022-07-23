package ordering

import (
	"comies/app/core/entities/item"
	"comies/app/core/entities/order"
	"comies/app/core/entities/product"
	"comies/app/core/workflows/ordering"
	"comies/app/gateway/api/handler"
	"comies/app/gateway/api/response"
	"net/http"
)

var failures = response.ErrorBinding{
	product.ErrStockNegative: response.WithError(http.StatusPreconditionFailed, response.Error{
		Code:    "PRODUCT_STOCK_EMPTY",
		Message: "Ops! This product has not enough available stock",
	}),
	product.ErrStockAlreadyFull: response.WithError(http.StatusPreconditionFailed, response.Error{
		Code:    "PRODUCT_STOCK_FULL",
		Message: "Ops! This product's stock is already full",
	}),
	item.ErrInvalidQuantity: response.WithError(http.StatusUnprocessableEntity, response.Error{
		Code:    "ITEM_INVALID_QUANTITY",
		Message: "Ops! This item's quantity should be greater than 0 to be ordered",
	}),
	order.ErrAlreadyOrdered: response.WithError(http.StatusPreconditionFailed, response.Error{
		Code:    "ORDER_ALREADY_ORDERED",
		Message: "Ops! This order is already in process and can not be re-ordered",
	}),
	order.ErrAlreadyPreparing: response.WithError(http.StatusPreconditionFailed, response.Error{
		Code:    "ORDER_ALREADY_PREPARING",
		Message: "Ops! This order is already in process and can not be canceled",
	}),
	order.ErrInvalidNumberOfItems: response.WithError(http.StatusPreconditionFailed, response.Error{
		Code:    "ORDER_MUST_HAVE_ITEMS",
		Message: "Ops! This order has no items to be ordered",
	}),
}

type Service struct {
	RequestOrderTicketRoute handler.Route `method:"POST" path:"/orders" middleware:"tx"`
	OrderRoute              handler.Route `method:"POST" path:"/orders/{order_id}/order" middleware:"tx"`
	AddToOrderRoute         handler.Route `method:"POST" path:"/orders/{order_id}/items" body:"Item" middleware:"tx"`

	CancelOrderRoute handler.Route `method:"DELETE" path:"/orders/{order_id}" url:"order_id" middleware:"tx"`

	ListOrdersInFlowRoute handler.Route `method:"GET" path:"/ofl"`
	ListOrdersRoute       handler.Route `method:"GET" path:"/orders"`
	GetOrderByIDRoute     handler.Route `method:"GET" path:"/orders/{order_id}" url:"order_id"`
	ListItemsRoute        handler.Route `method:"GET" path:"/orders/{order_id}/items" url:"order_id"`

	SetOrderStatusRoute       handler.Route `method:"PUT" path:"/orders/{order_id}/status" url:"order_id" middleware:"tx"`
	SetOrderDeliveryModeRoute handler.Route `method:"PUT" path:"/orders/{order_id}/delivery-mode" url:"order_id" middleware:"tx"`
	SetItemStatusRoute        handler.Route `method:"PUT" path:"/orders/{order_id}/items/{item_id}" url:"item_id" middleware:"tx"`

	ordering ordering.Workflow
}

func NewService(ordering ordering.Workflow) *Service {
	return &Service{
		ordering: ordering,
	}
}
