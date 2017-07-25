package main

import (
    "fmt"
)

func main() {
    recoverFromDivideByZero()
    fmt.Printf("after call recoverFromDivideByZero()\n")
}

func recoverFromDivideByZero() {
    defer func() {
        r := recover()
        fmt.Printf("recover: %s\n", r)
    }()
    a := 1
    b := 0
    c := a / b
    fmt.Printf("%d", c)
}