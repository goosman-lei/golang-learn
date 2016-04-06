package main

import "fmt"

func main() {
	var i int = 1

	makeFuncFuncA := func() func() {
		var j int = 10
		return func() {
			i++
			fmt.Printf("i, j: %d, %d\n", i, j)
		}
	}
	makeFuncFuncB := func() func() {
		var j int = 10
		return func() {
			i++
			fmt.Printf("i, j: %d, %d\n", i, j)
		}
	}

	funcAA := makeFuncFuncA()
	funcBA := makeFuncFuncB()

	i = 2

	funcAB := makeFuncFuncA()
	funcBB := makeFuncFuncB()

	funcAA()
	funcBA()
	funcAB()
	funcBB()

	i = 3

	funcAA()
	funcBA()
	funcAB()
	funcBB()
}

/*
$ go run 010.closure.go
i, j: 7, 10
i, j: 8, 10
i, j: 8, 10
i, j: 9, 10

Notice: 任何类型的闭包, 对同一个变量的使用, 都是内存共享的. 任意改变会影响其他闭包
*/
