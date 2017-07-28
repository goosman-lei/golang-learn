package main

import (
    "fmt"
    "time"
)

type Empty interface{}
type Semaphore chan Empty

func NewSemaphore(n int) Semaphore {
    return make(Semaphore, n)
}

func (s Semaphore) P(n int) {
    for n > 0 {
        s <- new(Empty)
        time.Sleep(1e1)
        n --
    }
}
func (s Semaphore) V(n int) {
    for n > 0 {
        <-s
        n --
    }
}

func (s Semaphore) Lock() {
    s.P(1)
}
func (s Semaphore) Unlock() {
    s.V(1)
}

// here have some problem:
//      when multi goroutine concurrence wait,
//      everyone locked some resource,
//      but nobody requirement was satisified
func (s Semaphore) Wait(n int) {
    s.P(n)
}
func (s Semaphore) Signal() {
    s.V(1)
}

func main() {
    s := NewSemaphore(10)

    go func() {
        s.Wait(5)
        fmt.Printf("first goroutine waited success\n")
    }()
    go func() {
        s.Wait(5)
        fmt.Printf("second goroutine waited success\n")
    }()
    go func() {
        s.Wait(5)
        fmt.Printf("third goroutine waited success\n")
    }()
    time.Sleep(1e9)
}