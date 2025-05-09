package dto

import (
	"gorm.io/gorm"
	"time"
)

type BaseDtoResponse struct {
	ID        string         `json:"id,omitempty"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
