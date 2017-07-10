package main

import (
    "fmt"
    "time"
    "golang.org/x/sys/unix"
)

func main() {
    var startRusage unix.Rusage
    var endRusage unix.Rusage

    start := time.Now()
    unix.Getrusage(0, &startRusage)
    f(40)
    unix.Getrusage(0, &endRusage)
    end := time.Now()

    delta := end.Sub(start)
    fmt.Printf("%s\n", delta)

    fmt.Printf("start rusage: %+v\n", startRusage)
    fmt.Printf("start rusage: %+v\n", endRusage)
}

func f(n int) int {
    if (n <= 1) {
        return 1
    }

    return f(n - 1) + f(n - 2)
}