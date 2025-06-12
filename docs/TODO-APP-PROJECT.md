# TODO App Project

A simple TODO application built with Go and Gin framework.

## Features

- RESTful API for TODO management
- JSON responses
- Lightweight and fast with Gin framework

## Project Structure

```
hello-claude-code/
├── main.go          # Main application entry point
├── go.mod           # Go module dependencies
├── go.sum           # Dependency checksums
└── docs/
    └── TODO-APP-PROJECT.md # Project documentation
```

## Getting Started

### Prerequisites

- Go 1.21.4 or later
- Git

### Installation

1. Clone the repository:
```bash
git clone git@github.com:iotassss/hello-claude-code.git
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

The server will start on `http://localhost:8080`

## API Endpoints

### Current Endpoints

- `GET /` - Returns a Hello World message

### Planned TODO Endpoints

- `GET /todos` - Get all TODO items
- `POST /todos` - Create a new TODO item
- `GET /todos/:id` - Get a specific TODO item
- `PUT /todos/:id` - Update a TODO item
- `DELETE /todos/:id` - Delete a TODO item

## Development

### Running in Development Mode

```bash
go run main.go
```

### Building for Production

```bash
go build -o todo-app main.go
./todo-app
```

## TODO Features to Implement

- [ ] Database integration (SQLite/PostgreSQL)
- [ ] TODO CRUD operations
- [ ] User authentication
- [ ] Due dates and priorities
- [ ] Categories/tags
- [ ] Search functionality
- [ ] RESTful API documentation
- [ ] Frontend interface
- [ ] Docker containerization
- [ ] Unit tests
- [ ] CI/CD pipeline

## Dependencies

- [Gin](https://github.com/gin-gonic/gin) - HTTP web framework

## Contributing

1. Fork the repository
2. Create a feature branch
3. Commit your changes
4. Push to the branch
5. Create a Pull Request

## License

This project is open source and available under the MIT License.