package api

import (
	"comies/app"
	"comies/app/gateway/api/menu"
	"google.golang.org/grpc"
)

func NewAPI(application app.Application) *grpc.Server {

	srv := grpc.NewServer()

	menu.NewService(srv, application.Menu, application.Stocking)

	return srv
}
