package main

import "fmt"

func main() {
    var fibonaccis [50]int

    fibonaccis[0] = 1
    fibonaccis[1] = 1
    for i := 2; i < len(fibonaccis); i ++ {
        fibonaccis[i] = fibonaccis[i - 1] + fibonaccis[i - 2]
    }
    fmt.Printf("%+v\n", fibonaccis)
}