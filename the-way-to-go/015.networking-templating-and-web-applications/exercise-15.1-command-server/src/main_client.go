package main

import (
    "net"
    "fmt"
    "bufio"
    "strings"
)

func main() {
    conn, err := net.Dial("tcp", "127.0.0.1:8088")
    if err != nil {
        fmt.Printf("connect to server failed: %s\n", err.Error())
        return
    }
    reader := bufio.NewReader(conn)

    conn.Write([]byte("sayhi: Hi server\n"))
    conn.Write([]byte("heartbeat\n"))
    conn.Write([]byte("echo: This is echo message\n"))
    if resp, err := reader.ReadString('\n'); err != nil {
        fmt.Printf("Read message from server error: %s\n", err.Error())
    } else {
        fmt.Printf("Read message from server: %s\n", strings.TrimRight(resp, "\r\n"))
    }
}
