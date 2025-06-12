package models

import (
	"time"
	"gorm.io/gorm"
)

type Category struct {
	ID          uint           `json:"id" gorm:"primarykey"`
	Name        string         `json:"name" gorm:"not null;unique;size:50" validate:"required,max=50"`
	Color       string         `json:"color" gorm:"not null;default:#3498db" validate:"required,hexcolor"`
	Description string         `json:"description" gorm:"size:200"`
	TaskCount   int64          `json:"task_count" gorm:"-"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}

type CreateCategoryRequest struct {
	Name        string `json:"name" validate:"required,max=50"`
	Color       string `json:"color" validate:"hexcolor"`
	Description string `json:"description" validate:"max=200"`
}

type UpdateCategoryRequest struct {
	Name        *string `json:"name" validate:"omitempty,max=50"`
	Color       *string `json:"color" validate:"omitempty,hexcolor"`
	Description *string `json:"description" validate:"omitempty,max=200"`
}