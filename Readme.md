# Todo API

A simple TODO application with basic CRUD operations using Go and an in-memory store.

## Setup Instructions

### Prerequisites

- Go 1.18 or higher

### Running the Application

1. Clone the repository:
    ```bash
    git clone https://github.com/Anil-S-ML/todo.git
    ```

2. Navigate to the project directory:
    ```bash
    cd todo
    ```

3. Install dependencies (if any):
    ```bash
    go mod tidy
    ```

4. Run the application:
    ```bash
    go run cmd/server/main.go or go run cmd/cli/main.go for CLI Apllication
    ```

The server will start at `http://localhost:8080`.

---

## API Endpoints

### Health Check

- **GET /health**: Returns "OK" if the server is up.

### Todos

- **GET /todos**: Fetch all todo items.
  
- **POST /todos**: Create a new todo item. Example body:
    ```json
    {
        "title": "Finish Go project"
    }
    ```

- **GET /todos/{id}**: Fetch a todo by its ID.

- **PUT /todos/{id}/complete**: Mark a todo item as complete.

- **DELETE /todos/{id}**: Delete a todo item by its ID.

### Authentication

The API uses Basic Authentication for restricted endpoints.

---

## Tests

To run the tests:

```bash
go test ./...
