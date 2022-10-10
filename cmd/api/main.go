package main

import (
	"comies/app/config"
	"comies/app/core/id"
	"comies/app/data/conn"
	"fmt"
	"log"
	"net"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// @title        Comies Backend API
// @version      1.0
// @description  An API wrapping all functionalities of the ordering and menu services of Comies
func main() {
	cfg := config.Load()

	err := conn.Connect(cfg.Database)
	if err != nil {
		log.Fatalf("Could not startup database: %v", err)
	}

	err = conn.Migrate()
	if err != nil {
		log.Fatalf("Could not migrate database: %v", err)
	}

	err = id.NewNode(cfg.IDGeneration)
	if err != nil {
		log.Fatalf("Could not startup idgen: %v", err)
	}

	err = startupLogger(cfg.Logger)
	if err != nil {
		log.Fatalf("Could not startup logger: %v", err)
	}

	address := fmt.Sprintf(":%v", cfg.Server.ListenPort)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Could not listen to port %v: %v", cfg.Server.ListenPort, err)
	}
	log.Printf("Listening on %s", lis.Addr())

	// err = http.Serve(lis, api.NewAPI(application))
	// if err != nil {
	// 	log.Fatalf("Server stopped listening on port %v: %v", cfg.Server.ListenPort, err)
	// }
}

func startupLogger(cfg config.Logger) error {

	stackTraceOption := zap.AddStacktrace(zapcore.PanicLevel)

	builderByEnv := map[string]func(options ...zap.Option) (*zap.Logger, error){
		"development": zap.NewDevelopment,
		"production":  zap.NewProduction,
	}

	logger, err := builderByEnv[cfg.Environment](stackTraceOption)
	if err != nil {
		return err
	}

	zap.ReplaceGlobals(logger)

	return nil
}
