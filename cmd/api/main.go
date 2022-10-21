package main

import (
	"comies/app/config"
	"comies/app/data/conn"
	"comies/app/data/ids"
	"comies/app/telemetry"
	"fmt"
	"net"
	"os"

	"go.uber.org/zap"
)

// @title        Comies Backend API
// @version      1.0
// @description  An API wrapping all functionalities of the ordering and menu services of Comies
func main() {
	cfg := config.Load()

	log := telemetry.NewLogger(os.Stdout)

	pool, err := conn.Connect(cfg.Database)
	if err != nil {
		log.Fatal("Could not startup database", zap.Error(err))
	}

	err = conn.Migrate(pool)
	if err != nil {
		log.Fatal("Could not migrate database", zap.Error(err))
	}

	err = ids.NewNode(cfg.IDGeneration)
	if err != nil {
		log.Fatal("Could not startup idgen", zap.Error(err))
	}

	address := fmt.Sprintf(":%v", cfg.Server.ListenPort)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("Could not listen to port", zap.String("port", cfg.Server.ListenPort), zap.Error(err))
	}

	log.Info("Server is ready to receive requests", zap.String("address", lis.Addr().String()))

	// err = http.Serve(lis, api.NewAPI(application))
	// if err != nil {
	// 	log.Fatalf("Server stopped listening on port %v: %v", cfg.Server.ListenPort, zap.Error(err))
	// }
}
