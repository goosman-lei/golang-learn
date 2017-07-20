package main

import (
    "fmt"
    "reflect"
)

type NotknownType struct {
    S1, s2, S3 string
}

func (n NotknownType) String() string {
    return n.S1 + "-" + n.s2 + "-" + n.S3
}

var secret interface{} = NotknownType{"Ada", "Go", "Oberon"}

func main() {
    value := reflect.ValueOf(secret) // <main.NotknownType Value>
    valueP := reflect.ValueOf(&secret)
    typ := reflect.TypeOf(secret) // main.NotknownType
    // alternative:
    // typ := value.Type() // main.NotknownType

    fmt.Println(&value)
    fmt.Println(typ)
    knd := value.Kind() // struct
    fmt.Println(knd)

    // iterate through the fields of the struct:
    for i := 0; i < value.NumField(); i ++ {
        fmt.Printf("Field %d: %v\n", i, value.Field(i))
        fmt.Printf("%v\n", value.Field(i).CanSet())
        // value.Field(i).Elem().SetString("C#")
    }

    // call the first method, which is String():
    results := value.Method(0).Call(nil)
    fmt.Println(results) // [Ada - Go - Oberon]

    fmt.Printf("valueP.Elem(): %#v valueP.Elem().CanAddr(): %t valueP.Elem().CanSet(): %t\n", valueP.Elem(), valueP.Elem().CanAddr(), valueP.Elem().CanSet())

    s := NotknownType{"Ada", "Go", "Oberon"}
    v := reflect.ValueOf(&s).Elem()
    for i := 0; i < v.NumField(); i ++ {
        fmt.Printf("Field %d: %p\n", i, value.Field(i))
    }
}