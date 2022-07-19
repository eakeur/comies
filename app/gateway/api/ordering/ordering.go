package ordering

import (
	"comies/app/core/entities/item"
	"comies/app/core/entities/order"
	"comies/app/core/workflows/ordering"
	"comies/app/gateway/api/errors"
	"comies/app/gateway/api/gen/ordering/protos"
	"comies/app/sdk/throw"
	"comies/app/sdk/types"
	"context"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var _ protos.OrderingServer = service{}

var failures = errors.ErrorBinding{}

type service struct {
	protos.UnimplementedOrderingServer
	ordering ordering.Workflow
}

func (s service) RequestOrderTicket(ctx context.Context, _ *protos.Empty) (*protos.RequestOrderTicketResponse, error) {
	ticket, err := s.ordering.RequestOrderTicket(ctx)
	if err != nil {
		return nil, failures.HandleError(throw.Error(err))
	}

	return &protos.RequestOrderTicketResponse{OrderId: int64(ticket)}, nil
}

func (s service) AddToOrder(ctx context.Context, request *protos.AddToOrderRequest) (*protos.AddToOrderResponse, error) {
	ignore := make([]types.ID, len(request.Item.Ignored))
	replacements := make(map[types.ID]types.ID, len(request.Item.Replacements))
	for i, id := range request.Item.Ignored {
		ignore[i] = types.ID(id)
	}
	for from, to := range request.Item.Replacements {
		replacements[types.ID(from)] = types.ID(to)
	}

	it, err := s.ordering.AddToOrder(ctx, item.Item{
		OrderID:      types.ID(request.Item.OrderId),
		ProductID:    types.ID(request.Item.ProductId),
		Quantity:     types.Quantity(request.Item.Quantity),
		Observations: request.Item.Observations,
		Details: item.Details{
			ReplaceIngredients: replacements,
			IgnoreIngredients:  ignore,
		},
	})
	if err != nil {
		return nil, failures.HandleError(throw.Error(err))
	}

	failures := make([]*protos.ReservationFailure, len(it.Failed))
	for _, f := range it.Failed {
		failures = append(failures, &protos.ReservationFailure{
			ProductId: int64(f.ProductID),
			Error:     f.Error.Error(),
		})
	}

	return &protos.AddToOrderResponse{Failures: failures}, nil
}

func (s service) Order(ctx context.Context, request *protos.OrderRequest) (*protos.OrderResponse, error) {
	o, err := s.ordering.Order(ctx, ordering.OrderConfirmation{
		OrderID:      types.ID(request.OrderId),
		DeliveryMode: order.DeliveryMode(request.DeliveryMode),
	})
	if err != nil {
		return nil, failures.HandleError(throw.Error(err))
	}

	return &protos.OrderResponse{
		Order: &protos.Order{
			Id:             int64(o.ID),
			Identification: o.Identification,
			PlacedAt:       timestamppb.New(o.PlacedAt),
			Status:         protos.OrderStatus(o.Status),
			DeliveryMode:   protos.DeliveryMode(o.DeliveryMode),
			Observation:    o.Observations,
			FinalPrice:     int64(o.FinalPrice),
			Address:        o.Address,
			Phone:          o.Phone,
		},
	}, nil
}

func (s service) SetOrderDeliveryMode(ctx context.Context, request *protos.SetOrderDeliveryModeRequest) (*protos.Empty, error) {
	err := s.ordering.SetOrderDeliveryMode(ctx, types.ID(request.OrderId), order.DeliveryMode(request.DeliveryMode))
	if err != nil {
		return nil, failures.HandleError(throw.Error(err))
	}

	return &protos.Empty{}, nil
}

func (s service) SetOrderStatus(ctx context.Context, request *protos.SetOrderStatusRequest) (*protos.Empty, error) {
	err := s.ordering.SetOrderStatus(ctx, types.ID(request.OrderId), order.Status(request.Status))
	if err != nil {
		return nil, failures.HandleError(throw.Error(err))
	}

	return &protos.Empty{}, nil
}

func (s service) ListOrders(ctx context.Context, request *protos.ListOrdersRequest) (*protos.ListOrdersResponse, error) {
	var statuses []order.Status
	if l := len(request.Filter.Statuses); l > 0 {
		statuses = make([]order.Status, l)
		for i, o := range request.Filter.Statuses {
			statuses[i] = order.Status(o)
		}
	}
	list, err := s.ordering.ListOrders(ctx, order.Filter{
		Status:       statuses,
		PlacedBefore: request.Filter.PlacedBefore.AsTime(),
		PlacedAfter:  request.Filter.PlacedAfter.AsTime(),
		DeliveryMode: order.DeliveryMode(request.Filter.DeliveryMode),
	})
	if err != nil {
		return nil, failures.HandleError(throw.Error(err))
	}

	orders := make([]*protos.Order, len(list))
	for i, o := range list {
		orders[i] = &protos.Order{
			Id:             int64(o.ID),
			Identification: o.Identification,
			PlacedAt:       timestamppb.New(o.PlacedAt),
			Status:         protos.OrderStatus(o.Status),
			DeliveryMode:   protos.DeliveryMode(o.DeliveryMode),
			Observation:    o.Observations,
			FinalPrice:     int64(o.FinalPrice),
			Address:        o.Address,
			Phone:          o.Phone,
		}
	}

	return &protos.ListOrdersResponse{
		Orders: orders,
	}, nil
}

func (s service) GetOrderByID(ctx context.Context, request *protos.GetOrderByIdRequest) (*protos.GetOrderByIdResponse, error) {
	o, err := s.ordering.GetOrderByID(ctx, types.ID(request.Id))
	if err != nil {
		return nil, failures.HandleError(throw.Error(err))
	}

	return &protos.GetOrderByIdResponse{
		Order: &protos.Order{
			Id:             int64(o.ID),
			Identification: o.Identification,
			PlacedAt:       timestamppb.New(o.PlacedAt),
			Status:         protos.OrderStatus(o.Status),
			DeliveryMode:   protos.DeliveryMode(o.DeliveryMode),
			Observation:    o.Observations,
			FinalPrice:     int64(o.FinalPrice),
			Address:        o.Address,
			Phone:          o.Phone,
		},
	}, nil
}

func (s service) CancelOrder(ctx context.Context, request *protos.CancelOrderRequest) (*protos.Empty, error) {
	err := s.ordering.CancelOrder(ctx, types.ID(request.Id))
	if err != nil {
		return nil, throw.Error(err)
	}

	return &protos.Empty{}, nil
}

func (s service) ListOrdersInFlow(_ *protos.Empty, server protos.Ordering_ListOrdersInFlowServer) error {
	var lastRun time.Time

	for {
		lastRun = time.Now()
		orders, err := s.ordering.ListOrders(server.Context(), order.Filter{
			PlacedAfter: lastRun,
		})
		if err != nil {
			return throw.Error(err)
		}

		for _, o := range orders {
			err := server.Send(&protos.ListOrdersInFlowResponse{
				Order: &protos.Order{
					Id:             int64(o.ID),
					Identification: o.Identification,
					PlacedAt:       timestamppb.New(o.PlacedAt),
					Observation:    o.Observations,
					FinalPrice:     int64(o.FinalPrice),
					Address:        o.Address,
					Phone:          o.Phone,
				},
			})
			if err != nil {
				return throw.Error(err)
			}
		}

		time.Sleep(time.Second * 5)
	}
}

func NewService(server *grpc.Server, ordering ordering.Workflow) protos.OrderingServer {
	s := service{
		ordering: ordering,
	}

	protos.RegisterOrderingServer(server, s)
	return s
}
