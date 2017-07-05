package main

import "fmt"

type Logger struct {
}

func (logger Logger) Info() {
	fmt.Println("log info in Logger.Info")
}

func (logger Logger) Error() {
	fmt.Println("log error in Logger.Info")
}

type Logic struct {
	*Logger
}

func (logic Logger) Buy() {
	fmt.Println("Logger.Buy")
}

func main() {
	logger := new(Logger)
	logic := new(Logic)
	logic.Logger = logger

	logic.Info()
	logic.Error()
	logic.Buy()
}
