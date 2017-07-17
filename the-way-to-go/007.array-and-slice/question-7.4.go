package main

import "fmt"

func main() {
    s1 := []byte{'p', 'o', 'e', 'm'}
    s2 := s1[2:]

    s2[1] = 't'
    fmt.Printf("%s %s\n", s1, s2)
}