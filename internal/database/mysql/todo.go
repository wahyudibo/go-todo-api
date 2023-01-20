package mysqldb

import (
	"gorm.io/gorm"

	"github.com/wahyudibo/go-todo-api/internal/repository"
)

type todoRepository struct {
	client *gorm.DB
}

func NewTodoRepository(dbClient *gorm.DB) repository.TodoRepository {
	return &todoRepository{
		client: dbClient,
	}
}

func (r *todoRepository) List() error {
	return nil
}

func (r *todoRepository) GetByID() error {
	return nil
}

func (r *todoRepository) Create() error {
	return nil
}

func (r *todoRepository) Update() error {
	return nil
}

func (r *todoRepository) Delete() error {
	return nil
}
