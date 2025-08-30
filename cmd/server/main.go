package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/pvnptl/task-management-system/internal/handler"
	"github.com/pvnptl/task-management-system/internal/repository"
	"github.com/pvnptl/task-management-system/internal/service"
)

func main() {
	// Get port from environment variable or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	
	// Initialize dependencies
	taskRepo := repository.NewMemoryTaskRepository()
	taskService := service.NewTaskService(taskRepo)
	taskHandler := handler.NewTaskHandler(taskService)
	
	// Setup router
	router := mux.NewRouter()
	
	// API routes
	api := router.PathPrefix("/api/v1").Subrouter()
	
	// Task routes
	api.HandleFunc("/tasks", taskHandler.CreateTask).Methods("POST")
	api.HandleFunc("/tasks", taskHandler.GetTasks).Methods("GET")
	api.HandleFunc("/tasks/{id}", taskHandler.GetTask).Methods("GET")
	api.HandleFunc("/tasks/{id}", taskHandler.UpdateTask).Methods("PUT")
	api.HandleFunc("/tasks/{id}", taskHandler.DeleteTask).Methods("DELETE")
	
	// Health check endpoint
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "healthy"}`))
	}).Methods("GET")
	
	// Start server
	log.Printf("Task Management Service starting on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
