package middleware

import (
	"comies/app"
	"google.golang.org/grpc"
)

type (
	Middlewares grpc.ServerOption
	middleware  struct {
		managers app.Managers
	}
)

func NewMiddlewares(managers app.Managers) Middlewares {
	mid := middleware{
		managers: managers,
	}
	return grpc.ChainUnaryInterceptor(mid.Logging(), mid.Transaction())
}
