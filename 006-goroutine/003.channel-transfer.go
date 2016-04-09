package main

import "fmt"
import "time"

type PipeData struct {
	value int
	next  chan PipeData
}

func main() {
	ch1 := make(chan PipeData)
	ch2 := make(chan PipeData)

	go step1(ch1)
	go step2(ch2)

	pd := PipeData{1, ch2}

	ch1 <- pd

	time.Sleep(1 * time.Second)
}

func step1(ch chan PipeData) {
	pd := <-ch
	fmt.Println("step1:", pd.value)
	pd.next <- pd
}
func step2(ch chan PipeData) {
	pd := <-ch
	fmt.Println("step2:", pd.value)
	pd.next <- pd
}
