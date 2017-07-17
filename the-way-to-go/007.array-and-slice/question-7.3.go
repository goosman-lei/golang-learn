package main
import "fmt"

func main() {
    s := make([]byte, 5)
    fmt.Println(len(s), cap(s))

    s = s[2:4]
    fmt.Println(len(s), cap(s))
}