package main

import (
    "fmt"
    "strings"
    "strconv"
)

func main() {
    fmt.Printf("\"Hello\" is prefix of \"Hello World\": %t\n", strings.HasPrefix("Hello World", "Hello"))
    fmt.Printf("\"World\" is prefix of \"Hello World\": %t\n", strings.HasPrefix("Hello World", "World"))
    fmt.Println()

    fmt.Printf("\"World\" is contained in \"Hello World\": %t\n", strings.Contains("Hello World", "World"))
    fmt.Printf("The letter \"l\" first index of \"Hello World\": %d\n", strings.Index("Hello World", "l"))
    fmt.Printf("The letter \"l\" last index of \"Hello World\": %d\n", strings.LastIndex("Hello World", "l"))
    fmt.Println()

    fmt.Printf("Replace all \"l\" with \"~\" in \"Hello World\": %s\n", strings.Replace("Hello World", "l", "~", -1))
    fmt.Printf("Times of  letter \"l\" occurs in \"Hello World\": %d\n", strings.Count("Hello World", "l"))
    fmt.Printf("Repeat letter \"l\" 5 times: %s\n", strings.Repeat("l", 5))
    fmt.Println()

    var buffer [10]byte
    str := "Hello world, This is a coder!"
    strReader := strings.NewReader(str)
    fmt.Printf("Len: %d offset: %+v\n", strReader.Len(), strReader)
    n, _ := strReader.Read(buffer[1:3])
    fmt.Printf("Read Len: %d Buffer: %s\n", n, buffer)
    fmt.Println()

    fmt.Printf("0x103 itoa: %s\n", strconv.Itoa(0x103))
    fmt.Printf("3.29372 format: %s\n", strconv.FormatFloat(3.29372, 'b', 3, 64))
    fmt.Println()
}