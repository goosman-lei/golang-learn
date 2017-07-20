package main

import (
    "fmt"
)

func main() {
    s1 := []int{1, 2, 3, 4, 5}
    s2 := []string{"A", "B", "C", "D", "E"}
    s := make([]interface{}, 5)

    fmt.Println(s1)
    for i, v := range s1 { s[i] = v }
    s = mapFunc(s)
    for i, v := range s { s1[i] = v.(int) }
    fmt.Println(s1)
    fmt.Println(s2)
    for i, v := range s2 { s[i] = v }
    s = mapFunc(s)
    for i, v := range s { s2[i] = v.(string) }
    fmt.Println(s2)
}

func mapFunc(s []interface{}) []interface{} {
    for idx, val := range s {
        switch v := val.(type) {
            case int:
                s[idx] = 2 * v
            case string:
                s[idx] = v + v
            default:
                s[idx] = v
        }
    }
    return s
}