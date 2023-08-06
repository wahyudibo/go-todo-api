package mysqldb

import (
	"errors"
	"time"

	"gorm.io/gorm"

	"github.com/wahyudibo/go-todo-api/internal/repository"
	"github.com/wahyudibo/go-todo-api/internal/repository/models"
)

type todoRepository struct {
	client *gorm.DB
}

func NewTodoRepository(dbClient *gorm.DB) repository.TodoRepository {
	return &todoRepository{
		client: dbClient,
	}
}

func (r *todoRepository) List() ([]models.Todo, error) {
	var todos []models.Todo
	result := r.client.Find(&todos)
	return todos, result.Error
}

func (r *todoRepository) GetByID(todoID int64) (*models.Todo, error) {
	var todo models.Todo
	result := r.client.Where("id = ?", todoID).First(&todo)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &todo, result.Error
}

func (r *todoRepository) Create(description string) error {
	todo := models.Todo{
		Description: description,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}
	result := r.client.Create(&todo)
	return result.Error
}

func (r *todoRepository) Update(todoID int64, updates map[string]interface{}) (*models.Todo, error) {
	var todo models.Todo
	result := r.client.Model(&todo).Where("id = ?", todoID).Updates(updates)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &todo, result.Error
}

func (r *todoRepository) Delete(todoID int64) (bool, error) {
	todo := models.Todo{ID: todoID}
	result := r.client.Delete(&todo)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false, nil
	}
	return true, result.Error
}
