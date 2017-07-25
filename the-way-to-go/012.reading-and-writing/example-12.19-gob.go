package main

import (
    "encoding/gob"
    "log"
    "os"
    "fmt"
)

type Address struct {
    Type string
    City string
    Country string
}

type VCard struct {
    FirstName string
    LastName string
    Addresses []*Address
    Remark string
}

func main() {
    pa := &Address{"private", "Aartselaar", "Belgium"}
    wa := &Address{"work", "Boom", "Belgium"}
    vc := VCard{"Jan", "Kersschot&<script>", []*Address{pa, wa}, "none"}
    fmt.Printf("vcard: %v\n", vc)

    // using an encoder:
    file, _ := os.OpenFile("vcard.gob", os.O_CREATE | os.O_WRONLY, 0755)
    defer file.Close()
    enc := gob.NewEncoder(file)
    err := enc.Encode(vc)
    if err != nil {
        log.Println("Error in encoding gob")
    }
}