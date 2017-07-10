package main

import "fmt"

func main() {
    fmt.Println(fibonaciiClosure(10))
}

func fibonaciiClosure(n int) int {
    if n <= 1 {
        return 1
    }
    datas := make([]int, n + 1, 2 * n)
    datas[0] = 1
    datas[1] = 1
    for i := 2; i <= n; i ++ {
        datas[i] = datas[i - 1] + datas[i - 2]
    }
    return datas[n]
}

func fibonaciiDeducation(n int) int {
    if n <= 1 {
        return 1
    }
    x := 1
    y := 1
    for i := 2; i <= n; i ++ {
        x, y = y, x + y
    }
    return y
}