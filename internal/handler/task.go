package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/codepgautam/TaskManagementSystem/internal/domain"
	"github.com/codepgautam/TaskManagementSystem/pkg/response"
)

// TaskHandler handles HTTP requests for tasks
type TaskHandler struct {
	service domain.TaskService
}

// NewTaskHandler creates a new task handler
func NewTaskHandler(service domain.TaskService) *TaskHandler {
	return &TaskHandler{
		service: service,
	}
}

// CreateTaskRequest represents the request body for creating a task
type CreateTaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

// UpdateTaskRequest represents the request body for updating a task
type UpdateTaskRequest struct {
	Title       string             `json:"title,omitempty"`
	Description string             `json:"description,omitempty"`
	Status      domain.TaskStatus  `json:"status,omitempty"`
}

// CreateTask handles POST /tasks
func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var req CreateTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid JSON payload")
		return
	}
	
	task, err := h.service.CreateTask(req.Title, req.Description)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}
	
	response.JSON(w, http.StatusCreated, task)
}

// GetTask handles GET /tasks/{id}
func (h *TaskHandler) GetTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	
	task, err := h.service.GetTask(id)
	if err != nil {
		response.Error(w, http.StatusNotFound, "Task not found")
		return
	}
	
	response.JSON(w, http.StatusOK, task)
}

// GetTasks handles GET /tasks with filtering and pagination
func (h *TaskHandler) GetTasks(w http.ResponseWriter, r *http.Request) {
	// Parse query parameters
	query := r.URL.Query()
	
	// Parse pagination (only HTTP parsing, no business logic)
	page, _ := strconv.Atoi(query.Get("page"))
	pageSize, _ := strconv.Atoi(query.Get("page_size"))
	
	// Use domain constructor for validation and defaults
	pagination := domain.NewPagination(page, pageSize)
	
	// Parse filters
	filter := domain.TaskFilter{}
	if statusParam := query.Get("status"); statusParam != "" {
		status := domain.TaskStatus(statusParam)
		filter.Status = &status
	}
	
	// Get tasks
	tasks, total, err := h.service.GetTasks(filter, pagination)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to retrieve tasks")
		return
	}
	
	// Calculate total pages
	totalPages := (total + pageSize - 1) / pageSize
	
	meta := &response.Meta{
		Page:       page,
		PageSize:   pageSize,
		Total:      total,
		TotalPages: totalPages,
	}
	
	response.JSONWithMeta(w, http.StatusOK, tasks, meta)
}

// UpdateTask handles PUT /tasks/{id}
func (h *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	
	var req UpdateTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid JSON payload")
		return
	}
	
	task, err := h.service.UpdateTask(id, req.Title, req.Description, req.Status)
	if err != nil {
		if err.Error() == "task not found" {
			response.Error(w, http.StatusNotFound, "Task not found")
		} else {
			response.Error(w, http.StatusBadRequest, err.Error())
		}
		return
	}
	
	response.JSON(w, http.StatusOK, task)
}

// DeleteTask handles DELETE /tasks/{id}
func (h *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	
	err := h.service.DeleteTask(id)
	if err != nil {
		if err.Error() == "task not found" {
			response.Error(w, http.StatusNotFound, "Task not found")
		} else {
			response.Error(w, http.StatusInternalServerError, "Failed to delete task")
		}
		return
	}
	
	response.JSON(w, http.StatusOK, map[string]string{"message": "Task deleted successfully"})
}
