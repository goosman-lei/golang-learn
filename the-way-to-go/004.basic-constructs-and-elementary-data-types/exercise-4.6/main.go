package main

import (
    "fmt"
    "unicode/utf8"
)

func main() {
    fmt.Printf("strLen is: %d\n", strLen("asSASA ddd dsjkdsjs dk中国"))
    fmt.Printf("utf8StrLen is: %d\n", utf8StrLen("asSASA ddd dsjkdsjs dk中国"))
}

func utf8StrLen(str string) (len int) {
    len = utf8.RuneCountInString(str)
    return
}

func strLen(str string) (len int) {
    i := -1
    for i, _ = range str {
    }
    len = i + 1
    return
}