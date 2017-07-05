package main

import "fmt"

func main() {

	str := "Hello, World!"
	n := len(str)

	for i := 0; i < n; i++ {
		ch := str[i]
		fmt.Printf("%2d\t%3d\t%c\n", i, ch, ch)
	}
}
