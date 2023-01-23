package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"

	"github.com/wahyudibo/go-todo-api/internal/config"
	"github.com/wahyudibo/go-todo-api/internal/database/mysql"
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
	db, err := mysqldb.Connect(ctx, &cfg)
	if err != nil {
		log.Fatalf("failed when initiating database: %v", err)
	}

	// migrate the database
	err = mysqldb.Migrate(db)
	if err != nil {
		log.Fatalf("Error running schema migration %v", err)
	}

	// initialize repositories
	todoRepo := mysqldb.NewTodoRepository(db)

	// initializes service
	todoService := todoservice.NewTodoService(todoRepo)

	// initialize router
	routeBuilder := router.New(todoService)
	router := routeBuilder.Build()

	log.Infof("starting server at port: %d", cfg.AppPort)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.AppPort), router); err != nil {
		log.Fatalf("failed to start app at port %d: %v", cfg.AppPort, err)
	}
}
