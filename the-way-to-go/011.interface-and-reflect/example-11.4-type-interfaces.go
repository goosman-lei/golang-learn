package main

import (
    "fmt"
    "math"
)

type Square struct {
    side float32
}

type Circle struct {
    radius float32
}

type Shaper interface {
    Area() float32
}

func main() {
    // areaIntf := Shaper(&Square{5})
    areaIntf := Shaper(&Circle{5})

    switch t := areaIntf.(type) {
        case *Square:
            fmt.Printf("Type Square %T with value %v\n", t, t)
        case *Circle:
            fmt.Printf("Type Circle %T with value %v\n", t, t)
        case nil:
            fmt.Printf("nil value: nothing to check?\n")
        default:
            fmt.Printf("Unexpected type %T\n", t)
    }

    if t, ok := areaIntf.(*Square); ok {
        fmt.Printf("The type of areaIntf is: %T\n", t)
    }

    if u, ok := areaIntf.(*Circle); ok {
        fmt.Printf("The type of areaIntf is: %T\n", u)
    } else {
        fmt.Println("areaIntf does not contain a variable of type Circle")
    }
}

func (sq *Square) Area() float32 {
    return sq.side * sq.side
}

func (c *Circle) Area() float32 {
    return math.Pi * c.radius * c.radius
}