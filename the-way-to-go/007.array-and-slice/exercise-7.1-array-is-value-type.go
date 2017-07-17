package main

import "fmt"

func main() {
    var arr1 [5]int = [5]int{1, 2, 3}
    var arr2 [5]int

    arr2 = arr1
    arr2[3] = 8
    arr2[4] = 9

    for i := 0; i < len(arr1); i ++ {
        fmt.Println("arr1[", i, "]: ", arr1[i])
    }

    for i := 0; i < len(arr2); i ++ {
        fmt.Println("arr1[", i, "]: ", arr2[i])
    }
}