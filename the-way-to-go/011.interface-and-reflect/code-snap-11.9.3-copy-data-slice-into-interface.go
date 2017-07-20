package main

import "fmt"

func main() {
    var dataSlice []int = []int{1, 2, 3, 4, 5}
    var interfaceSlice []interface{} = make([]interface{}, len(dataSlice))
    for idx, d := range dataSlice {
        interfaceSlice[idx] = d
    }

    fmt.Printf("%T %#v\n", interfaceSlice, interfaceSlice)
    for _, d := range interfaceSlice {
        fmt.Printf("%T %#v\n", d, d)
    }
}