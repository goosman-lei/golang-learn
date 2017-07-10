package main

import "fmt"

func main() {
    counter1 := genCounter(1, 1)
    counter2 := genCounter(100, 100)

    fmt.Printf("counter1(): %d\n", counter1())
    fmt.Printf("counter1(): %d\n", counter1())
    fmt.Printf("counter1(): %d\n", counter1())
    fmt.Println()

    fmt.Printf("counter2(): %d\n", counter2())
    fmt.Printf("counter2(): %d\n", counter2())
    fmt.Printf("counter2(): %d\n", counter2())
    fmt.Println()
}

func genCounter(start, step int) func() int {
    curVal := start
    return func() int {
        retVal := curVal
        curVal = curVal + step
        return retVal
    }
}