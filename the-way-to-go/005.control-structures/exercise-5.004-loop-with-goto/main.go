package main

import "fmt"

func main() {
    i := 0
loop:
    if i < 10 {
        i ++
        fmt.Println(i)
        goto loop
    }
}