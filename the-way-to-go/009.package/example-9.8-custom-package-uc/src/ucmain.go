package main

import (
    "fmt"
    "./uc"
)

func main() {
    str := "USING package uc!"
    fmt.Println(uc.UpperCase(str))

    fmt.Println(uc.Text)
}