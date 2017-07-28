package main

import (
    "fmt"
)

func generate() chan int {
    ch := make(chan int)
    go func() {
        for i := 2; i < 200 ; i ++ {
            ch <- i
        }
        fmt.Printf("\t\tgenerate output chan closed\n")
        close(ch)
    }()
    return ch
}

func filter(in <-chan int, prime int) chan int {
    out := make(chan int)
    go func() {
        for {
            if i, ok := <-in; ok {
                if i % prime != 0 {
                    out <- i
                }
            } else {
                fmt.Printf("\t\tfilter[%d] output chan closed\n", prime)
                close(out)
                break
            }
        }
    }()
    return out
}

func sieve() chan int {
    out := make(chan int)
    go func() {
        ch := generate()
        for {
            if prime, ok := <-ch; ok {
                ch = filter(ch, prime)
                out <- prime
            } else {
                fmt.Printf("\t\tsieve output chan closed\n")
                close(out)
                break
            }
        }
    }()
    return out
}

func main() {
    primes := sieve()
    for {
        if prime, ok := <-primes; ok {
            fmt.Printf("%d\n", prime)
        } else {
            break
        }
    }
}