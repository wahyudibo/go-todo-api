package repository

import "github.com/wahyudibo/go-todo-api/internal/repository/models"

// TodoRepository provides method to interact with todo resource.
type TodoRepository interface {
	// List lists all todos
	List() ([]models.Todo, error)
	// GetByID gets todo by its id
	GetByID(todoID int64) (*models.Todo, error)
	// Create creates new todo
	Create(description string) error
	// Update updates / modifies existing todo by its id
	Update(todoID int64, updates map[string]interface{}) (*models.Todo, error)
	// Delete deletes todo by its id
	Delete(todoID int64) (bool, error)
}
