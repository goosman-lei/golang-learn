package main

import (
    "fmt"
    "bytes"
)

func main() {
    buffer := bytes.NewBuffer([]byte("Hello World"))
    left, right := splitBuffer(buffer, 5)

    fmt.Printf("left: %v\n", left)
    fmt.Printf("right: %v\n", right)
}

func splitBuffer(buffer *bytes.Buffer, n int) (*bytes.Buffer, *bytes.Buffer) {
    byteArr := buffer.Bytes()

    return bytes.NewBuffer(byteArr[:n]), bytes.NewBuffer(byteArr[n:])
}