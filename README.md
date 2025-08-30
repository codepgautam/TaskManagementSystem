# Task Management System

A microservice-based task management system built with Go, demonstrating clean architecture principles and microservices design patterns.

## Table of Contents

- [Overview](#overview)
- [Architecture](#architecture)
- [API Documentation](#api-documentation)

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