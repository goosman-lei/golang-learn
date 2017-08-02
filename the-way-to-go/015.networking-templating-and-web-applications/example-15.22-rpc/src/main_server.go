package main

import (
    "multiply"
    "net/http"
    "log"
    "net"
    "net/rpc"
    "time"
)

func main() {
    o := new(multiply.Args)

    rpc.Register(o)
    rpc.HandleHTTP()

    listener, e := net.Listen("tcp", "0.0.0.0:8088")
    if e != nil {
        log.Fatal("Starting RPC-Server listen error:", e)
    }

    go http.Serve(listener, nil)
    time.Sleep(1000e9)
}
