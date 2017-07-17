package main

import (
    "./greeting"
    "fmt"
)

func main() {
    if greeting.IsMorning() {
        fmt.Println(greeting.GreetingMorning)
    } else if greeting.IsAfternoon() {
        fmt.Println(greeting.GreetingAfternoon)
    } else if greeting.IsEvening() {
        fmt.Println(greeting.GreetingNight)
    }
}
