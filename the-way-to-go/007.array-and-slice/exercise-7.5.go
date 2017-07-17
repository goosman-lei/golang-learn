package main

import (
    "fmt"
    "bytes"
)

func main() {
    var slice []byte = []byte("Hello world")
    slice2 := Append(slice, []byte(", I'm Guoguo, I'm a coder of golang"))
    fmt.Printf("%s\n", slice)
    fmt.Printf("%s\n", slice2)

}

func Append(slice, data []byte) []byte {
    buffer := bytes.NewBuffer(slice)
    buffer.Write(data)
    return buffer.Bytes()
}