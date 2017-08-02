package main

import (
    "fmt"
    "log"
    "net/rpc"
    "multiply"
)

func main() {
    client, err := rpc.DialHTTP("tcp", "localhost:8088")
    if err != nil {
        log.Fatal("Error dialing:", err)
    }

    // Synchronous call
    args := &multiply.Args{7, 8}
    var reply int
    err = client.Call("Args.Multiply", args, &reply)
    if err != nil {
        log.Fatal("Args error:", err)
    }
    fmt.Printf("Args: %d * %d = %d\n", args.N, args.M, reply)
}
