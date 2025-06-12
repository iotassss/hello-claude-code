package models

import (
	"time"
	"gorm.io/gorm"
)

type Task struct {
	ID          uint           `json:"id" gorm:"primarykey"`
	Title       string         `json:"title" gorm:"not null;size:100" validate:"required,max=100"`
	Description string         `json:"description" gorm:"size:500"`
	Status      string         `json:"status" gorm:"not null;default:pending" validate:"required,oneof=pending in_progress completed"`
	Priority    string         `json:"priority" gorm:"not null;default:medium" validate:"required,oneof=low medium high critical"`
	DueDate     *time.Time     `json:"due_date"`
	CategoryID  *uint          `json:"category_id"`
	Category    *Category      `json:"category,omitempty" gorm:"foreignKey:CategoryID"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}

type CreateTaskRequest struct {
	Title       string     `json:"title" validate:"required,max=100"`
	Description string     `json:"description" validate:"max=500"`
	Priority    string     `json:"priority" validate:"oneof=low medium high critical"`
	DueDate     *time.Time `json:"due_date"`
	CategoryID  *uint      `json:"category_id"`
}

type UpdateTaskRequest struct {
	Title       *string    `json:"title" validate:"omitempty,max=100"`
	Description *string    `json:"description" validate:"omitempty,max=500"`
	Status      *string    `json:"status" validate:"omitempty,oneof=pending in_progress completed"`
	Priority    *string    `json:"priority" validate:"omitempty,oneof=low medium high critical"`
	DueDate     *time.Time `json:"due_date"`
	CategoryID  *uint      `json:"category_id"`
}