package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("defer func called")
		isPanicking := recover()
		fmt.Println("recover ret:", isPanicking)
	}()

	panic(404)
}

/*
Notice:
	1. panic: 中断当前执行流程, 逐级跳出并逆序执行defer
	2. recover: 需要定义在defer中. 截断panic中断过程
	3. 没有recover的panic流执行完成会导致所在goroutine所属进程打印异常信息退出
*/
