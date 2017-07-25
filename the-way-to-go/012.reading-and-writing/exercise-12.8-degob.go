package main

import (
    "fmt"
    "encoding/gob"
    "os"
)

type Address struct {
    Type string
    City string
    Country string
}

type VCard struct {
    FirstName string
    LastName string
    Address []*Address
    Remark string
}

func main() {
    ifp, err := os.Open("vcard.gob")
    if err != nil {
        fmt.Printf("open vcard.gob fail: %s\n", err.Error())
        os.Exit(1)
    }

    var v VCard
    decoder := gob.NewDecoder(ifp)
    err = decoder.Decode(&v)
    if err != nil {
        fmt.Printf("decode gob failed: %s\n", err.Error())
        os.Exit(1)
    }

    fmt.Println(v)
}
