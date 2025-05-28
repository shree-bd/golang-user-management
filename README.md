# Go MongoDB User Management API

A simple REST API built with Go and MongoDB for user management operations.

## Features

- Create, read, update, and delete users
- MongoDB integration
- Web frontend interface
- RESTful API endpoints

## Prerequisites

- Go 1.24+
- MongoDB

## Installation

1. Clone the repository
2. Install dependencies:
   ```bash
   go mod tidy
   ```
3. Start MongoDB service
4. Run the application:
   ```bash
   go run main.go
   ```
5. Open http://localhost:9000 in your browser

## API Endpoints

- `GET /api/users` - Get all users
- `GET /api/users/:id` - Get user by ID
- `POST /api/users` - Create new user
- `DELETE /api/users/:id` - Delete user

## Tech Stack

- Go
- MongoDB
- HTML/CSS/JavaScript
- httprouter
- MongoDB Go Driver 