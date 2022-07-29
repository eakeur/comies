package main

import (
	"comies/app"
	"comies/app/config"
	"comies/app/gateway/api"
	"comies/app/gateway/persistence/postgres"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
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

	var db *pgxpool.Pool
	if cfg.Database.URL == "" {
		db, err = postgres.ConnectAndMount(ctx, postgres.Config{
			User:     cfg.Database.User,
			Password: cfg.Database.Password,
			Host:     cfg.Database.Host,
			Name:     cfg.Database.Name,
			SSLMode:  cfg.Database.SSLMode,
		})
	} else {
		db, err = postgres.ConnectAndMountURL(ctx, cfg.Database.URL)
	}
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

	var logger *zap.Logger
	if cfg.Logger.Environment == "development" {
		logger, err = zap.NewDevelopment(zap.AddStacktrace(zapcore.PanicLevel))
		if err != nil {
			log.Fatalf("Could not create logger instance: %v", err)
		}
		log.Println("Successfully created logger instance in development mode")
	} else {
		logger, err = zap.NewProduction(zap.AddStacktrace(zapcore.PanicLevel))
		if err != nil {
			log.Fatalf("Could not create logger instance: %v", err)
		}
		log.Println("Successfully created logger instance in production mode")
	}

	application := app.NewApplication(app.Gateways{
		Database:      db,
		SnowflakeNode: snflake,
		Logger:        logger.Sugar(),
	})

	address := fmt.Sprintf(":%v", cfg.Server.ListenPort)
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
