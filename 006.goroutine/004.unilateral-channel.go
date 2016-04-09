package main

import "fmt"

func main() {
	defer func() {
		recover()
	}()
	var ch chan int        // 正常channel
	var chWrite chan<- int // 单向写channel
	var chRead <-chan int  // 单向读channel

	ch = make(chan int, 3)
	chWrite = chan<- int(ch) // 通过类型转换, 设置单向channel的值
	chRead = (<-chan int)(ch)

	chWrite <- 1
	n := <-chRead
	fmt.Println("readed:", n)

	close(ch)

	var isOk bool

	_, isOk = <-ch
	fmt.Println(isOk)
	_, isOk = <-chRead
	fmt.Println(isOk)
}
