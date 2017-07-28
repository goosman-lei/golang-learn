package main

import (
    "fmt"
    "time"
)

func main() {
    suck(pump())

    time.Sleep(1e9)
}

func pump() chan int {
    ch := make(chan int)
    go func() {
        for i := 0; ; i ++ {
            ch <- i
            time.Sleep(1e8)
        }
    }()
    return ch
}

func suck(ch chan int) {
    go func() {
        for i, v := range ch {
            fmt.Printf("i: %s v: %s\n", i, v)
        }
    }()
}