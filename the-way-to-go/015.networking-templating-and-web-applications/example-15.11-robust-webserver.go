package main

import (
    "net/http"
    "io"
    "log"
    "fmt"
)

type HandlerFunc func(http.ResponseWriter, *http.Request)

func main() {
    http.HandleFunc("/", wrapperHandler(action_index_handler))
    err := http.ListenAndServe("0.0.0.0:8088", nil)
    if err != nil {
        fmt.Printf("web-server startup failed: %s\n", err.Error())
    }
}

func wrapperHandler(handler HandlerFunc) HandlerFunc {
    return func(w http.ResponseWriter, req *http.Request) {
        defer func() {
            if r := recover(); r != nil {
                log.Printf("panic: %s\n", r)
            }
        }()
        handler(w, req)
    }
}

func action_index_handler(w http.ResponseWriter, req *http.Request) {
    io.WriteString(w, "Hello World\n")
    panic("Hello World")
}