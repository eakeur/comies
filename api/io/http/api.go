package http

import (
	"comies/telemetry"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
)

func Serve(endpoint string, handler http.Handler, writeTimeout, readTimeout time.Duration) {
	log := telemetry.Logger()

	//nolint:gosec
	srv := &http.Server{
		Handler:      handler,
		Addr:         endpoint,
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

	log.Info(fmt.Sprintf("Listening at %s", endpoint))
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("Listen and serve failed", zap.Error(err))
	}

	log.Info("Waiting idle connections...")
	<-idleConnections

	log.Info("Bye bye")
}
