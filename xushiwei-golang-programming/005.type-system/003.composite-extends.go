package main

import "fmt"
import "reflect"

type Flyable struct {
	flyHeight int
}

func (f Flyable) Fly() {
	fmt.Println("Fly by Flyable")
}

type Animal struct {
	weight int
}

func (a Animal) Cry() {
	fmt.Println("Cry by Animal")
}
func (a Animal) Eat() {
	fmt.Println("Eat by Animal")
}

type Bird struct {
	Animal
	Flyable
}

func (b Bird) Cry() {
	fmt.Println("Cry by Bird")
}

type Sparrow struct {
	Bird
	flyHeight float64
}

func (s Sparrow) Eat() {
	fmt.Println("Eat by Sparrow")
}

func main() {
	sparrow := new(Sparrow)
	sparrow.flyHeight = 100
	sparrow.weight = 10

	// 方法继承关系
	sparrow.Cry()
	sparrow.Fly()
	sparrow.Eat()

	// 继承到的成员
	fmt.Println("weight:", sparrow.weight, "flyHeight:", sparrow.flyHeight, "flyHeight type:", reflect.TypeOf(sparrow.flyHeight))
}

/*
Notice:
	1. 允许多重继承. 但是如果继承到的方法名有冲突, 则会提示歧义
	2. 成员属性的类型同样可以被继承覆盖
*/
