package main

import (
    "fmt"
    "time"
)

func main() {
    ch := make(chan int)
    syCh := make(chan string)

    go producer(ch, syCh)
    go consumer(ch, syCh)

    <-syCh
    <-syCh
}

func producer(ch chan int, syCh chan string) {
    for i := 0; i < 100; i += 10 {
        ch <- i
        time.Sleep(1e8)
    }
    syCh <- "done"
    fmt.Printf("producer done\n")
    close(ch)
}
func consumer(ch chan int, syCh chan string) {
    for {
        if n, ok := <-ch; ok {
            fmt.Printf("received value: %d\n", n)
        } else {
            fmt.Printf("receive channel closed\n")
            break;
        }
    }
    syCh <- "done"
}