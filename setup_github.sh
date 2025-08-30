#!/bin/bash

# GitHub Repository Setup Script
# This script helps you initialize the Git repository and push to GitHub

set -e

echo "🚀 Setting up Git repository for Task Management System"

# Initialize git repository
git init

# Add all files
git add .

# Create initial commit
git commit -m "Initial commit: Task Management System microservice

- Implemented clean architecture with domain, service, repository layers
- Added CRUD operations for tasks with validation
- Implemented pagination and filtering for GET /tasks endpoint
- Added comprehensive middleware (logging, CORS, recovery)
- Included Docker support and docker-compose setup
- Added unit tests for service layer
- Created comprehensive documentation and API examples"

echo "✅ Git repository initialized with initial commit"

echo "📝 Next steps:"
echo "1. Create a new repository on GitHub named 'task-management-system'"
echo "2. Run the following commands:"
echo "   git remote add origin https://github.com/pvnptl/task-management-system.git"
echo "   git branch -M main"
echo "   git push -u origin main"

echo ""
echo "🔧 To test the application locally:"
echo "1. Install Go if not already installed"
echo "2. Run: make run"
echo "3. Or with Docker: make docker-run"

echo ""
echo "📚 Additional development commits you can make:"
echo "1. Add database integration (PostgreSQL/MongoDB)"
echo "2. Implement JWT authentication"
echo "3. Add more comprehensive tests"
echo "4. Add monitoring and metrics"
echo "5. Implement caching layer"
