package api

import (
	"comies/app"
	"comies/app/gateway/api/menu"
	"comies/app/gateway/api/middleware"
	"comies/app/gateway/api/ordering"
	"crypto/tls"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"

	"google.golang.org/grpc"
)

func NewAPI(application app.Application, cert *tls.Certificate) *grpc.Server {

	cred := grpc.Creds(insecure.NewCredentials())
	if cert != nil {
		cred = grpc.Creds(credentials.NewTLS(&tls.Config{
			Certificates: []tls.Certificate{*cert},
			ClientAuth:   tls.NoClientCert,
		}))
	}

	srv := grpc.NewServer(middleware.NewMiddlewares(application.Managers), cred)

	menu.NewService(srv, application.Menu)
	ordering.NewService(srv, application.Ordering)

	return srv
}
