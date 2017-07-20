package main

import (
    "fmt"
    "time"
)

type MyTime struct {
    time.Time
}

func (this MyTime) first3Chars() string {
    return this.String()[0:3]
}

func main() {
    m := MyTime{time.Now()}
    fmt.Printf("Full time now: %s\n", m.String())
    fmt.Printf("First 3 chars: %s\n", m.first3Chars())
}