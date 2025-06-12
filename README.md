# TODO App

A modern TODO application built with Go (Gin) and vanilla JavaScript, featuring task management with categories and a clean web interface.

## Features

- Create, read, update, and delete tasks
- Organize tasks by categories
- Task status tracking (completed/pending)
- Task statistics and analytics
- Clean and responsive web interface
- RESTful API endpoints
- SQLite database with GORM

## Tech Stack

- **Backend**: Go with Gin framework
- **Database**: SQLite with GORM ORM
- **Frontend**: Vanilla JavaScript, HTML5, CSS3
- **Development**: Air for hot reloading
- **Containerization**: Docker & Docker Compose

## Quick Start

### Prerequisites

- Go 1.21.4 or higher
- Docker (optional)

### Running Locally

1. Clone the repository:
```bash
git clone <repository-url>
cd hello-claude-code
```

2. Install dependencies:
```bash
go mod tidy
```

3. Run the application:
```bash
go run main.go
```

4. Open your browser and navigate to `http://localhost:8080`

### Using Docker

1. Build and run with Docker Compose:
```bash
docker-compose up --build
```

2. Access the application at `http://localhost:8080`

### Development Mode

For development with hot reloading:

```bash
# Install Air (if not already installed)
go install github.com/cosmtrek/air@latest

# Run in development mode
air
```

Or using Docker:
```bash
docker-compose -f docker-compose.yml up dev
```

## API Endpoints

### Tasks
- `GET /api/v1/tasks` - Get all tasks
- `GET /api/v1/tasks/stats` - Get task statistics
- `GET /api/v1/tasks/:id` - Get a specific task
- `POST /api/v1/tasks` - Create a new task
- `PUT /api/v1/tasks/:id` - Update a task
- `DELETE /api/v1/tasks/:id` - Delete a task
- `PATCH /api/v1/tasks/:id/toggle` - Toggle task completion status

### Categories
- `GET /api/v1/categories` - Get all categories
- `GET /api/v1/categories/:id` - Get a specific category
- `POST /api/v1/categories` - Create a new category
- `PUT /api/v1/categories/:id` - Update a category
- `DELETE /api/v1/categories/:id` - Delete a category

## Project Structure

```
.
├── database/           # Database connection and models
├── handlers/           # HTTP request handlers
├── models/            # Data models
├── static/            # Static assets (CSS, JS)
├── templates/         # HTML templates
├── docs/              # Project documentation
├── Dockerfile         # Production Docker image
├── Dockerfile.dev     # Development Docker image
├── docker-compose.yml # Docker Compose configuration
├── .air.toml         # Air configuration for hot reloading
├── main.go           # Application entry point
└── README.md         # This file
```

## Database Schema

The application uses two main entities:

- **Tasks**: ID, Title, Description, Completed status, Category ID, timestamps
- **Categories**: ID, Name, Color, timestamps

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Test your changes
5. Submit a pull request

## License

This project is open source and available under the MIT License.