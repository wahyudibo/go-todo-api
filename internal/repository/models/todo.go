package models

import "time"

var ErrRecordNotFound = "record not found"

// Todo represents todo model in the database
type Todo struct {
	ID          int64 `gorm:"primaryKey"`
	Description string
	IsCompleted bool `gorm:"default:false"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
