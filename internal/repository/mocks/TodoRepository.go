// Code generated by mockery v2.28.1. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	models "github.com/wahyudibo/go-todo-api/internal/repository/models"
)

// TodoRepository is an autogenerated mock type for the TodoRepository type
type TodoRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: description
func (_m *TodoRepository) Create(description string) error {
	ret := _m.Called(description)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(description)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: todoID
func (_m *TodoRepository) Delete(todoID int64) (bool, error) {
	ret := _m.Called(todoID)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(int64) (bool, error)); ok {
		return rf(todoID)
	}
	if rf, ok := ret.Get(0).(func(int64) bool); ok {
		r0 = rf(todoID)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(todoID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: todoID
func (_m *TodoRepository) GetByID(todoID int64) (*models.Todo, error) {
	ret := _m.Called(todoID)

	var r0 *models.Todo
	var r1 error
	if rf, ok := ret.Get(0).(func(int64) (*models.Todo, error)); ok {
		return rf(todoID)
	}
	if rf, ok := ret.Get(0).(func(int64) *models.Todo); ok {
		r0 = rf(todoID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Todo)
		}
	}

	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(todoID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields:
func (_m *TodoRepository) List() ([]models.Todo, error) {
	ret := _m.Called()

	var r0 []models.Todo
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]models.Todo, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []models.Todo); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Todo)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: todoID, updates
func (_m *TodoRepository) Update(todoID int64, updates map[string]interface{}) (*models.Todo, error) {
	ret := _m.Called(todoID, updates)

	var r0 *models.Todo
	var r1 error
	if rf, ok := ret.Get(0).(func(int64, map[string]interface{}) (*models.Todo, error)); ok {
		return rf(todoID, updates)
	}
	if rf, ok := ret.Get(0).(func(int64, map[string]interface{}) *models.Todo); ok {
		r0 = rf(todoID, updates)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Todo)
		}
	}

	if rf, ok := ret.Get(1).(func(int64, map[string]interface{}) error); ok {
		r1 = rf(todoID, updates)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewTodoRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewTodoRepository creates a new instance of TodoRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTodoRepository(t mockConstructorTestingTNewTodoRepository) *TodoRepository {
	mock := &TodoRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
