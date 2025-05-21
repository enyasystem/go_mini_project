package utils

func Add(a int, b int) int {
    return a + b
}

func Subtract(a int, b int) int {
    return a - b
}

// Todo represents a single todo item
type Todo struct {
    ID    int
    Title string
    Done  bool
}

// TodoList manages a list of todos
type TodoList struct {
    todos []Todo
    nextID int
}

// NewTodoList creates a new TodoList
func NewTodoList() *TodoList {
    return &TodoList{todos: []Todo{}, nextID: 1}
}

// AddTodo adds a new todo item
func (tl *TodoList) AddTodo(title string) Todo {
    todo := Todo{ID: tl.nextID, Title: title, Done: false}
    tl.todos = append(tl.todos, todo)
    tl.nextID++
    return todo
}

// ListTodos returns all todo items
func (tl *TodoList) ListTodos() []Todo {
    return tl.todos
}

// RemoveTodo removes a todo by ID
func (tl *TodoList) RemoveTodo(id int) bool {
    for i, t := range tl.todos {
        if t.ID == id {
            tl.todos = append(tl.todos[:i], tl.todos[i+1:]...)
            return true
        }
    }
    return false
}

// MarkDone marks a todo as done by ID
func (tl *TodoList) MarkDone(id int) bool {
    for i, t := range tl.todos {
        if t.ID == id {
            tl.todos[i].Done = true
            return true
        }
    }
    return false
}

// Edit a todo's title by ID
func (tl *TodoList) EditTodo(id int, newTitle string) bool {
    for i, t := range tl.todos {
        if t.ID == id {
            tl.todos[i].Title = newTitle
            return true
        }
    }
    return false
}

// Clear all completed todos
func (tl *TodoList) ClearCompleted() int {
    var active []Todo
    removed := 0
    for _, t := range tl.todos {
        if !t.Done {
            active = append(active, t)
        } else {
            removed++
        }
    }
    tl.todos = active
    return removed
}
