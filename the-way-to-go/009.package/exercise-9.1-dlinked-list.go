package main

import (
    "fmt"
    "container/list"
)

func main() {
    l := list.New()

    l.PushBack(101)
    l.PushBack(102)
    l.PushBack(103)

    for i, e := 0, l.Front(); i < l.Len(); i, e = i + 1, e.Next() {
        fmt.Printf("index: %d, element: %+v, element.Value: %+v\n", i, e, e.Value)
    }
}