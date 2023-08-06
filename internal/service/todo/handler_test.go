package todoservice

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	mockStorageAdapter "github.com/wahyudibo/go-todo-api/internal/adapter/storage/mocks"
	"github.com/wahyudibo/go-todo-api/internal/dto"
	mockRepo "github.com/wahyudibo/go-todo-api/internal/repository/mocks"
	"github.com/wahyudibo/go-todo-api/internal/repository/models"
)

func TestTodoService(t *testing.T) {
	t.Run("testTodoServiceListFailedInternalServerError", testTodoServiceListFailedInternalServerError)
	t.Run("testTodoServiceListEmptyList", testTodoServiceListEmptyList)
	t.Run("testTodoServiceListSuccess", testTodoServiceListSuccess)
	t.Run("testTodoServiceGetByIDFailedInvalidTodoID", testTodoServiceGetByIDFailedInvalidTodoID)
	t.Run("testTodoServiceGetByIDFailedInternalServerError", testTodoServiceGetByIDFailedInternalServerError)
	t.Run("testTodoServiceGetByIDFailedTodoIDNotFound", testTodoServiceGetByIDFailedTodoIDNotFound)
	t.Run("testTodoServiceGetByIDSuccess", testTodoServiceGetByIDSuccess)
	t.Run("testTodoServiceCreateFailedInvalidRequestBody", testTodoServiceCreateFailedInvalidRequestBody)
	t.Run("testTodoServiceCreateFailedInternalServerError", testTodoServiceCreateFailedInternalServerError)
	t.Run("testTodoServiceCreateSuccess", testTodoServiceCreateSuccess)
	t.Run("testTodoServiceUpdateFailedInvalidTodoID", testTodoServiceUpdateFailedInvalidTodoID)
	t.Run("testTodoServiceUpdateFailedInvalidRequestBody", testTodoServiceUpdateFailedInvalidRequestBody)
	t.Run("testTodoServiceUpdateFailedInternalServerError", testTodoServiceUpdateFailedInternalServerError)
	t.Run("testTodoServiceUpdateFailedNotFound", testTodoServiceUpdateFailedNotFound)
	t.Run("testTodoServiceUpdateSuccess", testTodoServiceUpdateSuccess)
	t.Run("testTodoServiceDeleteFailedInvalidTodoID", testTodoServiceDeleteFailedInvalidTodoID)
	t.Run("testTodoServiceDeleteFailedInternalServerError", testTodoServiceDeleteFailedInternalServerError)
	t.Run("testTodoServiceDeleteFailedTodoIDNotFound", testTodoServiceDeleteFailedTodoIDNotFound)
	t.Run("testTodoServiceDeleteSuccess", testTodoServiceDeleteSuccess)
}

func testTodoServiceListFailedInternalServerError(t *testing.T) {
	todoRepo := mockRepo.NewTodoRepository(t)
	expectedErr := errors.New("Internal server error")
	todoRepo.On("List").Return(nil, expectedErr)

	storageAdapter := mockStorageAdapter.NewStorageAdapter(t)

	todoSvc := NewTodoService(todoRepo, storageAdapter)
	endpoint := "/api/todos"

	r := chi.NewRouter()
	r.Get(endpoint, todoSvc.List)
	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	require.NoError(t, err)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	respBody, err := io.ReadAll(w.Body)
	require.NoError(t, err)

	var errResp dto.ErrorResponse
	err = json.Unmarshal(respBody, &errResp)
	require.NoError(t, err)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, expectedErr.Error(), errResp.Message)
}

func testTodoServiceListEmptyList(t *testing.T) {
	todoRepo := mockRepo.NewTodoRepository(t)
	emptyTodoList := make([]models.Todo, 0)
	todoRepo.On("List").Return(emptyTodoList, nil)

	storageAdapter := mockStorageAdapter.NewStorageAdapter(t)

	todoSvc := NewTodoService(todoRepo, storageAdapter)
	endpoint := "/api/todos"

	r := chi.NewRouter()
	r.Get(endpoint, todoSvc.List)
	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	require.NoError(t, err)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	respBody, err := io.ReadAll(w.Body)
	require.NoError(t, err)

	var result []models.Todo
	err = json.Unmarshal(respBody, &result)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, emptyTodoList, result)
}

func testTodoServiceListSuccess(t *testing.T) {}

func testTodoServiceGetByIDFailedInvalidTodoID(t *testing.T) {}

func testTodoServiceGetByIDFailedInternalServerError(t *testing.T) {}

func testTodoServiceGetByIDFailedTodoIDNotFound(t *testing.T) {}

func testTodoServiceGetByIDSuccess(t *testing.T) {}

func testTodoServiceCreateFailedInvalidRequestBody(t *testing.T) {}

func testTodoServiceCreateFailedInternalServerError(t *testing.T) {}

func testTodoServiceCreateSuccess(t *testing.T) {}

func testTodoServiceUpdateFailedInvalidTodoID(t *testing.T) {}

func testTodoServiceUpdateFailedInvalidRequestBody(t *testing.T) {}

func testTodoServiceUpdateFailedInternalServerError(t *testing.T) {}

func testTodoServiceUpdateFailedNotFound(t *testing.T) {}

func testTodoServiceUpdateSuccess(t *testing.T) {}

func testTodoServiceDeleteFailedInvalidTodoID(t *testing.T) {}

func testTodoServiceDeleteFailedInternalServerError(t *testing.T) {}

func testTodoServiceDeleteFailedTodoIDNotFound(t *testing.T) {}

func testTodoServiceDeleteSuccess(t *testing.T) {}
