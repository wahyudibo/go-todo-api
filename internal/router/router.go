package router

import (
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"

	todoservice "github.com/wahyudibo/go-todo-api/internal/service/todo"
)

type RouteBuilder struct {
	todoService *todoservice.Handler
}

func New(todoSvc *todoservice.Handler) *RouteBuilder {
	return &RouteBuilder{
		todoService: todoSvc,
	}
}

func (rb *RouteBuilder) Build() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/api/todos", rb.todoService.List)
	r.Get("/api/todos/{todoId}", rb.todoService.GetByID)
	r.Post("/api/todos", rb.todoService.Create)
	r.Put("/api/todos/{todoId}", rb.todoService.Update)
	r.Delete("/api/todos/{todoId}", rb.todoService.Delete)

	return r
}
