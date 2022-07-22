package middleware

import (
	"comies/app"
)

type (
	Middlewares struct {
		managers app.Managers
	}
)

func NewMiddlewares(managers app.Managers) Middlewares {
	return Middlewares{
		managers: managers,
	}
}
