package ordering

import (
	"comies/app/core/workflows/ordering"
	"comies/app/gateway/api/errors"
	"comies/app/gateway/api/gen/ordering/protos"
	"google.golang.org/grpc"
)

var _ protos.OrderingServer = service{}

var failures = errors.ErrorBinding{}

type service struct {
	protos.UnimplementedOrderingServer
	ordering ordering.Workflow
}

func NewService(server *grpc.Server, ordering ordering.Workflow) protos.OrderingServer {
	s := service{
		ordering: ordering,
	}

	protos.RegisterOrderingServer(server, s)
	return s
}
