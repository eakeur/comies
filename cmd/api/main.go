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
		User:      cfg.Database.User,
		Password:  cfg.Database.Password,
		Host:      cfg.Database.Host,
		Port:      cfg.Database.Port,
		Name:      cfg.Database.Name,
		SSLMode:   cfg.Database.SSLMode,
		Migration: cfg.Database.MigrationsPath,
	})
	if err != nil {
		log.Fatalf("Could not connect and populate postgres database: %v", err)
	}

	nodeNumber, err := strconv.Atoi(cfg.IDGeneration.NodeNumber)
	if err != nil {
		log.Fatalf("Could not parse id generation node number %v: %v", cfg.IDGeneration.NodeNumber, err)
	}

	snflake, err := snowflake.NewNode(int64(nodeNumber))
	if err != nil {
		log.Fatalf("Could not create snowflake node: %v", err)
	}

	application := app.NewApplication(app.Gateways{
		Database:      db,
		SnowflakeNode: snflake,
	})

	address := fmt.Sprintf(":%v", cfg.Server.ListenPort)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Could not listen to port %v: %v", cfg.Server.ListenPort, err)
	}

	srv := api.NewAPI(application, nil)

	log.Printf("Listening on address https://%v", address)
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("Server stopped listenin on port %v: %v", cfg.Server.ListenPort, err)
	}

}
