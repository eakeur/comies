package main

import (
	"comies/app"
	"comies/app/config"
	"comies/app/gateway/api"
	"comies/app/gateway/persistence/postgres"
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"

	"github.com/bwmarrin/snowflake"
)

func main() {
	ctx := context.Background()

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Could not load configurations from environment: %v", err)
	}

	db, err := postgres.ConnectAndMount(ctx, postgres.Config{
		User:     cfg.Database.User,
		Password: cfg.Database.Password,
		Host:     cfg.Database.Host,
		Name:     cfg.Database.Name,
		SSLMode:  cfg.Database.SSLMode,
	})
	if err != nil {
		log.Fatalf("Could not connect and populate postgres database: %v", err)
	}
	log.Printf("Successfully connected to database %s", db.Config().ConnConfig.Database)

	nodeNumber, err := strconv.Atoi(cfg.IDGeneration.NodeNumber)
	if err != nil {
		log.Fatalf("Could not parse id generation node number %v: %v", cfg.IDGeneration.NodeNumber, err)
	}

	snflake, err := snowflake.NewNode(int64(nodeNumber))
	if err != nil {
		log.Fatalf("Could not create snowflake node: %v", err)
	}
	log.Println("Successfully created snowflake node")

	application := app.NewApplication(app.Gateways{
		Database:      db,
		SnowflakeNode: snflake,
	})

	address := fmt.Sprintf(":%v", "8080")
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Could not listen to port %v: %v", cfg.Server.ListenPort, err)
	}
	log.Printf("Listening on %s", lis.Addr())

	err = http.Serve(lis, api.NewAPI(application))
	if err != nil {
		log.Fatalf("Server stopped listening on port %v: %v", cfg.Server.ListenPort, err)
	}
}
