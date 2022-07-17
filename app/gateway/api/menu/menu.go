package menu

import (
	"comies/app/core/workflows/menu"
	"comies/app/core/workflows/stocking"
	client "comies/app/gateway/api/gen/menu"
	"google.golang.org/grpc"
)

var _ client.MenuServer = service{}

type service struct {
	client.UnimplementedMenuServer
	menu   menu.Workflow
	stocks stocking.Workflow
}

func NewService(server *grpc.Server, menu menu.Workflow, stocks stocking.Workflow) client.MenuServer {
	s := service{
		menu:   menu,
		stocks: stocks,
	}

	client.RegisterMenuServer(server, s)

	return s
}
