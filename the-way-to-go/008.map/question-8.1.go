package main

import "fmt"

func main() {
    capitals := map[string]string{"Frace": "Paris", "Italy": "Rome", "Japan": "Tokyo"}

    for key := range capitals {
        fmt.Println("Map item: Capital of", key, "is", capitals[key])
    }
}