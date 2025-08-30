package domain

import (
	"time"
)

// TaskStatus represents the status of a task
type TaskStatus string

const (
	TaskStatusPending    TaskStatus = "Pending"
	TaskStatusInProgress TaskStatus = "InProgress"
	TaskStatusCompleted  TaskStatus = "Completed"
)

// Task represents a task entity
type Task struct {
	ID          string      `json:"id"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Status      TaskStatus  `json:"status"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
}

// TaskFilter represents filtering options for tasks
type TaskFilter struct {
	Status *TaskStatus `json:"status,omitempty"`
}

// Pagination represents pagination parameters
type Pagination struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

// TaskRepository defines the interface for task storage operations
type TaskRepository interface {
	Create(task *Task) error
	GetByID(id string) (*Task, error)
	GetAll(filter TaskFilter, pagination Pagination) ([]*Task, int, error)
	Update(task *Task) error
	Delete(id string) error
}

// TaskService defines the interface for task business logic
type TaskService interface {
	CreateTask(title, description string) (*Task, error)
	GetTask(id string) (*Task, error)
	GetTasks(filter TaskFilter, pagination Pagination) ([]*Task, int, error)
	UpdateTask(id, title, description string, status TaskStatus) (*Task, error)
	DeleteTask(id string) error
}
