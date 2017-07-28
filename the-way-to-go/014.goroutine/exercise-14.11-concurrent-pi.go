package main

import (
    "runtime"
    "fmt"
)

func main() {
    runtime.GOMAXPROCS(1)
    n := 10000000
    fmt.Printf("PI in %d: %0.20f\n", n, schedule(n))
}

func schedule(n int) float64 {
    iCh := generator(n)

    oCh1 := calculator(iCh)
    /*
    oCh2 := calculator(iCh)
    oCh3 := calculator(iCh)
    oCh4 := calculator(iCh)
    */

    pi := 0.0
    FORLOOP:
    for {
        select {
            case v, ok := <-oCh1:
                if !ok {
                    break FORLOOP
                }
                pi += v
                //fmt.Printf("schedule receive from channel 1: %f: PI: %f\n", v, pi)
            /*
            case v, ok := <-oCh2:
                if !ok {
                    break FORLOOP
                }
                pi += v
                //fmt.Printf("schedule receive from channel 2: %f: PI: %f\n", v, pi)
            case v, ok := <-oCh3:
                if !ok {
                    break FORLOOP
                }
                pi += v
                //fmt.Printf("schedule receive from channel 3: %f: PI: %f\n", v, pi)
            case v, ok := <-oCh4:
                if !ok {
                    break FORLOOP
                }
                pi += v
                //fmt.Printf("schedule receive from channel 4: %f: PI: %f\n", v, pi)
            */
        }
    }
    return pi
}
func generator(n int) chan int {
    ch := make(chan int, 256)

    go func() {
        defer close(ch)
        for i, x := 0, 1; i < n; i ++ {
            //fmt.Printf("generator: %d\n", x)
            ch <- x
            x += 2
        }
    }()

    return ch
}
func calculator(iCh chan int) chan float64 {
    oCh := make(chan float64, 64)

    go func() {
        defer close(oCh)
        for n := range iCh {
            //fmt.Printf("calculator receive: %d\n", n)
            oCh <- float64((n + 1) / 2 % 2 * 8 - 4) / float64(n)
        }
    }()

    return oCh
}