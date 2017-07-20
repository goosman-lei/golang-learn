package main

import (
    "fmt"
    "reflect"
)

func main() {
    x := 3.4
    v := reflect.ValueOf(x)
    // vE := v.Elem() // panic when Value contains value is not interface or pointer
    vP := reflect.ValueOf(&x)
    vPE := vP.Elem()

    fmt.Printf("x: %p\t&x: %p\n", x, &x)
    fmt.Printf("v: %p\t&v: %p\n", v, &v)
    fmt.Printf("vP: %p\t&vP: %p\n", vP, &vP)
    fmt.Printf("vPE: %p\t&vPE: %p\n", vPE, &vPE)

    fmt.Println("settablility of v:", v.CanSet())
    vPE.SetFloat(3.1415)
    fmt.Println(v)

    fmt.Println("type of vP:", vP.Type())
    fmt.Println("settablility of vP:", vP.CanSet())

    fmt.Println("The Elem of vP is: ", vPE)
    fmt.Println("settablility of vPE:", vPE.CanSet())
    vPE.SetFloat(3.1415926)
    fmt.Println(vPE.Interface())
    fmt.Println(vPE)
}