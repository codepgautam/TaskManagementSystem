package service

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/codepgautam/TaskManagementSystem/internal/domain"
)

var (
	ErrInvalidTaskData = errors.New("invalid task data")
	ErrTaskNotFound    = errors.New("task not found")
)

// TaskService implements the domain.TaskService interface
type TaskService struct {
	repo domain.TaskRepository
}

// NewTaskService creates a new task service
func NewTaskService(repo domain.TaskRepository) *TaskService {
	return &TaskService{
		repo: repo,
	}
}

// CreateTask creates a new task with validation
func (s *TaskService) CreateTask(title, description string) (*domain.Task, error) {
	// Validate input
	if strings.TrimSpace(title) == "" {
		return nil, ErrInvalidTaskData
	}
	
	// Create new task
	task := &domain.Task{
		ID:          uuid.New().String(),
		Title:       strings.TrimSpace(title),
		Description: strings.TrimSpace(description),
		Status:      domain.TaskStatusPending,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	
	// Store the task
	if err := s.repo.Create(task); err != nil {
		return nil, err
	}
	
	return task, nil
}

// GetTask retrieves a task by ID
func (s *TaskService) GetTask(id string) (*domain.Task, error) {
	if strings.TrimSpace(id) == "" {
		return nil, ErrInvalidTaskData
	}
	
	task, err := s.repo.GetByID(id)
	if err != nil {
		return nil, ErrTaskNotFound
	}
	
	return task, nil
}

// GetTasks retrieves tasks with filtering and pagination
func (s *TaskService) GetTasks(filter domain.TaskFilter, pagination domain.Pagination) ([]*domain.Task, int, error) {
	// Set default pagination values
	if pagination.Page < 1 {
		pagination.Page = 1
	}
	if pagination.PageSize < 1 || pagination.PageSize > 100 {
		pagination.PageSize = 10
	}
	
	return s.repo.GetAll(filter, pagination)
}

// UpdateTask updates an existing task
func (s *TaskService) UpdateTask(id, title, description string, status domain.TaskStatus) (*domain.Task, error) {
	if strings.TrimSpace(id) == "" {
		return nil, ErrInvalidTaskData
	}
	
	// Get existing task
	existingTask, err := s.repo.GetByID(id)
	if err != nil {
		return nil, ErrTaskNotFound
	}
	
	// Update fields
	if strings.TrimSpace(title) != "" {
		existingTask.Title = strings.TrimSpace(title)
	}
	if strings.TrimSpace(description) != "" {
		existingTask.Description = strings.TrimSpace(description)
	}
	if status != "" {
		existingTask.Status = status
	}
	
	existingTask.UpdatedAt = time.Now()
	
	// Save updated task
	if err := s.repo.Update(existingTask); err != nil {
		return nil, err
	}
	
	return existingTask, nil
}

// DeleteTask removes a task by ID
func (s *TaskService) DeleteTask(id string) error {
	if strings.TrimSpace(id) == "" {
		return ErrInvalidTaskData
	}
	
	// Check if task exists
	_, err := s.repo.GetByID(id)
	if err != nil {
		return ErrTaskNotFound
	}
	
	return s.repo.Delete(id)
}
