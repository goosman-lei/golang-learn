package main

import (
    "fmt"
)

func main() {
    data_arr := []int{1, 2, 3, 4, 5, 6}
    data_slice := data_arr[:]
    data_slice2 := data_arr[2:4]

    fmt.Printf("%p %p %p %p\n", data_arr, data_slice, data_slice2, &data_arr[2])

    fmt.Printf("data_arr: %v, data_slice: %v, data_slice2: %v\n", data_arr, data_slice, data_slice2)
    data_slice[1] = 8
    fmt.Printf("data_arr: %v, data_slice: %v, data_slice2: %v\n", data_arr, data_slice, data_slice2)
}