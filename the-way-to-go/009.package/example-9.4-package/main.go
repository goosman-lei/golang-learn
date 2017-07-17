package main

import (
    "fmt"
    "./pack"
)

func main() {
    var test1 string
    test1 = pack.ReturnStr()

    fmt.Printf("ReturnStr from package pack: %s\n", test1)
    fmt.Printf("Integer from package pack: %d\n", pack.PackInt)
    //fmt.Printf("Float from package pack: %f\n", pack.packFloat)
}