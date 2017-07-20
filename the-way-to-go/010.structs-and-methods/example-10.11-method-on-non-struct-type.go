package main

import "fmt"

type IntVector []int

func (this IntVector) Sum() (s int) {
    for _, x := range this {
        s += x
    }
    return
}

func main() {
    fmt.Println(IntVector{1, 2, 3, 4}.Sum())
}