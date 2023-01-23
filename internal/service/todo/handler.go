package todoservice

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	log "github.com/sirupsen/logrus"

	"github.com/wahyudibo/go-todo-api/internal/dto"
	"github.com/wahyudibo/go-todo-api/internal/repository"
	"github.com/wahyudibo/go-todo-api/internal/repository/models"
)

type Handler struct {
	TodoRepository repository.TodoRepository
}

func NewTodoService(todoRepo repository.TodoRepository) *Handler {
	return &Handler{
		TodoRepository: todoRepo,
	}
}

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	todos, err := h.TodoRepository.List()
	if err != nil {
		log.Errorf("failed when listing todo: %v", err)
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, &dto.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, todos)
}

func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	todoIDURLParam := chi.URLParam(r, "todoId")
	todoID, err := strconv.ParseInt(todoIDURLParam, 10, 64)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, &dto.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	todo, err := h.TodoRepository.GetByID(todoID)
	if err != nil {
		log.Errorf("failed when get todo by id: %v", err)
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, &dto.ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	if todo == nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, &dto.ErrorResponse{
			Message: models.ErrRecordNotFound,
		})
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, todo)
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	reqBody := dto.CreateTodoRequest{}
	if err := render.DecodeJSON(r.Body, &reqBody); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, &dto.ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	if err := reqBody.Validate(); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, &dto.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	if err := h.TodoRepository.Create(reqBody.Description); err != nil {
		log.Errorf("failed when creating todo: %v", err)
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, &dto.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	todoIDURLParam := chi.URLParam(r, "todoId")
	todoID, err := strconv.ParseInt(todoIDURLParam, 10, 64)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, &dto.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	reqBody := dto.UpdateTodoRequest{}
	if err := render.DecodeJSON(r.Body, &reqBody); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, &dto.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	rowsAffected, err := h.TodoRepository.Update(todoID, reqBody.ToModel())
	if err != nil {
		log.Errorf("failed when updating todo: %v", err)
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, &dto.ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	if rowsAffected == 0 {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, &dto.ErrorResponse{
			Message: models.ErrRecordNotFound,
		})
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	todoIDURLParam := chi.URLParam(r, "todoId")
	todoID, err := strconv.ParseInt(todoIDURLParam, 10, 64)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, &dto.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	rowsAffected, err := h.TodoRepository.Delete(todoID)
	if err != nil {
		log.Errorf("failed when deleting todo: %v", err)
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, &dto.ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	if rowsAffected == 0 {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, &dto.ErrorResponse{
			Message: models.ErrRecordNotFound,
		})
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
