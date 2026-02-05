# Employee Management Backend API

This is a backend API for managing employees, built with Go, Gin, and PostgreSQL. It includes features for user authentication (JWT), CRUD operations for employees, salary calculation based on country, and aggregated salary metrics.

## Table of Contents
- [Features](#features)
- [Project Structure](#project-structure)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Local Development](#local-development)
  - [Running with Docker](#running-with-docker)
- [API Overview](#api-overview)
  - [Authentication Flow](#authentication-flow)
  - [Endpoints](#endpoints)
- [Database Migration Strategy](#database-migration-strategy)
- [Design Decisions and Assumptions](#design-decisions-and-assumptions)
- [Entity Relationship Diagram (ERD)](#entity-relationship-diagram-erd)

## Features
- User Registration and Login with JWT authentication
- CRUD operations for Employee resources (Full Name, Job Title, Country, Gross Salary)
- Net Salary calculation based on country-specific tax deductions
- Aggregated Salary Metrics (min, max, average salary by country; average salary by job title)
- Externalized configuration using YAML and Viper
- Database persistence with PostgreSQL and GORM
- Dockerization for easy deployment and local development

## Project Structure

```
employee-management/
├── cmd/
│   └── main.go           # Entry point
├── config/               # Application configuration
│   └── config.go         # Loads config.yaml using Viper
├── routers/              # Main API routing definitions
│   └── router.go         # Initializes Gin and registers all module routes
├── internal/
│   ├── users/            # User management module (handlers, services, repository, routes)
│   │   ├── handler.go    # HTTP handlers for user operations (register, login)
│   │   ├── service.go    # Business logic for users (hashing, auth)
│   │   ├── repository.go # Data access logic for users
│   │   └── router.go     # Module-specific user routes
│   ├── employee/         # Employee management module (handlers, services, repository, routes)
│   │   ├── handler.go    # HTTP handlers for employee CRUD, salary calculation, metrics
│   │   ├── service.go    # Business logic for employees (CRUD, salary logic)
│   │   ├── repository.go # Data access logic for employees
│   │   └── router.go     # Module-specific employee routes
│   ├── jwt/              # JWT token generation and parsing utilities
│   │   └── jwt.go
│   ├── db/               # Database connection setup
│   │   └── db.go         # GORM database initialization
│   ├── middleware/       # General purpose middleware
│   │   └── auth_middleware.go # JWT authentication middleware
│   └── models/           # Centralized database models
│       ├── users.go      # User model definition
│       └── employee.go   # Employee model definition
├── migrations/           # Database migration files
│   ├── 000001_initial_schema.up.sql
│   └── 000001_initial_schema.down.sql
├── Dockerfile            # Docker image definition
├── docker-compose.yml    # Docker orchestration for local development
├── config.yaml           # Configuration file
└── README.md             # Project documentation
```

## Getting Started

### Prerequisites
- Go (1.22+)
- PostgreSQL
- Docker and Docker Compose (optional, for containerized setup)

### Local Development

1.  **Clone the repository:**
    ```bash
    git clone https://github.com/venkythirupari/employee-management.git
    cd employee-management
    ```

2.  **Install Go Modules:**
    ```bash
    go mod tidy
    ```

3.  **Configure Database:**
    Ensure a PostgreSQL instance is running and accessible. Update `config.yaml` with your database connection details.

    **`config.yaml` example:**
    ```yaml
    server:
      port: "8080"
    database:
      host: "localhost"
      port: "5432"
      user: "postgres"
      password: "password"
      dbname: "employee_db"
    jwt:
      secret: "your_secret_key"
    ```
    Replace `your_secret_key` with a strong, unique secret.

4.  **Run Database Migrations:**
    Install the `migrate` CLI tool:
    ```bash
    go install github.com/golang-migrate/migrate/v4/cmd/migrate@latest
    ```
    Run migrations:
    ```bash
    migrate -path migrations -database "postgres://postgres:password@localhost:5432/employee_db?sslmode=disable" up
    ```
    Adjust the database URL as per your `config.yaml` settings.

5.  **Run the Application:**
    ```bash
    go run cmd/main.go
    ```
    The API will be available at `http://localhost:8080`.

### Running with Docker

1.  **Build and Run with Docker Compose:**
    ```bash
    docker-compose up --build
    ```
    This will build the Go application image, start a PostgreSQL container, and run the application. The API will be available at `http://localhost:8080`.

2.  **Run Database Migrations (inside Docker):**
    You can run migrations against the Dockerized PostgreSQL. First, ensure the `db` service is up. Then, execute the `migrate` command using your `app` container:
    ```bash
    docker-compose exec app migrate -path migrations -database "postgres://postgres:password@db:5432/employee_db?sslmode=disable" up
    ```
    Note that the database host is `db` (the service name in `docker-compose.yml`) when connecting from within another container.

## API Overview

The API supports user authentication, employee management, salary calculation, and salary metrics. All non-authentication endpoints require a valid JWT in the `Authorization` header (`Bearer <token>`).

### Authentication Flow
1.  **Register:** `POST /api/v1/users/register` with `username` and `password`.
2.  **Login:** `POST /api/v1/users/login` with `username` and `password` to receive a JWT.
3.  **Access Protected Endpoints:** Include the received JWT in the `Authorization: Bearer <token>` header for all other API calls.

### Endpoints
*   **Users:**
    *   `POST /api/v1/users/register` - Register a new user.
    *   `POST /api/v1/users/login` - Authenticate a user and get a JWT.
*   **Employees:** (Requires JWT)
    *   `POST /api/v1/employees/` - Create a new employee.
    *   `GET /api/v1/employees/` - Get all employees.
    *   `GET /api/v1/employees/{id}` - Get an employee by ID.
    *   `PUT /api/v1/employees/{id}` - Update an employee.
    *   `DELETE /api/v1/employees/{id}` - Delete an employee.
    *   `GET /api/v1/employees/{id}/salary/net` - Calculate net salary for an employee.
*   **Metrics:** (Requires JWT)
    *   `GET /api/v1/metrics/salary/country?country={country_name}` - Get salary metrics (min, max, avg) for a given country.
    *   `GET /api/v1/metrics/salary/job-title?job_title={job_title}` - Get average salary for a given job title.

## Database Migration Strategy
Database schema changes are managed using `golang-migrate/migrate`. Migration files are located in the `migrations/` directory. Each change set consists of an `up.sql` file (for applying changes) and a `down.sql` file (for reverting changes). This ensures versioned, reproducible schema evolution.

## Design Decisions and Assumptions
-   **Go Modules:** Used for dependency management.
-   **Gin Framework:** Chosen for its performance and ease of use in building RESTful APIs.
-   **PostgreSQL & GORM:** PostgreSQL is a robust relational database, and GORM provides an expressive ORM layer.
-   **Viper:** Used for externalizing configuration via `config.yaml`, providing flexibility and environment-specific settings.
-   **JWT Authentication:** A standard and secure method for stateless API authentication.
-   **Project Structure:** A layered architecture (`handlers -> services -> repositories -> database`) is used to promote separation of concerns, testability, and maintainability.
-   **Error Handling:** Basic error handling is implemented. For production, more detailed custom error types and centralized error responses might be beneficial.
-   **API Versioning:** `/api/v1` prefix is used for basic API versioning.
-   **Salary Deduction Rules:** Hardcoded in the `employee/service.go`. In a real-world scenario, these might be configurable or fetched from a dedicated service.

## Entity Relationship Diagram (ERD)

```mermaid
erDiagram
    users {
        SERIAL id PK
        VARCHAR unique_username
        VARCHAR password_hash
        TIMESTAMP created_at
        TIMESTAMP updated_at
    }

    employees {
        SERIAL id PK
        VARCHAR full_name
        VARCHAR job_title
        VARCHAR country
        NUMERIC gross_salary
        TIMESTAMP created_at
        TIMESTAMP updated_at
    }
```
