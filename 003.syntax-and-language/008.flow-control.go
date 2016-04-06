package main

import "fmt"

func main() {
	i := 0
LOOP:
	fmt.Println(i)
	i++
	if i < 10 {
		goto LOOP
	}
}
