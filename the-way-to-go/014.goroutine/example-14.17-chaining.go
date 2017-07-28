package main

import (
    "flag"
    "fmt"
    "runtime"
    "encoding/json"
)

var ngoroutine = flag.Int("n", 10000, "how many goroutines")

func f(left, right chan int) {
    left <- 1 + <-right
}

func main() {
    flag.Parse()
    leftmost := make(chan int)
    var left, right chan int = nil, leftmost
    for i := 0; i < *ngoroutine; i ++ {
        left, right = right, make(chan int)
        go f(left, right)
    }

    /* dump running goroutine info */
    numGoroutine := runtime.NumGoroutine()
    profiles := make([]runtime.StackRecord, numGoroutine)
    _, _ = runtime.GoroutineProfile(profiles)
    profileJson, _ := json.MarshalIndent(profiles, "", "    ")
    fmt.Printf("%d\n", numGoroutine)
    fmt.Printf("%s\n", profileJson)

    right <- 0
    x := <-leftmost
    fmt.Println(x)
}