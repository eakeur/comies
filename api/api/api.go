package api

import (
	"comies/telemetry"
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
)

func Serve(lis net.Listener, handler http.Handler, writeTimeout, readTimeout time.Duration) {
	log := telemetry.Logger()

	//nolint:gosec
	srv := &http.Server{
		Handler:      handler,
		WriteTimeout: writeTimeout,
		ReadTimeout:  readTimeout,
	}

	c := make(chan os.Signal, 1)
	idleConnections := make(chan struct{})
	signal.Notify(c, os.Interrupt, syscall.SIGINT)

	go func() {
		<-c
		// create context with timeout
		ctx, cancel := context.WithTimeout(context.Background(), writeTimeout)
		defer cancel()

		// start http shutdown
		if err := srv.Shutdown(ctx); err != nil {
			log.Error("Shutdown " + err.Error())
		}

		close(idleConnections)
	}()

	log.Info(fmt.Sprintf("Listening at %s", lis.Addr().String()))
	if err := srv.Serve(lis); err != nil && err != http.ErrServerClosed {
		log.Fatal("Listen and serve failed", zap.Error(err))
	}

	log.Info("Waiting idle connections...")
	<-idleConnections

	log.Info("Bye bye")
}
