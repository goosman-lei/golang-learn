package main

import "fmt"
import "sync"
import "time"

var once sync.Once

func main() {
	ch := make(chan int)
	go doSomething(1, ch)
	go doSomething(2, ch)
	go doSomething(3, ch)
	ch <- 1
	ch <- 2
	ch <- 3
	time.Sleep(1 * time.Second)
}

func doInit() {
	fmt.Println("doInit")
}

func doSomething(n int, ch chan int) {
	once.Do(doInit)
	for {
		i := <-ch
		fmt.Println("doSomething", n, i)
	}
}
