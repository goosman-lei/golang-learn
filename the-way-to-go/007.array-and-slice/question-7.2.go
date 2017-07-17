package main

import "fmt"

func main() {
    b := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
    fmt.Printf("%s\n", b)
    fmt.Printf("%s %s %s %s\n", b[1:4], b[:2], b[2:], b[:])
}