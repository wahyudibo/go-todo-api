package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"

	mysqldbadapter "github.com/wahyudibo/go-todo-api/internal/adapter/database/mysql"
	storageadapter "github.com/wahyudibo/go-todo-api/internal/adapter/storage"
	localstorageadapter "github.com/wahyudibo/go-todo-api/internal/adapter/storage/local"
	s3storageadapter "github.com/wahyudibo/go-todo-api/internal/adapter/storage/s3"
	"github.com/wahyudibo/go-todo-api/internal/config"
	"github.com/wahyudibo/go-todo-api/internal/router"
	"github.com/wahyudibo/go-todo-api/internal/service/todo"
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
	db, err := mysqldbadapter.Connect(ctx, &cfg)
	if err != nil {
		log.Fatalf("failed when initiating database: %v", err)
	}

	// migrate the database
	err = mysqldbadapter.Migrate(db)
	if err != nil {
		log.Fatalf("error running schema migration %v", err)
	}

	// initializes storage adapters
	var storageAdapter storageadapter.StorageAdapter
	switch cfg.StorageDriver {
	case "s3":
		storageAdapter, err = s3storageadapter.New(ctx, &cfg)
	case "local":
		storageAdapter, err = localstorageadapter.New(&cfg)
	default:
		log.Fatalf("unknown storage driver. Acceptable storage driver are: local, s3")
	}
	if err != nil {
		log.Fatalf("failed to initializes storage adapter: %+v\n", err)
	}

	// initialize repositories
	todoRepo := mysqldbadapter.NewTodoRepository(db)

	// initializes service
	todoService := todoservice.NewTodoService(todoRepo, storageAdapter)

	// initialize router
	routeBuilder := router.New(todoService)
	router := routeBuilder.Build()

	log.Infof("starting server at port: %d", cfg.AppPort)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.AppPort), router); err != nil {
		log.Fatalf("failed to start app at port %d: %v", cfg.AppPort, err)
	}
}
