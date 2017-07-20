package main

import (
    "fmt"
    "reflect"
)

type Foo map[string]string
type Bar struct {
    thingOne string
    thingTwo int
}

func main() {
    // OK:
    y := new(Bar)
    (*y).thingOne = "hello"
    (*y).thingTwo = 1
    fmt.Printf("Bar y:\n%s\n", y)
    // Not OK:
    // z := make(Bar) // compile error: cannot make type Bar
    z := &Bar{}
    z.thingOne = "hello"
    z.thingTwo = 1
    fmt.Printf("Bar z:\n%s\n", z)
    // OK:
    x := make(Foo)
    x["x"] = "goodbye"
    x["y"] = "world"
    fmt.Printf("map x:\n")
    dumpMap(x)
    // Not OK:
    u := new(Foo)
    fmt.Printf("Bar u:\n")
    fmt.Printf("\tu type: %v\n", reflect.TypeOf(u))
    fmt.Printf("\tu value: %v\n", reflect.ValueOf(u))
    //(*u)["x"] = "goodbye" // !! panic !!: runtime error: assignment to entry in nil map
    //(*u)["y"] = "world"
    //fmt.Printf("Bar u:\n%s\n", *u)
}

func dumpMap(m map[string]string) {
    for k, v := range m {
        fmt.Printf("\t%s: %s\n", k, v)
    }
}
func (this *Bar)String() string {
    return fmt.Sprintf("\tthingOne: %s, thingTwo: %d", this.thingOne, this.thingTwo)
}