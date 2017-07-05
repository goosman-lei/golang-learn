package main

import "fmt"
import "sync"
import "time"
import "runtime"

var pool []int
var poolLock sync.RWMutex

func main() {
	go producer(1)
	go producer(2)
	go producer(3)
	go consumer(1)
	go consumer(2)
	go consumer(3)
	time.Sleep(1 * time.Second)
}

func consumer(n int) {
	var ele int
	for {
		//time.Sleep(1000 * time.Nanosecond)
		poolLock.RLock()
		if len(pool) == 0 {
			fmt.Println("\tconsume", n, "pool is empty")
		} else {
			ele, pool = pool[0], pool[1:]
			fmt.Println("\tconsum", n, "value", ele)
		}
		poolLock.RUnlock()
		fmt.Println("\tr unlock", n)
		runtime.Gosched()
	}
}

func producer(n int) {
	i := 0
	for {
		//time.Sleep(1000 * time.Nanosecond)
		poolLock.Lock()
		if len(pool) == 10 {
			fmt.Println("produce", n, "pool is full")
		} else {
			pool = append(pool, i+n*10000000)
			fmt.Println("produce", n, "value", i)
			i++
		}
		poolLock.Unlock()
		fmt.Println("unlock", n)
		runtime.Gosched()
	}

}
