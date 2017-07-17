package main

import (
    "fmt"
)

func main() {
    slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
    fmt.Println(filter(slice, func(idx, ele int) bool {
        return ele % 2 == 0
    }))
    fmt.Println(filter(slice, func(idx, ele int) bool {
        return ele < 5
    }))
}

func filter(slice []int, fn func(idx, ele int) bool) []int {
    result := make([]int, 0, len(slice))
    for idx, ele := range slice {
        if (fn(idx, ele)) {
            result = append(result, ele)
        }
    }
    return result
}