package ordering

import (
	"comies/app/core/entities/order"
	"comies/app/core/workflows/ordering"
	"comies/app/gateway/api/gen/ordering/protos"
	"comies/app/sdk/throw"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

var _ protos.OrderingServer = service{}

type service struct {
	protos.UnimplementedOrderingServer
	ordering ordering.Workflow
}

func (s service) ListOrdersInFlow(_ *protos.Empty, server protos.Ordering_ListOrdersInFlowServer) error {
	at := time.Now()

	orders, err := s.ordering.ListOrders(server.Context(), order.Filter{
		PlacedAfter: at,
	})
	if err != nil {
		return throw.Error(err)
	}

	for _, o := range orders {
		err := server.Send(&protos.ListOrdersInFlowResponse{
			Id:             int64(o.ID),
			Identification: o.Identification,
			PlacedAt:       timestamppb.New(o.PlacedAt),
			Observation:    o.Observations,
			FinalPrice:     int64(o.FinalPrice),
			Address:        o.Address,
			Phone:          o.Phone,
		})
		if err != nil {
			return throw.Error(err)
		}
	}

	return nil
}

func NewService(server *grpc.Server, ordering ordering.Workflow) protos.OrderingServer {
	s := service{
		ordering: ordering,
	}

	protos.RegisterOrderingServer(server, s)
	return s
}
