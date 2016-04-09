package main

import "fmt"

func Count(ch chan int) {
	received := <-ch
	fmt.Println("Counting:", received)
	ch <- 0
}

func main() {
	fmt.Println("channel community for goroutines")
	chs := make([]chan int, 10)
	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go Count(chs[i])
	}
	for i, ch := range chs {
		ch <- i
		<-ch
		fmt.Println("\tchannel:", i, "done")
	}
	fmt.Println()

	var ch chan int

	fmt.Println("channel select syntax")
	ch = make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case ch <- 0:
		case ch <- 1:
		}
		n := <-ch
		fmt.Println("Value received:", n)
	}
	fmt.Println()

	fmt.Println("buffered channel")
	ch = make(chan int, 10)
	for i := 0; i < 10; i++ {
		ch <- i
		fmt.Println("Value sended:", i)
	}

	for i := 0; i < 10; i++ {
		n, err := <-ch
		fmt.Println("Value received:", n, err)
	}
	fmt.Println()
}
