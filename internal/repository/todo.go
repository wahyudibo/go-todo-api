package repository

// TodoRepository provides method to interact with todo resource.
type TodoRepository interface {
	// List lists all todos
	List() error
	// GetByID gets todo by its id
	GetByID() error
	// Create creates new todo
	Create() error
	// Update updates / modifies existing todo by its id
	Update() error
	// Delete deletes todo by its id
	Delete() error
}
