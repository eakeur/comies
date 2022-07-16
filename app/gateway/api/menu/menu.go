package menu

import (
	"comies/app/core/entities/product"
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

func ExternalProductType(p product.Type) client.ProductType {
	switch p {
	case product.InputType:
		return client.ProductType_INPUT
	case product.OutputType:
		return client.ProductType_INPUT
	default:
		return client.ProductType_OUTPUT
	}
}

func InternalProductType(p client.ProductType) product.Type {
	return product.Type(p.Descriptor().Name())
}
