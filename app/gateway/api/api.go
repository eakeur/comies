package api

import (
	"comies/app"
	"comies/app/gateway/api/menu"
	"comies/app/gateway/api/middleware"
	"comies/app/gateway/api/ordering"

	"google.golang.org/grpc"
)

func NewAPI(application app.Application) *grpc.Server {

	srv := grpc.NewServer(middleware.NewMiddlewares(application.Managers))

	menu.NewService(srv, application.Menu)
	ordering.NewService(srv, application.Ordering)

	return srv
}
