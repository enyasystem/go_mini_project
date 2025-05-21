package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
    "my-go-project/pkg"
)

func main() {
    todoList := pkg.NewTodoList()
    reader := bufio.NewReader(os.Stdin)

    for {
        fmt.Println("\nTodo List App")
        fmt.Println("1. Add Todo")
        fmt.Println("2. List Todos")
        fmt.Println("3. Remove Todo")
        fmt.Println("4. Mark Todo as Done")
        fmt.Println("5. Edit Todo")
        fmt.Println("6. Clear Completed Todos")
        fmt.Println("7. Exit")
        fmt.Print("Choose an option: ")
        input, _ := reader.ReadString('\n')
        input = strings.TrimSpace(input)

        switch input {
        case "1":
            fmt.Print("Enter todo title: ")
            title, _ := reader.ReadString('\n')
            title = strings.TrimSpace(title)
            todo := todoList.AddTodo(title)
            fmt.Printf("Added: [%d] %s\n", todo.ID, todo.Title)
        case "2":
            todos := todoList.ListTodos()
            if len(todos) == 0 {
                fmt.Println("No todos yet.")
            } else {
                fmt.Println("Todos:")
                for _, t := range todos {
                    status := " "
                    if t.Done {
                        status = "x"
                    }
                    fmt.Printf("[%d] [%s] %s\n", t.ID, status, t.Title)
                }
            }
        case "3":
            fmt.Print("Enter todo ID to remove: ")
            idStr, _ := reader.ReadString('\n')
            idStr = strings.TrimSpace(idStr)
            id, err := strconv.Atoi(idStr)
            if err != nil {
                fmt.Println("Invalid ID.")
                continue
            }
            if todoList.RemoveTodo(id) {
                fmt.Println("Todo removed.")
            } else {
                fmt.Println("Todo not found.")
            }
        case "4":
            fmt.Print("Enter todo ID to mark as done: ")
            idStr, _ := reader.ReadString('\n')
            idStr = strings.TrimSpace(idStr)
            id, err := strconv.Atoi(idStr)
            if err != nil {
                fmt.Println("Invalid ID.")
                continue
            }
            if todoList.MarkDone(id) {
                fmt.Println("Todo marked as done.")
            } else {
                fmt.Println("Todo not found.")
            }
        case "5":
            fmt.Print("Enter todo ID to edit: ")
            idStr, _ := reader.ReadString('\n')
            idStr = strings.TrimSpace(idStr)
            id, err := strconv.Atoi(idStr)
            if err != nil {
                fmt.Println("Invalid ID.")
                continue
            }
            fmt.Print("Enter new title: ")
            newTitle, _ := reader.ReadString('\n')
            newTitle = strings.TrimSpace(newTitle)
            if todoList.EditTodo(id, newTitle) {
                fmt.Println("Todo updated.")
            } else {
                fmt.Println("Todo not found.")
            }
        case "6":
            removed := todoList.ClearCompleted()
            fmt.Printf("%d completed todos cleared.\n", removed)
        case "7":
            fmt.Println("Goodbye!")
            return
        default:
            fmt.Println("Invalid option.")
        }
    }
}
