package main

import "fmt"
import "time"

func main() {
	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(1 * time.Second)
		timeout <- true
	}()

	ch := make(chan int, 1)

	var now time.Time
	now = time.Now()
	fmt.Println(now.Unix(), now.Nanosecond())
	select {
	case <-ch:
		fmt.Println("Received")
	case <-timeout:
		fmt.Println("Timeout")
	}
	now = time.Now()
	fmt.Println(now.Unix(), now.Nanosecond())
}
