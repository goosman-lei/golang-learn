package main

import (
    "fmt"
)

type Person struct {
    firstName string
    lastName string
}

func (p *Person) FirstName() string {
    return p.firstName
}
func (p *Person) SetFirstName(name string) {
    p.firstName = name
}

func main() {
    p := new(Person)
    p.SetFirstName("Eric")
    fmt.Println(p.FirstName())
}