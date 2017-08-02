package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "log"
    "flag"
)

var url = flag.String("url", "", "The url what you want to fetch")

func main() {
    flag.Parse()
    if len(*url) <= 0 {
        fmt.Printf("Error: have no url\n")
        return
    }

    res, err := http.Get(*url)
    CheckError(err)

    data, err := ioutil.ReadAll(res.Body)
    CheckError(err)

    fmt.Printf("Got: %q", string(data))
}

func CheckError(err error) {
    if err != nil {
        log.Fatalf("Get: %v", err)
    }
}