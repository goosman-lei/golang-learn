package main

import "fmt"

func main() {
    // 无类型浮点常量
    const Pi float64 = 3.1415926535897932354626
    const zero = 0.0
    // 无类型整型常量
    const (
        size int64 = 1024
        eof = -1
    )
    // u = 0.0, v = 3.0, 常量的多重赋值
    const u, v float32 = 0, 3
    // 无类型整型和字符串常量 a = 3, b = 4, c = "foo"
    const a, b, c = 3, 4, "foo"

    // 常量定义的右值可以是可在编译期运行的常量表达式
    const mask = 1 << 3
    // 下面是错误的写法
    // const Home = os.GetEnv("HOME")
}