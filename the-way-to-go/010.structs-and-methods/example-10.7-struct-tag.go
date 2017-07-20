package main

import (
    "fmt"
    "reflect"
)

type TagType struct {
    field1 bool "An important answer"
    field2 string "The name of the thing"
    field3 int "How much there are"
}

func main() {
    tt := TagType{true, "Barak Obama", 1}
    for i := 0; i < 3; i ++ {
        ttType := reflect.TypeOf(tt)
        idxField := ttType.Field(i)
        fmt.Printf("%s.%s.%s[%s]: %v\n", idxField.PkgPath, ttType.Name(), idxField.Name, idxField.Type, idxField.Tag)
    }
}