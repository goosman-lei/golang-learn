package main

import "fmt"

func main() {
    fabonacciArr := fabonacciArr(13)
    fmt.Printf("%v\n", fabonacciArr)
}

func fabonacciArr(n int) []int {
    retArr := make([]int, n + 1)
    switch {
        case n < 0:
            return retArr
        case n == 0:
            retArr[0] = 1
            return retArr
        case n == 1:
            retArr[0] = 1
            retArr[1] = 1
            return retArr
        default:
            retArr[0] = 1
            retArr[1] = 1
            for i := 2; i <= n; i ++ {
                retArr[i] = retArr[i - 1] + retArr[i - 2]
            }
            return retArr
    }
}