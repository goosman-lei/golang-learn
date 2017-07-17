package main

import (
    "fmt"
)

func main() {
    items := make([]map[string]int, 5)
    for i := range items {
        items[i] = make(map[string]int)
        items[i]["index"] = i
    }
    items[0]["Mon"] = 1
    items[0]["Tue"] = 2
    items[0]["Wed"] = 3
    items[0]["Thu"] = 4
    items[0]["Fri"] = 5
    items[0]["Sat"] = 6
    items[0]["Sun"] = 7
    fmt.Printf("Version A: Value of items %v\n", items)
    fmt.Println("Version A verbose:")
    for i, m := range items {
        fmt.Printf("slice-index: %d, slice-value-pointer: %p\n", i, m)
    }
    fmt.Println()

    items2 := make([]map[string]int, 5)
    for i, item := range items2 {
        // here, item is only a copy of slice element
        item = make(map[string]int)
        // item will be lost in next iteration
        item["index"] = i
    }
    fmt.Printf("Version B: Value of items %v\n", items2)
    fmt.Println("Version B verbose:")
    for i, m := range items {
        fmt.Printf("slice-index: %d, slice-value-pointer: %p\n", i, m)
    }
    fmt.Println()
}