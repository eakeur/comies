package menu

import (
	"comies/app/core/workflows/menu"
	client "comies/app/gateway/api/gen/menu"

	"google.golang.org/grpc"
)

var _ client.MenuServer = service{}

type service struct {
	client.UnimplementedMenuServer
	menu menu.Workflow
}

func NewService(server *grpc.Server, menu menu.Workflow) client.MenuServer {
	s := service{
		menu: menu,
	}

	client.RegisterMenuServer(server, s)

	return s
}
