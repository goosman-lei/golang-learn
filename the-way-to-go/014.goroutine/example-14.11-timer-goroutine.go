package main

import (
    "fmt"
    "time"
)

func main() {
    tick := time.Tick(1e8)
    boom := time.After(5e8)

    for {
        select {
            case t := <-tick:
                fmt.Printf("Tick value: %v\n", t)
            case b := <-boom:
                fmt.Printf("Boom value: %v\n", b)
                return
            default:
                fmt.Println("Unknown      .")
                time.Sleep(5e7)
        }
    }
}