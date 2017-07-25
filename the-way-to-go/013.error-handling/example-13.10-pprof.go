package main

import (
    _ "net/http/pprof"
    "net/http"
)

func main() {
    http.HandleFunc("/hello", HelloAction)
    http.ListenAndServe("101.200.125.244:8080", nil)
}

func HelloAction(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello"))
}