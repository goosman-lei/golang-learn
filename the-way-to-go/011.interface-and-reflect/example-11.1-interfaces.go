package main

import "fmt"

type Shaper interface {
    Area() float32
}

type Square struct {
    side float32
}

func (sq *Square) Area() float32 {
    return sq.side * sq.side
}

func main() {
    sq1 := new(Square)
    sq1.side = 5

    /*
    // 定义接口类型变量, 然后赋值实现的实例
    var shaper Shaper
    shaper = sq1

    // 强制类型转换并初始化赋值接口变量
    shaper := Shaper(sq1)
    */

    shaper := sq1
    fmt.Printf("The square has area: %f\n", shaper.Area())
}