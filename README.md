# Task Management System

A microservice-based task management system built with Go, demonstrating clean architecture principles and microservices design patterns.

## Table of Contents

- [Overview](#overview)
- [Architecture](#architecture)
- [Design Decisions](#design-decisions)
- [Getting Started](#getting-started)
- [API Documentation](#api-documentation)
- [Microservices Concepts](#microservices-concepts)
- [Scaling Considerations](#scaling-considerations)
- [Inter-Service Communication](#inter-service-communication)

## Overview

This project implements a simple task management system as a microservice with the following features:

- **CRUD Operations**: Create, Read, Update, Delete tasks
- **Pagination**: Efficient handling of large task lists
- **Filtering**: Filter tasks by status
- **RESTful API**: Clean and consistent API design
- **Microservices Architecture**: Single responsibility and clear separation of concerns

## Architecture

The system follows a layered architecture pattern:

```
cmd/
├── server/          # Application entry point
internal/
├── domain/          # Domain models and interfaces
├── repository/      # Data access layer
├── service/         # Business logic layer
└── handler/         # HTTP handlers (presentation layer)
pkg/
└── response/        # Standardized API responses
```

### Key Components

1. **Domain Layer**: Defines core business entities and interfaces
2. **Repository Layer**: Handles data persistence (currently in-memory)
3. **Service Layer**: Contains business logic and validation
4. **Handler Layer**: HTTP request/response handling

## Design Decisions

### 1. Clean Architecture
- **Dependency Inversion**: High-level modules don't depend on low-level modules
- **Interface Segregation**: Small, focused interfaces
- **Single Responsibility**: Each layer has a distinct purpose

### 2. In-Memory Storage
- Chosen for simplicity and demonstration purposes
- Simple map-based storage for easy understanding
- Easy to replace with database implementation

### 3. Error Handling
- Standardized error responses
- Proper HTTP status codes
- Error wrapping for better debugging

### 4. API Design
- RESTful endpoints following HTTP conventions
- Consistent JSON response format
- Query parameters for filtering and pagination

## Getting Started

### Prerequisites

- Go 1.21 or higher
- Docker (optional, for containerized deployment)

### Installation

1. Clone the repository:
```bash
git clone https://github.com/codepgautam/TaskManagementSystem.git
cd TaskManagementSystem
```

2. Install dependencies:
```bash
go mod download
```

3. Run the service:
```bash
go run cmd/server/main.go
```

The service will start on port 8080 by default.

### Using Docker

1. Build and run with Docker Compose:
```bash
docker-compose up --build
```

2. Or build and run manually:
```bash
docker build -t task-management-system .
docker run -p 8080:8080 task-management-system
```

### Running Tests

```bash
go test ./...
```

## API Documentation

Base URL: `http://localhost:8080/api/v1`

### Endpoints

#### 1. Create Task
- **POST** `/tasks`
- **Request Body**:
```json
{
  "title": "Complete project",
  "description": "Finish the task management system"
}
```
- **Response**:
```json
{
  "success": true,
  "data": {
    "id": "123e4567-e89b-12d3-a456-426614174000",
    "title": "Complete project",
    "description": "Finish the task management system",
    "status": "Pending",
    "created_at": "2024-01-15T10:30:00Z",
    "updated_at": "2024-01-15T10:30:00Z"
  }
}
```

#### 2. Get All Tasks
- **GET** `/tasks?page=1&page_size=10&status=Pending`
- **Query Parameters**:
  - `page`: Page number (default: 1)
  - `page_size`: Items per page (default: 10, max: 100)
  - `status`: Filter by status (Pending, InProgress, Completed)
- **Response**:
```json
{
  "success": true,
  "data": [
    {
      "id": "123e4567-e89b-12d3-a456-426614174000",
      "title": "Complete project",
      "description": "Finish the task management system",
      "status": "Pending",
      "created_at": "2024-01-15T10:30:00Z",
      "updated_at": "2024-01-15T10:30:00Z"
    }
  ],
  "meta": {
    "page": 1,
    "page_size": 10,
    "total": 1,
    "total_pages": 1
  }
}
```

#### 3. Get Task by ID
- **GET** `/tasks/{id}`
- **Response**:
```json
{
  "success": true,
  "data": {
    "id": "123e4567-e89b-12d3-a456-426614174000",
    "title": "Complete project",
    "description": "Finish the task management system",
    "status": "Pending",
    "created_at": "2024-01-15T10:30:00Z",
    "updated_at": "2024-01-15T10:30:00Z"
  }
}
```

#### 4. Update Task
- **PUT** `/tasks/{id}`
- **Request Body** (all fields optional):
```json
{
  "title": "Updated title",
  "description": "Updated description",
  "status": "Completed"
}
```
- **Response**: Same as Get Task

#### 5. Delete Task
- **DELETE** `/tasks/{id}`
- **Response**:
```json
{
  "success": true,
  "data": {
    "message": "Task deleted successfully"
  }
}
```

#### 6. Health Check
- **GET** `/health`
- **Response**:
```json
{
  "status": "healthy"
}
```

### Error Responses

All error responses follow this format:
```json
{
  "success": false,
  "error": "Error message description"
}
```

## Microservices Concepts

### 1. Single Responsibility Principle
Each component has a clear, single responsibility:
- **Domain**: Business entities and rules
- **Repository**: Data access and persistence
- **Service**: Business logic and validation
- **Handler**: HTTP request/response handling
- **Middleware**: Cross-cutting concerns

### 2. API Design
- RESTful endpoints following HTTP conventions
- Consistent response format across all endpoints
- Proper HTTP status codes
- Versioned API (`/api/v1`)

### 3. Separation of Concerns
- Clean separation between layers
- Dependency injection for loose coupling
- Interface-based design for testability

### 4. Stateless Design
- No server-side session state
- Each request contains all necessary information
- Enables horizontal scaling

## Scaling Considerations

### Horizontal Scaling

1. **Load Balancing**: Multiple service instances behind a load balancer
2. **Database Scaling**: Replace in-memory storage with distributed database
3. **Caching**: Add Redis for frequently accessed data
4. **Container Orchestration**: Use Kubernetes for auto-scaling

### Vertical Scaling

1. **Resource Optimization**: Increase CPU/memory for existing instances
2. **Connection Pooling**: Optimize database connections
3. **Performance Monitoring**: Monitor and optimize bottlenecks

### Implementation Strategy

```bash
# Example Kubernetes deployment
kubectl apply -f k8s/
kubectl scale deployment task-service --replicas=3
```

## Inter-Service Communication

### Adding a User Service

When extending the system with a User Service, consider these communication patterns:

#### 1. Synchronous Communication

**REST API Calls**:
```go
// In Task Service - validate user exists
userClient := &http.Client{}
resp, err := userClient.Get(fmt.Sprintf("%s/users/%s", userServiceURL, userID))
```

**gRPC** (for better performance):
```go
// user_service.proto
service UserService {
  rpc GetUser(GetUserRequest) returns (User);
}

// In Task Service
user, err := userClient.GetUser(ctx, &pb.GetUserRequest{Id: userID})
```

#### 2. Asynchronous Communication

**Message Queues** (RabbitMQ, Apache Kafka):
```go
// Publish task events
publisher.Publish("task.created", TaskCreatedEvent{
    TaskID: task.ID,
    UserID: task.UserID,
})

// User Service subscribes to task events
subscriber.Subscribe("task.created", handleTaskCreated)
```

#### 3. Event-Driven Architecture

```go
// Domain events
type TaskCreatedEvent struct {
    TaskID    string    `json:"task_id"`
    UserID    string    `json:"user_id"`
    Timestamp time.Time `json:"timestamp"`
}

// Event bus for decoupled communication
eventBus.Publish(TaskCreatedEvent{...})
```

### Service Discovery

For production deployments:
- **Consul**: Service registry and health checking
- **Kubernetes**: Built-in service discovery
- **API Gateway**: Single entry point for all services

### Example Multi-Service Architecture

```yaml
# docker-compose.yml for multi-service setup
version: '3.8'
services:
  task-service:
    build: ./task-service
    ports: ["8080:8080"]
    environment:
      - USER_SERVICE_URL=http://user-service:8081
  
  user-service:
    build: ./user-service
    ports: ["8081:8081"]
    environment:
      - TASK_SERVICE_URL=http://task-service:8080
  
  message-queue:
    image: rabbitmq:3-management
    ports: ["5672:5672", "15672:15672"]
```

## Development

### Project Structure
```
task-management-system/
├── cmd/server/              # Application entry point
├── internal/
│   ├── domain/             # Business entities and interfaces
│   ├── repository/         # Data access layer
│   ├── service/           # Business logic
│   └── handler/           # HTTP handlers
├── pkg/
│   └── response/         # API response utilities
├── Dockerfile             # Container definition
├── docker-compose.yml     # Multi-container setup
├── go.mod                # Go module definition
└── README.md             # This file
```

### Adding New Features

1. Define domain entities in `internal/domain/`
2. Implement repository interface in `internal/repository/`
3. Add business logic in `internal/service/`
4. Create HTTP handlers in `internal/handler/`
5. Update routing in `cmd/server/main.go`

### Best Practices Implemented

- **Interface-based design** for testability and flexibility
- **Dependency injection** for loose coupling
- **Error handling** with proper HTTP status codes
- **Input validation** at service layer
- **Simple in-memory storage** with basic map operations
- **Standardized responses** for consistent API

## License

This project is created for demonstration purposes as part of the Alle Backend Assignment.
