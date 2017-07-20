package main

import (
    "fmt"
)

type Integer int
type IntegerStruct struct {
    n int
}

func main() {
    v1 := Integer(1)
    v2 := IntegerStruct{2}

    f(v1.get())
    f(v2.get())
}

func (v Integer) get() int {
    return int(v)
}

func (v IntegerStruct) get() int {
    return v.n
}

func f(i int) {
    fmt.Printf("Called function f(%d)\n", i)
}