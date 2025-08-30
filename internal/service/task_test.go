package service

import (
	"testing"

	"github.com/codepgautam/TaskManagementSystem/internal/domain"
	"github.com/codepgautam/TaskManagementSystem/internal/repository"
)

func TestTaskService_CreateTask(t *testing.T) {
	repo := repository.NewMemoryTaskRepository()
	service := NewTaskService(repo)
	
	tests := []struct {
		name        string
		title       string
		description string
		wantErr     bool
	}{
		{
			name:        "valid task",
			title:       "Test Task",
			description: "Test Description",
			wantErr:     false,
		},
		{
			name:        "empty title",
			title:       "",
			description: "Test Description",
			wantErr:     true,
		},
		{
			name:        "whitespace title",
			title:       "   ",
			description: "Test Description",
			wantErr:     true,
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			task, err := service.CreateTask(tt.title, tt.description)
			
			if tt.wantErr {
				if err == nil {
					t.Error("expected error but got none")
				}
				return
			}
			
			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}
			
			if task.ID == "" {
				t.Error("expected task ID to be generated")
			}
			
			if task.Status != domain.TaskStatusPending {
				t.Errorf("expected status to be Pending, got %v", task.Status)
			}
		})
	}
}

func TestTaskService_GetTask(t *testing.T) {
	repo := repository.NewMemoryTaskRepository()
	service := NewTaskService(repo)
	
	// Create a test task
	task, err := service.CreateTask("Test Task", "Test Description")
	if err != nil {
		t.Fatalf("failed to create test task: %v", err)
	}
	
	tests := []struct {
		name    string
		id      string
		wantErr bool
	}{
		{
			name:    "existing task",
			id:      task.ID,
			wantErr: false,
		},
		{
			name:    "non-existing task",
			id:      "non-existing-id",
			wantErr: true,
		},
		{
			name:    "empty id",
			id:      "",
			wantErr: true,
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := service.GetTask(tt.id)
			
			if tt.wantErr {
				if err == nil {
					t.Error("expected error but got none")
				}
				return
			}
			
			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}
			
			if result.ID != tt.id {
				t.Errorf("expected ID %v, got %v", tt.id, result.ID)
			}
		})
	}
}

func TestTaskService_UpdateTask(t *testing.T) {
	repo := repository.NewMemoryTaskRepository()
	service := NewTaskService(repo)
	
	// Create a test task
	task, err := service.CreateTask("Original Title", "Original Description")
	if err != nil {
		t.Fatalf("failed to create test task: %v", err)
	}
	
	// Update the task
	updatedTask, err := service.UpdateTask(
		task.ID,
		"Updated Title",
		"Updated Description",
		domain.TaskStatusCompleted,
	)
	
	if err != nil {
		t.Errorf("unexpected error: %v", err)
		return
	}
	
	if updatedTask.Title != "Updated Title" {
		t.Errorf("expected title 'Updated Title', got %v", updatedTask.Title)
	}
	
	if updatedTask.Status != domain.TaskStatusCompleted {
		t.Errorf("expected status Completed, got %v", updatedTask.Status)
	}
	
	// Test updating non-existing task
	_, err = service.UpdateTask("non-existing", "Title", "Description", domain.TaskStatusPending)
	if err == nil {
		t.Error("expected error for non-existing task")
	}
}
