package main

import (
    "fmt"
    "my-go-project/pkg"
)

func main() {
    sum := pkg.Add(5, 3)
    difference := pkg.Subtract(5, 3)

    fmt.Printf("Sum: %d\n", sum)
    fmt.Printf("Difference: %d\n", difference)
}