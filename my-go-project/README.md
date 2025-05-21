# Todo List Go App

This is a simple command-line Todo List application written in Go. It allows you to add, list, remove, edit, and mark todos as done, as well as clear all completed todos.

## Project Structure

```
my-go-project
├── cmd
│   └── main.go        # Entry point of the application (CLI)
├── pkg
│   └── utils.go       # Todo list logic and utility functions
├── go.mod             # Module definition file
└── README.md          # Project documentation
```

## Getting Started

To build and run the application, follow these steps:

1. **Clone the repository:**
   ```
   git clone <repository-url>
   cd my-go-project
   ```

2. **Install dependencies:**
   ```
   go mod tidy
   ```

3. **Build the application:**
   ```
   go build -o my-go-app ./cmd
   ```

4. **Run the application:**
   ```
   ./my-go-app
   ```

## Features

- Add a new todo
- List all todos
- Remove a todo by ID
- Mark a todo as done
- Edit a todo's title
- Clear all completed todos

## Usage

When you run the app, you'll see a menu:

```
Todo List App
1. Add Todo
2. List Todos
3. Remove Todo
4. Mark Todo as Done
5. Edit Todo
6. Clear Completed Todos
7. Exit
```

Follow the prompts to manage your todo list. Todos are stored in memory for the current session.

## Contributing

Feel free to submit issues or pull requests to improve the project.
