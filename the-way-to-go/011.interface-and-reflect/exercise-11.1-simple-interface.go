package main

import (
    "fmt"
)

type Simpler interface {
    Get() int
    Set(int)
}

func demo1(s Simpler, v int) int {
    s.Set(v)
    return s.Get()
}

func demo2(s Simpler, v int) (n int) {
    n = s.Get()
    s.Set(v)
    return
}

type DemoSimpler struct {
    n int
}
func (this *DemoSimpler) Get() int {
    return this.n
}
func (this *DemoSimpler) Set(v int) {
    this.n = v
}

func main() {
    s := Simpler(&DemoSimpler{3})

    fmt.Printf("DemoSimpler Value: %+v\n", s)
    fmt.Printf("demo2(s, 9): %+v\n", demo2(s, 9))
    fmt.Printf("demo2(s, 20): %+v\n", demo2(s, 9))
    fmt.Printf("demo1(s, 8): %+v\n", demo1(s, 8))
    fmt.Printf("demo1(s, 18): %+v\n", demo1(s, 8))
}