package main

import (
    "fmt"
    "runtime"
)

func main() {
    where := func() {
        _, file, line, _ := runtime.Caller(1)
        fmt.Printf("%s:%d\n", file, line)
    }

    where()
    where()
    where()
}