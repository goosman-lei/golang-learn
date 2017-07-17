package main

import (
    "fmt"
    "sync"
)

type Info struct {
    mu sync.Mutex
    Str string
}

func Update(info *Info) {
    info.mu.Lock()
    info.Str = "new value"
    info.mu.Unlock()
}

func main() {
    var once sync.Once

    for i := 0; i < 10; i ++ {
        fmt.Printf("index: %d\n", i)
        once.Do(func() {
            fmt.Printf("once print index: %d\n", i)
        })
    }
}