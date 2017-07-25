package main

import (
    "fmt"
    "errors"
//  "strconv"
)

/*
//Canonicity Value
const maxInt int64 = (1 << (strconv.IntSize - 1)) - 1
const minInt int64 = ^maxInt
*/

// Test Value
const maxInt int64 = 1e10
const minInt int64 = -1e10

func main() {
    v1, e := IntFromInt64(1e18);
    fmt.Printf("%E => %d [e: %v]\n", 1e18, v1, e)
    v2, e := IntFromInt64(1e8);
    fmt.Printf("%E => %d [e: %v]\n", 1e13, v2, e)
    v3, e := IntFromInt64(-1e18);
    fmt.Printf("%E => %d [e: %v]\n", -1e18, v3, e)
}

func IntFromInt64(num int64) (ret int, err error) {
    defer func() {
        if r := recover(); r != nil {
            err = errors.New(r.(string))
        }
    }()
    ret = ConvertInt64ToInt(num)
    return
}

func ConvertInt64ToInt(num int64) int {
    if num > maxInt {
        panic(fmt.Sprintf("Exceed max value: %d", maxInt))
    } else if num < minInt {
        panic(fmt.Sprintf("Exceed min value: %d", minInt))
    }
    return int(num)
}