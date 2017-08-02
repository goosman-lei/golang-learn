package main

import (
    "fmt"
    "net/http"
    "time"
)

var urls = []string{
    "http://www.baidu.com",
    "http://www.qq.com",
    "http://www.tmall.com",
}

func main() {
    c := http.Client{Timeout: 3 * time.Second}
    for _, url := range urls {
        resp, err := c.Head(url)
        if err != nil {
            fmt.Printf("Error[%s]: %s\n", url, err)
        } else {
            fmt.Printf("Success[%s]: %s\n", url, resp.Status)
        }
    }
}