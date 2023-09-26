# CRUD Application with React.js, Vite, Axios, and Golang

This is a simple CRUD (Create, Read, Update, Delete) application built using React.js for the frontend, Vite for the development environment, Axios for API requests, and Golang for the backend. This application demonstrates how to perform basic CRUD operations with a modern web stack.

## Features

- Create, Read, Update, and Delete operations on a resource (e.g., Todos, Users, etc.).
- Separation of frontend and backend code for modularity.
- RESTful API design with Golang.
- Easy development setup with Vite.
- Axios for handling HTTP requests.

## Prerequisites

Before you begin, ensure you have met the following requirements:

- Node.js and npm installed (for the frontend).
- Golang installed (for the backend).
- Git installed.
- Clone this repository:

```bash
git clone https://github.com/autumnleaf-ra/go-crud-react.git
```

# Frontend (React.js with Vite)
```bash
cd client
npm install
```
## Usage
```bash
npm run dev
```
The Front-End will start at http://localhost:5173.

# Backend (Golang)

## Installation
```bash
go mod tidy
```

## Setting Port & Database
```
1. Open '.env' for setting database
2. Config.js API_URL "http://localhost:port/api"
3. Bootstrap/app.go - Setting listenPort API Backend
```
## Usage
```bash
go run main.go
```
The API server will start at http://localhost:4000.

## API Endpoints
- GET /api/users
- GET /api/users/:id
- POST /api/users
- PATCH /api/users/:id
- DELETE /api/users/:id
