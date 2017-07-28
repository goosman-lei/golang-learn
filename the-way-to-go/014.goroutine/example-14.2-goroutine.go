package main

import (
    "fmt"
    "time"
)

func main() {
    ch := make(chan string, 3)

    go sendData(ch)
    go getData(ch)

    fmt.Printf("[%s]In main before sleep\n", time.Now())
    time.Sleep(1e9)
    fmt.Printf("[%s]In main after sleep\n", time.Now())
}

func sendData(ch chan string) {
    ch <- "Washington"
    fmt.Printf("[%s]send Washington\n", time.Now())
    ch <- "Tripoli"
    fmt.Printf("[%s]send Tripoli\n", time.Now())
    ch <- "London"
    fmt.Printf("[%s]send London\n", time.Now())
    ch <- "Beijing"
    fmt.Printf("[%s]send Beijing\n", time.Now())
    ch <- "Tokio"
    fmt.Printf("[%s]send Tokio\n", time.Now())
    close(ch)
}

func getData(ch chan string) {
    for {
        time.Sleep(1e6)
        if input, ok := <-ch; ok {
            fmt.Printf("[%s]receive message from channel: %s\n", time.Now(), input)
        } else {
            fmt.Printf("[%s]channel is closed\n", time.Now())
            break
        }
    }
}