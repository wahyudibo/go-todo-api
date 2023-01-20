package main

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"

	"github.com/wahyudibo/go-todo-api/internal/config"
	"github.com/wahyudibo/go-todo-api/internal/database/mysql"
)

func main() {
	ctx := context.Background()

	// logger settings
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.InfoLevel)

	// initialize config
	var cfg config.Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatalf("failed when parsing config: %v", err)
	}

	// initialize db connection
	db, err := mysqldb.Connect(ctx, &cfg)
	if err != nil {
		log.Fatalf("failed when initiating database: %v", err)
	}

	// initialize repositories
	todoRepo := mysqldb.NewTodoRepository(db)

	// initializes service

}
