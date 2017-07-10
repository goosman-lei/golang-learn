package main

import (
    "fmt"
)

func main() {
    str := "Good Bye"
    var p *string = &str

    *p = "Bye-bye"

    fmt.Printf("Here is the pointer p: %p\n", p)
    fmt.Printf("Here is the string *p: %s\n", *p)
    fmt.Printf("Here is the string str: %s\n", str)
}