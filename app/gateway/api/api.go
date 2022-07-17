package api

import (
	"comies/app"
	"comies/app/gateway/api/menu"
	"comies/app/gateway/api/middleware"

	"google.golang.org/grpc"
)

func NewAPI(application app.Application) *grpc.Server {

	srv := grpc.NewServer(middleware.NewMiddlewares(application.Managers))

	menu.NewService(srv, application.Menu)

	return srv
}
