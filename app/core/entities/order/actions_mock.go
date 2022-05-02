// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package order

import (
	"context"
	"gomies/app/sdk/types"
	"sync"
)

// Ensure, that ActionsMock does implement Actions.
// If this is not the case, regenerate this file with moq.
var _ Actions = &ActionsMock{}

// ActionsMock is a mock implementation of Actions.
//
// 	func TestSomethingThatUsesActions(t *testing.T) {
//
// 		// make and configure a mocked Actions
// 		mockedActions := &ActionsMock{
// 			CreateContentFunc: func(ctx context.Context, c ...Content) ([]Content, error) {
// 				panic("mock out the CreateContent method")
// 			},
// 			CreateItemFunc: func(ctx context.Context, item Item) (Item, error) {
// 				panic("mock out the CreateItem method")
// 			},
// 			CreateOrderFunc: func(ctx context.Context, o Order) (Order, error) {
// 				panic("mock out the CreateOrder method")
// 			},
// 			GetOrderFunc: func(ctx context.Context, id types.ID) (Order, error) {
// 				panic("mock out the GetOrder method")
// 			},
// 			ListContentFunc: func(ctx context.Context, itemUID types.ID) ([]Content, error) {
// 				panic("mock out the ListContent method")
// 			},
// 			ListItemsFunc: func(ctx context.Context, orderUID types.ID) ([]Item, error) {
// 				panic("mock out the ListItems method")
// 			},
// 			ListOrdersFunc: func(ctx context.Context, f Filter) ([]Order, int, error) {
// 				panic("mock out the ListOrders method")
// 			},
// 			RemoveContentFunc: func(ctx context.Context, id types.ID) error {
// 				panic("mock out the RemoveContent method")
// 			},
// 			RemoveItemFunc: func(ctx context.Context, id types.ID) error {
// 				panic("mock out the RemoveItem method")
// 			},
// 			RemoveOrderFunc: func(ctx context.Context, o Order) error {
// 				panic("mock out the RemoveOrder method")
// 			},
// 			UpdateContentQuantityFunc: func(ctx context.Context, id types.ID, qt types.Quantity) error {
// 				panic("mock out the UpdateContentQuantity method")
// 			},
// 			UpdateContentStatusFunc: func(ctx context.Context, id types.ID, status PreparationStatus) error {
// 				panic("mock out the UpdateContentStatus method")
// 			},
// 			UpdateItemStatusFunc: func(ctx context.Context, id types.ID, status PreparationStatus) error {
// 				panic("mock out the UpdateItemStatus method")
// 			},
// 			UpdateOrderAddressIDFunc: func(ctx context.Context, id types.ID, addressID types.ID) error {
// 				panic("mock out the UpdateOrderAddressID method")
// 			},
// 			UpdateOrderDeliveryModeFunc: func(ctx context.Context, id types.ID, deliverType DeliveryMode) error {
// 				panic("mock out the UpdateOrderDeliveryMode method")
// 			},
// 			UpdateOrderStatusFunc: func(ctx context.Context, id types.ID, status Status) error {
// 				panic("mock out the UpdateOrderStatus method")
// 			},
// 		}
//
// 		// use mockedActions in code that requires Actions
// 		// and then make assertions.
//
// 	}
type ActionsMock struct {
	// CreateContentFunc mocks the CreateContent method.
	CreateContentFunc func(ctx context.Context, c ...Content) ([]Content, error)

	// CreateItemFunc mocks the CreateItem method.
	CreateItemFunc func(ctx context.Context, item Item) (Item, error)

	// CreateOrderFunc mocks the CreateOrder method.
	CreateOrderFunc func(ctx context.Context, o Order) (Order, error)

	// GetOrderFunc mocks the GetOrder method.
	GetOrderFunc func(ctx context.Context, id types.ID) (Order, error)

	// ListContentFunc mocks the ListContent method.
	ListContentFunc func(ctx context.Context, itemUID types.ID) ([]Content, error)

	// ListItemsFunc mocks the ListItems method.
	ListItemsFunc func(ctx context.Context, orderUID types.ID) ([]Item, error)

	// ListOrdersFunc mocks the ListOrders method.
	ListOrdersFunc func(ctx context.Context, f Filter) ([]Order, int, error)

	// RemoveContentFunc mocks the RemoveContent method.
	RemoveContentFunc func(ctx context.Context, id types.ID) error

	// RemoveItemFunc mocks the RemoveItem method.
	RemoveItemFunc func(ctx context.Context, id types.ID) error

	// RemoveOrderFunc mocks the RemoveOrder method.
	RemoveOrderFunc func(ctx context.Context, o Order) error

	// UpdateContentQuantityFunc mocks the UpdateContentQuantity method.
	UpdateContentQuantityFunc func(ctx context.Context, id types.ID, qt types.Quantity) error

	// UpdateContentStatusFunc mocks the UpdateContentStatus method.
	UpdateContentStatusFunc func(ctx context.Context, id types.ID, status PreparationStatus) error

	// UpdateItemStatusFunc mocks the UpdateItemStatus method.
	UpdateItemStatusFunc func(ctx context.Context, id types.ID, status PreparationStatus) error

	// UpdateOrderAddressIDFunc mocks the UpdateOrderAddressID method.
	UpdateOrderAddressIDFunc func(ctx context.Context, id types.ID, addressID types.ID) error

	// UpdateOrderDeliveryModeFunc mocks the UpdateOrderDeliveryMode method.
	UpdateOrderDeliveryModeFunc func(ctx context.Context, id types.ID, deliverType DeliveryMode) error

	// UpdateOrderStatusFunc mocks the UpdateOrderStatus method.
	UpdateOrderStatusFunc func(ctx context.Context, id types.ID, status Status) error

	// calls tracks calls to the methods.
	calls struct {
		// CreateContent holds details about calls to the CreateContent method.
		CreateContent []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// C is the c argument value.
			C []Content
		}
		// CreateItem holds details about calls to the CreateItem method.
		CreateItem []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Item is the item argument value.
			Item Item
		}
		// CreateOrder holds details about calls to the CreateOrder method.
		CreateOrder []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// O is the o argument value.
			O Order
		}
		// GetOrder holds details about calls to the GetOrder method.
		GetOrder []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID types.ID
		}
		// ListContent holds details about calls to the ListContent method.
		ListContent []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ItemUID is the itemUID argument value.
			ItemUID types.ID
		}
		// ListItems holds details about calls to the ListItems method.
		ListItems []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// OrderUID is the orderUID argument value.
			OrderUID types.ID
		}
		// ListOrders holds details about calls to the ListOrders method.
		ListOrders []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// F is the f argument value.
			F Filter
		}
		// RemoveContent holds details about calls to the RemoveContent method.
		RemoveContent []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID types.ID
		}
		// RemoveItem holds details about calls to the RemoveItem method.
		RemoveItem []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID types.ID
		}
		// RemoveOrder holds details about calls to the RemoveOrder method.
		RemoveOrder []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// O is the o argument value.
			O Order
		}
		// UpdateContentQuantity holds details about calls to the UpdateContentQuantity method.
		UpdateContentQuantity []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID types.ID
			// Qt is the qt argument value.
			Qt types.Quantity
		}
		// UpdateContentStatus holds details about calls to the UpdateContentStatus method.
		UpdateContentStatus []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID types.ID
			// Status is the status argument value.
			Status PreparationStatus
		}
		// UpdateItemStatus holds details about calls to the UpdateItemStatus method.
		UpdateItemStatus []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID types.ID
			// Status is the status argument value.
			Status PreparationStatus
		}
		// UpdateOrderAddressID holds details about calls to the UpdateOrderAddressID method.
		UpdateOrderAddressID []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID types.ID
			// AddressID is the addressID argument value.
			AddressID types.ID
		}
		// UpdateOrderDeliveryMode holds details about calls to the UpdateOrderDeliveryMode method.
		UpdateOrderDeliveryMode []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID types.ID
			// DeliverType is the deliverType argument value.
			DeliverType DeliveryMode
		}
		// UpdateOrderStatus holds details about calls to the UpdateOrderStatus method.
		UpdateOrderStatus []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID types.ID
			// Status is the status argument value.
			Status Status
		}
	}
	lockCreateContent           sync.RWMutex
	lockCreateItem              sync.RWMutex
	lockCreateOrder             sync.RWMutex
	lockGetOrder                sync.RWMutex
	lockListContent             sync.RWMutex
	lockListItems               sync.RWMutex
	lockListOrders              sync.RWMutex
	lockRemoveContent           sync.RWMutex
	lockRemoveItem              sync.RWMutex
	lockRemoveOrder             sync.RWMutex
	lockUpdateContentQuantity   sync.RWMutex
	lockUpdateContentStatus     sync.RWMutex
	lockUpdateItemStatus        sync.RWMutex
	lockUpdateOrderAddressID    sync.RWMutex
	lockUpdateOrderDeliveryMode sync.RWMutex
	lockUpdateOrderStatus       sync.RWMutex
}

// CreateContent calls CreateContentFunc.
func (mock *ActionsMock) CreateContent(ctx context.Context, c ...Content) ([]Content, error) {
	if mock.CreateContentFunc == nil {
		panic("ActionsMock.CreateContentFunc: method is nil but Actions.CreateContent was just called")
	}
	callInfo := struct {
		Ctx context.Context
		C   []Content
	}{
		Ctx: ctx,
		C:   c,
	}
	mock.lockCreateContent.Lock()
	mock.calls.CreateContent = append(mock.calls.CreateContent, callInfo)
	mock.lockCreateContent.Unlock()
	return mock.CreateContentFunc(ctx, c...)
}

// CreateContentCalls gets all the calls that were made to CreateContent.
// Check the length with:
//     len(mockedActions.CreateContentCalls())
func (mock *ActionsMock) CreateContentCalls() []struct {
	Ctx context.Context
	C   []Content
} {
	var calls []struct {
		Ctx context.Context
		C   []Content
	}
	mock.lockCreateContent.RLock()
	calls = mock.calls.CreateContent
	mock.lockCreateContent.RUnlock()
	return calls
}

// CreateItem calls CreateItemFunc.
func (mock *ActionsMock) CreateItem(ctx context.Context, item Item) (Item, error) {
	if mock.CreateItemFunc == nil {
		panic("ActionsMock.CreateItemFunc: method is nil but Actions.CreateItem was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		Item Item
	}{
		Ctx:  ctx,
		Item: item,
	}
	mock.lockCreateItem.Lock()
	mock.calls.CreateItem = append(mock.calls.CreateItem, callInfo)
	mock.lockCreateItem.Unlock()
	return mock.CreateItemFunc(ctx, item)
}

// CreateItemCalls gets all the calls that were made to CreateItem.
// Check the length with:
//     len(mockedActions.CreateItemCalls())
func (mock *ActionsMock) CreateItemCalls() []struct {
	Ctx  context.Context
	Item Item
} {
	var calls []struct {
		Ctx  context.Context
		Item Item
	}
	mock.lockCreateItem.RLock()
	calls = mock.calls.CreateItem
	mock.lockCreateItem.RUnlock()
	return calls
}

// CreateOrder calls CreateOrderFunc.
func (mock *ActionsMock) CreateOrder(ctx context.Context, o Order) (Order, error) {
	if mock.CreateOrderFunc == nil {
		panic("ActionsMock.CreateOrderFunc: method is nil but Actions.CreateOrder was just called")
	}
	callInfo := struct {
		Ctx context.Context
		O   Order
	}{
		Ctx: ctx,
		O:   o,
	}
	mock.lockCreateOrder.Lock()
	mock.calls.CreateOrder = append(mock.calls.CreateOrder, callInfo)
	mock.lockCreateOrder.Unlock()
	return mock.CreateOrderFunc(ctx, o)
}

// CreateOrderCalls gets all the calls that were made to CreateOrder.
// Check the length with:
//     len(mockedActions.CreateOrderCalls())
func (mock *ActionsMock) CreateOrderCalls() []struct {
	Ctx context.Context
	O   Order
} {
	var calls []struct {
		Ctx context.Context
		O   Order
	}
	mock.lockCreateOrder.RLock()
	calls = mock.calls.CreateOrder
	mock.lockCreateOrder.RUnlock()
	return calls
}

// GetOrder calls GetOrderFunc.
func (mock *ActionsMock) GetOrder(ctx context.Context, id types.ID) (Order, error) {
	if mock.GetOrderFunc == nil {
		panic("ActionsMock.GetOrderFunc: method is nil but Actions.GetOrder was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  types.ID
	}{
		Ctx: ctx,
		ID:  id,
	}
	mock.lockGetOrder.Lock()
	mock.calls.GetOrder = append(mock.calls.GetOrder, callInfo)
	mock.lockGetOrder.Unlock()
	return mock.GetOrderFunc(ctx, id)
}

// GetOrderCalls gets all the calls that were made to GetOrder.
// Check the length with:
//     len(mockedActions.GetOrderCalls())
func (mock *ActionsMock) GetOrderCalls() []struct {
	Ctx context.Context
	ID  types.ID
} {
	var calls []struct {
		Ctx context.Context
		ID  types.ID
	}
	mock.lockGetOrder.RLock()
	calls = mock.calls.GetOrder
	mock.lockGetOrder.RUnlock()
	return calls
}

// ListContent calls ListContentFunc.
func (mock *ActionsMock) ListContent(ctx context.Context, itemUID types.ID) ([]Content, error) {
	if mock.ListContentFunc == nil {
		panic("ActionsMock.ListContentFunc: method is nil but Actions.ListContent was just called")
	}
	callInfo := struct {
		Ctx     context.Context
		ItemUID types.ID
	}{
		Ctx:     ctx,
		ItemUID: itemUID,
	}
	mock.lockListContent.Lock()
	mock.calls.ListContent = append(mock.calls.ListContent, callInfo)
	mock.lockListContent.Unlock()
	return mock.ListContentFunc(ctx, itemUID)
}

// ListContentCalls gets all the calls that were made to ListContent.
// Check the length with:
//     len(mockedActions.ListContentCalls())
func (mock *ActionsMock) ListContentCalls() []struct {
	Ctx     context.Context
	ItemUID types.ID
} {
	var calls []struct {
		Ctx     context.Context
		ItemUID types.ID
	}
	mock.lockListContent.RLock()
	calls = mock.calls.ListContent
	mock.lockListContent.RUnlock()
	return calls
}

// ListItems calls ListItemsFunc.
func (mock *ActionsMock) ListItems(ctx context.Context, orderUID types.ID) ([]Item, error) {
	if mock.ListItemsFunc == nil {
		panic("ActionsMock.ListItemsFunc: method is nil but Actions.ListItems was just called")
	}
	callInfo := struct {
		Ctx      context.Context
		OrderUID types.ID
	}{
		Ctx:      ctx,
		OrderUID: orderUID,
	}
	mock.lockListItems.Lock()
	mock.calls.ListItems = append(mock.calls.ListItems, callInfo)
	mock.lockListItems.Unlock()
	return mock.ListItemsFunc(ctx, orderUID)
}

// ListItemsCalls gets all the calls that were made to ListItems.
// Check the length with:
//     len(mockedActions.ListItemsCalls())
func (mock *ActionsMock) ListItemsCalls() []struct {
	Ctx      context.Context
	OrderUID types.ID
} {
	var calls []struct {
		Ctx      context.Context
		OrderUID types.ID
	}
	mock.lockListItems.RLock()
	calls = mock.calls.ListItems
	mock.lockListItems.RUnlock()
	return calls
}

// ListOrders calls ListOrdersFunc.
func (mock *ActionsMock) ListOrders(ctx context.Context, f Filter) ([]Order, int, error) {
	if mock.ListOrdersFunc == nil {
		panic("ActionsMock.ListOrdersFunc: method is nil but Actions.ListOrders was just called")
	}
	callInfo := struct {
		Ctx context.Context
		F   Filter
	}{
		Ctx: ctx,
		F:   f,
	}
	mock.lockListOrders.Lock()
	mock.calls.ListOrders = append(mock.calls.ListOrders, callInfo)
	mock.lockListOrders.Unlock()
	return mock.ListOrdersFunc(ctx, f)
}

// ListOrdersCalls gets all the calls that were made to ListOrders.
// Check the length with:
//     len(mockedActions.ListOrdersCalls())
func (mock *ActionsMock) ListOrdersCalls() []struct {
	Ctx context.Context
	F   Filter
} {
	var calls []struct {
		Ctx context.Context
		F   Filter
	}
	mock.lockListOrders.RLock()
	calls = mock.calls.ListOrders
	mock.lockListOrders.RUnlock()
	return calls
}

// RemoveContent calls RemoveContentFunc.
func (mock *ActionsMock) RemoveContent(ctx context.Context, id types.ID) error {
	if mock.RemoveContentFunc == nil {
		panic("ActionsMock.RemoveContentFunc: method is nil but Actions.RemoveContent was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  types.ID
	}{
		Ctx: ctx,
		ID:  id,
	}
	mock.lockRemoveContent.Lock()
	mock.calls.RemoveContent = append(mock.calls.RemoveContent, callInfo)
	mock.lockRemoveContent.Unlock()
	return mock.RemoveContentFunc(ctx, id)
}

// RemoveContentCalls gets all the calls that were made to RemoveContent.
// Check the length with:
//     len(mockedActions.RemoveContentCalls())
func (mock *ActionsMock) RemoveContentCalls() []struct {
	Ctx context.Context
	ID  types.ID
} {
	var calls []struct {
		Ctx context.Context
		ID  types.ID
	}
	mock.lockRemoveContent.RLock()
	calls = mock.calls.RemoveContent
	mock.lockRemoveContent.RUnlock()
	return calls
}

// RemoveItem calls RemoveItemFunc.
func (mock *ActionsMock) RemoveItem(ctx context.Context, id types.ID) error {
	if mock.RemoveItemFunc == nil {
		panic("ActionsMock.RemoveItemFunc: method is nil but Actions.RemoveItem was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  types.ID
	}{
		Ctx: ctx,
		ID:  id,
	}
	mock.lockRemoveItem.Lock()
	mock.calls.RemoveItem = append(mock.calls.RemoveItem, callInfo)
	mock.lockRemoveItem.Unlock()
	return mock.RemoveItemFunc(ctx, id)
}

// RemoveItemCalls gets all the calls that were made to RemoveItem.
// Check the length with:
//     len(mockedActions.RemoveItemCalls())
func (mock *ActionsMock) RemoveItemCalls() []struct {
	Ctx context.Context
	ID  types.ID
} {
	var calls []struct {
		Ctx context.Context
		ID  types.ID
	}
	mock.lockRemoveItem.RLock()
	calls = mock.calls.RemoveItem
	mock.lockRemoveItem.RUnlock()
	return calls
}

// RemoveOrder calls RemoveOrderFunc.
func (mock *ActionsMock) RemoveOrder(ctx context.Context, o Order) error {
	if mock.RemoveOrderFunc == nil {
		panic("ActionsMock.RemoveOrderFunc: method is nil but Actions.RemoveOrder was just called")
	}
	callInfo := struct {
		Ctx context.Context
		O   Order
	}{
		Ctx: ctx,
		O:   o,
	}
	mock.lockRemoveOrder.Lock()
	mock.calls.RemoveOrder = append(mock.calls.RemoveOrder, callInfo)
	mock.lockRemoveOrder.Unlock()
	return mock.RemoveOrderFunc(ctx, o)
}

// RemoveOrderCalls gets all the calls that were made to RemoveOrder.
// Check the length with:
//     len(mockedActions.RemoveOrderCalls())
func (mock *ActionsMock) RemoveOrderCalls() []struct {
	Ctx context.Context
	O   Order
} {
	var calls []struct {
		Ctx context.Context
		O   Order
	}
	mock.lockRemoveOrder.RLock()
	calls = mock.calls.RemoveOrder
	mock.lockRemoveOrder.RUnlock()
	return calls
}

// UpdateContentQuantity calls UpdateContentQuantityFunc.
func (mock *ActionsMock) UpdateContentQuantity(ctx context.Context, id types.ID, qt types.Quantity) error {
	if mock.UpdateContentQuantityFunc == nil {
		panic("ActionsMock.UpdateContentQuantityFunc: method is nil but Actions.UpdateContentQuantity was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  types.ID
		Qt  types.Quantity
	}{
		Ctx: ctx,
		ID:  id,
		Qt:  qt,
	}
	mock.lockUpdateContentQuantity.Lock()
	mock.calls.UpdateContentQuantity = append(mock.calls.UpdateContentQuantity, callInfo)
	mock.lockUpdateContentQuantity.Unlock()
	return mock.UpdateContentQuantityFunc(ctx, id, qt)
}

// UpdateContentQuantityCalls gets all the calls that were made to UpdateContentQuantity.
// Check the length with:
//     len(mockedActions.UpdateContentQuantityCalls())
func (mock *ActionsMock) UpdateContentQuantityCalls() []struct {
	Ctx context.Context
	ID  types.ID
	Qt  types.Quantity
} {
	var calls []struct {
		Ctx context.Context
		ID  types.ID
		Qt  types.Quantity
	}
	mock.lockUpdateContentQuantity.RLock()
	calls = mock.calls.UpdateContentQuantity
	mock.lockUpdateContentQuantity.RUnlock()
	return calls
}

// UpdateContentStatus calls UpdateContentStatusFunc.
func (mock *ActionsMock) UpdateContentStatus(ctx context.Context, id types.ID, status PreparationStatus) error {
	if mock.UpdateContentStatusFunc == nil {
		panic("ActionsMock.UpdateContentStatusFunc: method is nil but Actions.UpdateContentStatus was just called")
	}
	callInfo := struct {
		Ctx    context.Context
		ID     types.ID
		Status PreparationStatus
	}{
		Ctx:    ctx,
		ID:     id,
		Status: status,
	}
	mock.lockUpdateContentStatus.Lock()
	mock.calls.UpdateContentStatus = append(mock.calls.UpdateContentStatus, callInfo)
	mock.lockUpdateContentStatus.Unlock()
	return mock.UpdateContentStatusFunc(ctx, id, status)
}

// UpdateContentStatusCalls gets all the calls that were made to UpdateContentStatus.
// Check the length with:
//     len(mockedActions.UpdateContentStatusCalls())
func (mock *ActionsMock) UpdateContentStatusCalls() []struct {
	Ctx    context.Context
	ID     types.ID
	Status PreparationStatus
} {
	var calls []struct {
		Ctx    context.Context
		ID     types.ID
		Status PreparationStatus
	}
	mock.lockUpdateContentStatus.RLock()
	calls = mock.calls.UpdateContentStatus
	mock.lockUpdateContentStatus.RUnlock()
	return calls
}

// UpdateItemStatus calls UpdateItemStatusFunc.
func (mock *ActionsMock) UpdateItemStatus(ctx context.Context, id types.ID, status PreparationStatus) error {
	if mock.UpdateItemStatusFunc == nil {
		panic("ActionsMock.UpdateItemStatusFunc: method is nil but Actions.UpdateItemStatus was just called")
	}
	callInfo := struct {
		Ctx    context.Context
		ID     types.ID
		Status PreparationStatus
	}{
		Ctx:    ctx,
		ID:     id,
		Status: status,
	}
	mock.lockUpdateItemStatus.Lock()
	mock.calls.UpdateItemStatus = append(mock.calls.UpdateItemStatus, callInfo)
	mock.lockUpdateItemStatus.Unlock()
	return mock.UpdateItemStatusFunc(ctx, id, status)
}

// UpdateItemStatusCalls gets all the calls that were made to UpdateItemStatus.
// Check the length with:
//     len(mockedActions.UpdateItemStatusCalls())
func (mock *ActionsMock) UpdateItemStatusCalls() []struct {
	Ctx    context.Context
	ID     types.ID
	Status PreparationStatus
} {
	var calls []struct {
		Ctx    context.Context
		ID     types.ID
		Status PreparationStatus
	}
	mock.lockUpdateItemStatus.RLock()
	calls = mock.calls.UpdateItemStatus
	mock.lockUpdateItemStatus.RUnlock()
	return calls
}

// UpdateOrderAddressID calls UpdateOrderAddressIDFunc.
func (mock *ActionsMock) UpdateOrderAddressID(ctx context.Context, id types.ID, addressID types.ID) error {
	if mock.UpdateOrderAddressIDFunc == nil {
		panic("ActionsMock.UpdateOrderAddressIDFunc: method is nil but Actions.UpdateOrderAddressID was just called")
	}
	callInfo := struct {
		Ctx       context.Context
		ID        types.ID
		AddressID types.ID
	}{
		Ctx:       ctx,
		ID:        id,
		AddressID: addressID,
	}
	mock.lockUpdateOrderAddressID.Lock()
	mock.calls.UpdateOrderAddressID = append(mock.calls.UpdateOrderAddressID, callInfo)
	mock.lockUpdateOrderAddressID.Unlock()
	return mock.UpdateOrderAddressIDFunc(ctx, id, addressID)
}

// UpdateOrderAddressIDCalls gets all the calls that were made to UpdateOrderAddressID.
// Check the length with:
//     len(mockedActions.UpdateOrderAddressIDCalls())
func (mock *ActionsMock) UpdateOrderAddressIDCalls() []struct {
	Ctx       context.Context
	ID        types.ID
	AddressID types.ID
} {
	var calls []struct {
		Ctx       context.Context
		ID        types.ID
		AddressID types.ID
	}
	mock.lockUpdateOrderAddressID.RLock()
	calls = mock.calls.UpdateOrderAddressID
	mock.lockUpdateOrderAddressID.RUnlock()
	return calls
}

// UpdateOrderDeliveryMode calls UpdateOrderDeliveryModeFunc.
func (mock *ActionsMock) UpdateOrderDeliveryMode(ctx context.Context, id types.ID, deliverType DeliveryMode) error {
	if mock.UpdateOrderDeliveryModeFunc == nil {
		panic("ActionsMock.UpdateOrderDeliveryModeFunc: method is nil but Actions.UpdateOrderDeliveryMode was just called")
	}
	callInfo := struct {
		Ctx         context.Context
		ID          types.ID
		DeliverType DeliveryMode
	}{
		Ctx:         ctx,
		ID:          id,
		DeliverType: deliverType,
	}
	mock.lockUpdateOrderDeliveryMode.Lock()
	mock.calls.UpdateOrderDeliveryMode = append(mock.calls.UpdateOrderDeliveryMode, callInfo)
	mock.lockUpdateOrderDeliveryMode.Unlock()
	return mock.UpdateOrderDeliveryModeFunc(ctx, id, deliverType)
}

// UpdateOrderDeliveryModeCalls gets all the calls that were made to UpdateOrderDeliveryMode.
// Check the length with:
//     len(mockedActions.UpdateOrderDeliveryModeCalls())
func (mock *ActionsMock) UpdateOrderDeliveryModeCalls() []struct {
	Ctx         context.Context
	ID          types.ID
	DeliverType DeliveryMode
} {
	var calls []struct {
		Ctx         context.Context
		ID          types.ID
		DeliverType DeliveryMode
	}
	mock.lockUpdateOrderDeliveryMode.RLock()
	calls = mock.calls.UpdateOrderDeliveryMode
	mock.lockUpdateOrderDeliveryMode.RUnlock()
	return calls
}

// UpdateOrderStatus calls UpdateOrderStatusFunc.
func (mock *ActionsMock) UpdateOrderStatus(ctx context.Context, id types.ID, status Status) error {
	if mock.UpdateOrderStatusFunc == nil {
		panic("ActionsMock.UpdateOrderStatusFunc: method is nil but Actions.UpdateOrderStatus was just called")
	}
	callInfo := struct {
		Ctx    context.Context
		ID     types.ID
		Status Status
	}{
		Ctx:    ctx,
		ID:     id,
		Status: status,
	}
	mock.lockUpdateOrderStatus.Lock()
	mock.calls.UpdateOrderStatus = append(mock.calls.UpdateOrderStatus, callInfo)
	mock.lockUpdateOrderStatus.Unlock()
	return mock.UpdateOrderStatusFunc(ctx, id, status)
}

// UpdateOrderStatusCalls gets all the calls that were made to UpdateOrderStatus.
// Check the length with:
//     len(mockedActions.UpdateOrderStatusCalls())
func (mock *ActionsMock) UpdateOrderStatusCalls() []struct {
	Ctx    context.Context
	ID     types.ID
	Status Status
} {
	var calls []struct {
		Ctx    context.Context
		ID     types.ID
		Status Status
	}
	mock.lockUpdateOrderStatus.RLock()
	calls = mock.calls.UpdateOrderStatus
	mock.lockUpdateOrderStatus.RUnlock()
	return calls
}
