package main

import "fmt"

func main() {
    var i, j, k int

    i, _, j = 1, 2, 3

    k = 3

    k ++

    fmt.Printf("i = %d, j = %d, k = %d\n", i, j, k)
}