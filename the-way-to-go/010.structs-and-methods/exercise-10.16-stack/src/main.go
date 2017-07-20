package main

import (
    "fmt"
    "stack"
)

func main() {
    s := stack.NewStack(5)
    fmt.Printf("Test for push:\n")
    fmt.Printf("\ts1.Push(1): %t\n", s.Push(1))
    fmt.Printf("\ts1.Push(2): %t\n", s.Push(2))
    fmt.Printf("\ts1.Push(3): %t\n", s.Push(3))
    fmt.Printf("\ts1.Push(4): %t\n", s.Push(4))
    fmt.Printf("\ts1.Push(5): %t\n", s.Push(5))
    fmt.Printf("\ts1.Push(6): %t\n", s.Push(6))
    fmt.Printf("\ts1.Push(7): %t\n", s.Push(7))
    fmt.Printf("%s\n", s)

    fmt.Printf("Test for pop:\n")
    fmt.Printf("\ts1.Pop(): %d\n", s.Pop())
    fmt.Printf("\ts1.Pop(): %d\n", s.Pop())
    fmt.Printf("\ts1.Pop(): %d\n", s.Pop())
    fmt.Printf("\ts1.Pop(): %d\n", s.Pop())
    fmt.Printf("\ts1.Pop(): %d\n", s.Pop())
    fmt.Printf("\ts1.Pop(): %d\n", s.Pop())
    fmt.Printf("\ts1.Pop(): %d\n", s.Pop())
    fmt.Printf("%s\n", s)

    fmt.Printf("Test for push and pop:\n")
    fmt.Printf("\ts1.Push(1): %t\n", s.Push(1))
    fmt.Printf("\ts1.Push(2): %t\n", s.Push(2))
    fmt.Printf("\ts1.Push(3): %t\n", s.Push(3))
    fmt.Printf("\ts1.Pop(): %d\n", s.Pop())
    fmt.Printf("\ts1.Pop(): %d\n", s.Pop())
    fmt.Printf("\ts1.Push(4): %t\n", s.Push(4))
    fmt.Printf("\ts1.Pop(): %d\n", s.Pop())
    fmt.Printf("\ts1.Pop(): %d\n", s.Pop())
    fmt.Printf("\ts1.Push(5): %t\n", s.Push(5))
    fmt.Printf("\ts1.Push(6): %t\n", s.Push(6))
    fmt.Printf("\ts1.Pop(): %d\n", s.Pop())
    fmt.Printf("\ts1.Push(7): %t\n", s.Push(7))
    fmt.Printf("%s\n", s)
}
