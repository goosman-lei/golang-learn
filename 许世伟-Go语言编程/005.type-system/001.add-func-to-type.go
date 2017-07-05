package main

import "fmt"

type Integer int

func (a Integer) Less(b Integer) bool {
	return a < b
}

func main() {
	var a Integer = 1

	if a.Less(2) {
		fmt.Println(a, "less 2")
	}
}

/*
Notice: 不能给非本地类型添加方法
*/
