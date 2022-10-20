package main

import (
	"comies/app/config"
	"comies/app/data/conn"
	"comies/app/data/ids"
	"comies/app/telemetry"
	"fmt"
	"log"
	"net"
	"os"
)

// @title        Comies Backend API
// @version      1.0
// @description  An API wrapping all functionalities of the ordering and menu services of Comies
func main() {
	cfg := config.Load()

	pool, err := conn.Connect(cfg.Database)
	if err != nil {
		log.Fatalf("Could not startup database: %v", err)
	}

	err = conn.Migrate(pool)
	if err != nil {
		log.Fatalf("Could not migrate database: %v", err)
	}

	err = ids.NewNode(cfg.IDGeneration)
	if err != nil {
		log.Fatalf("Could not startup idgen: %v", err)
	}

	telemetry.Register(&telemetry.Telemetry{
		Logger: telemetry.NewLogger(os.Stdout),
	})

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
