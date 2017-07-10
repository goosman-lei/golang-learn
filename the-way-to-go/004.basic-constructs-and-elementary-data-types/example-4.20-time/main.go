package main

import (
    "fmt"
    "time"
)

var week time.Duration

func main() {
    now := time.Now()
    fmt.Println(now)
    fmt.Printf("%02d.%02d.%04d\n", now.Day(), now.Month(), now.Year())
    fmt.Println()

    nowUTC := time.Now().UTC()
    fmt.Println(nowUTC)
    fmt.Println(time.Now())
    fmt.Println()

    week = 60 * 60 * 24 * 7 * 1e9
    weekFromNow := now.Add(week)
    fmt.Println(weekFromNow)
    fmt.Println()

    fmt.Println(now.Format(time.RFC822))
    fmt.Println(now.Format(time.ANSIC))
    fmt.Println(now.Format("02 Jan 2006 15:04"))
    str := now.Format("2006-01-02")
    fmt.Println(now, "=>", str)
    fmt.Println()
}