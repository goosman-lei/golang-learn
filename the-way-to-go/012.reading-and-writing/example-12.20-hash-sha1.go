package main

import (
    "fmt"
    "crypto/sha1"
    "io"
    "log"
)

func main() {
    hasher := sha1.New()
    io.WriteString(hasher, "test")
    b := []byte{}
    h1 := hasher.Sum(b)
    h2 := hasher.Sum(b)
    fmt.Printf("byte array: %p\n", &b)
    fmt.Printf("Result: %x %p\n", h1, h1)
    fmt.Printf("Result: %d %p\n", h2, h2)
    fmt.Printf("bytes array: %s\n", b[:])

    hasher.Reset()
    data := []byte("We shall overcome!")
    n, err := hasher.Write(data)
    if n != len(data) || err != nil {
        log.Printf("Hash write error: %v / %v", n, err)
    }
    checksum := hasher.Sum(b)
    fmt.Printf("Result: %x %p\n", checksum, checksum)
}