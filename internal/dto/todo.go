package dto

import (
	"time"

	"github.com/go-ozzo/ozzo-validation/v4"
)

type CreateTodoRequest struct {
	Description string `json:"description"`
}

func (r CreateTodoRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Description, validation.Required),
	)
}

type UpdateTodoRequest struct {
	Description string `json:"description"`
	IsCompleted bool   `json:"is_completed"`
}

func (r *UpdateTodoRequest) ToModel() map[string]interface{} {
	updates := make(map[string]interface{}, 0)
	if r.Description != "" {
		updates["description"] = r.Description
	}
	updates["is_completed"] = r.IsCompleted
	updates["updated_at"] = time.Now().UTC()

	return updates
}

type UploadResponse struct {
	ObjectKey string `json:"object_key"`
}
