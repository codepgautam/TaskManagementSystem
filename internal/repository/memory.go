package repository

import (
	"errors"
	"sort"
	"time"

	"github.com/codepgautam/TaskManagementSystem/internal/domain"
)

var (
	ErrTaskNotFound = errors.New("task not found")
)

// MemoryTaskRepository implements TaskRepository using in-memory storage
type MemoryTaskRepository struct {
	tasks map[string]*domain.Task
}

// NewMemoryTaskRepository creates a new in-memory task repository
func NewMemoryTaskRepository() *MemoryTaskRepository {
	return &MemoryTaskRepository{
		tasks: make(map[string]*domain.Task),
	}
}

// Create stores a new task
func (r *MemoryTaskRepository) Create(task *domain.Task) error {
	r.tasks[task.ID] = task
	return nil
}

// GetByID retrieves a task by its ID
func (r *MemoryTaskRepository) GetByID(id string) (*domain.Task, error) {
	task, exists := r.tasks[id]
	if !exists {
		return nil, ErrTaskNotFound
	}
	
	return task, nil
}

// GetAll retrieves all tasks with optional filtering and pagination
func (r *MemoryTaskRepository) GetAll(filter domain.TaskFilter, pagination domain.Pagination) ([]*domain.Task, int, error) {
	// Convert map to slice for easier manipulation
	allTasks := make([]*domain.Task, 0, len(r.tasks))
	for _, task := range r.tasks {
		// Apply status filter if provided
		if filter.Status != nil && task.Status != *filter.Status {
			continue
		}
		
		allTasks = append(allTasks, task)
	}
	
	// Sort by creation date (newest first)
	sort.Slice(allTasks, func(i, j int) bool {
		return allTasks[i].CreatedAt.After(allTasks[j].CreatedAt)
	})
	
	totalCount := len(allTasks)
	
	// Apply pagination (assumes pagination is already validated)
	start := (pagination.Page - 1) * pagination.PageSize
	end := start + pagination.PageSize
	
	if start >= len(allTasks) {
		return []*domain.Task{}, totalCount, nil
	}
	
	if end > len(allTasks) {
		end = len(allTasks)
	}
	
	return allTasks[start:end], totalCount, nil
}

// Update modifies an existing task
func (r *MemoryTaskRepository) Update(task *domain.Task) error {
	if _, exists := r.tasks[task.ID]; !exists {
		return ErrTaskNotFound
	}
	
	task.UpdatedAt = time.Now()
	r.tasks[task.ID] = task
	return nil
}

// Delete removes a task by its ID
func (r *MemoryTaskRepository) Delete(id string) error {
	if _, exists := r.tasks[id]; !exists {
		return ErrTaskNotFound
	}
	
	delete(r.tasks, id)
	return nil
}
