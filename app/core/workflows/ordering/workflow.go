package ordering

import (
	"context"
	"gomies/app/core/entities/order"
	"gomies/app/sdk/fault"
	"gomies/app/sdk/types"
)

var _ Workflow = workflow{}

type (
	Workflow interface {
		RequestOrderTicket(ctx context.Context, customerID types.ID) (types.ID, error)
		Order(ctx context.Context, o OrderConfirmation) (order.Order, error)
		AddToOrder(ctx context.Context, i order.Item) (ItemAdditionResult, error)

		UpdateOrderDeliveryMode(ctx context.Context, id types.ID, deliveryMode order.DeliveryMode) error
		UpdateOrderStatus(ctx context.Context, id types.ID, st order.Status) error
		UpdateOrderAddressID(ctx context.Context, id types.ID, addressID types.ID) error

		UpdateItemStatus(ctx context.Context, id types.ID, status order.PreparationStatus) error

		UpdateContentStatus(ctx context.Context, id types.ID, status order.PreparationStatus) error
		UpdateContentQuantity(ctx context.Context, id types.ID, qt types.Quantity) error

		ListOrders(ctx context.Context, f order.Filter) ([]order.Order, int, error)
		GetOrder(ctx context.Context, id types.ID) (order.Order, error)
		CancelOrder(ctx context.Context, id types.ID) error
	}

	workflow struct {
		products ProductService
		orders   order.Actions
	}
)

func (w workflow) RequestOrderTicket(ctx context.Context, customerID types.ID) (types.ID, error) {
	const operation = "Workflow.Ordering.RequestOrderTicket"

	if customerID.Empty() {
		return 0, fault.Wrap(fault.ErrMissingUID, operation)
	}

	o, err := w.orders.CreateOrder(ctx, order.Order{
		CustomerID: customerID,
		Status:     order.InTheCartStatus,
	})
	if err != nil {
		return 0, fault.Wrap(err, operation)
	}

	return o.ID, nil
}

func (w workflow) AddToOrder(ctx context.Context, i order.Item) (ItemAdditionResult, error) {
	const operation = "Workflow.Ordering.AddToOrder"

	if i.OrderID.Empty() {
		return ItemAdditionResult{}, fault.Wrap(fault.ErrMissingUID, operation)
	}

	i, err := w.orders.CreateItem(ctx, i)
	if err != nil {
		return ItemAdditionResult{}, fault.Wrap(err, operation)
	}

	errChan := make(chan error, len(i.Products))
	results := make(chan error, len(i.Products))

	for _, product := range i.Products {
		product := product
	}

	return o.ID, nil
}

func (w workflow) Order(ctx context.Context, o OrderConfirmation) (order.Order, error) {
	//TODO implement me
	panic("implement me")
}

func (w workflow) UpdateOrderDeliveryMode(ctx context.Context, id types.ID, deliveryMode order.DeliveryMode) error {
	//TODO implement me
	panic("implement me")
}

func (w workflow) UpdateOrderStatus(ctx context.Context, id types.ID, st order.Status) error {
	//TODO implement me
	panic("implement me")
}

func (w workflow) UpdateOrderAddressID(ctx context.Context, id types.ID, addressID types.ID) error {
	//TODO implement me
	panic("implement me")
}

func (w workflow) UpdateItemStatus(ctx context.Context, id types.ID, status order.PreparationStatus) error {
	//TODO implement me
	panic("implement me")
}

func (w workflow) UpdateContentStatus(ctx context.Context, id types.ID, status order.PreparationStatus) error {
	//TODO implement me
	panic("implement me")
}

func (w workflow) UpdateContentQuantity(ctx context.Context, id types.ID, qt types.Quantity) error {
	//TODO implement me
	panic("implement me")
}

func (w workflow) ListOrders(ctx context.Context, f order.Filter) ([]order.Order, int, error) {
	//TODO implement me
	panic("implement me")
}

func (w workflow) GetOrder(ctx context.Context, id types.ID) (order.Order, error) {
	//TODO implement me
	panic("implement me")
}

func (w workflow) CancelOrder(ctx context.Context, id types.ID) error {
	//TODO implement me
	panic("implement me")
}

func NewWorkflow(orders order.Actions, products ProductService) Workflow {
	return workflow{
		products: products,
		orders:   orders,
	}
}
