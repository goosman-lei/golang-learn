package main

import (
    "fmt"
)

type Any interface{}
type EvalFunc func(Any) (Any, Any)

type Fibonacci struct {
    n int
    c int
    p int
}

func main() {
    demoEven()
    demoFibonacci()
}

func demoFibonacci() {
    fibonacci := BuildLazyFibonacciEvaluator(Fibonacci{0, 1, 0})

    for i := 0; i < 20; i ++ {
        f := fibonacci()
        fmt.Printf("%vth fibonacci: %v\n", f.n, f.c)
    }

}

func demoEven() {
    evenFunc := func(state Any) (Any, Any) {
        os := state.(int)
        ns := os + 2
        return os, ns
    }

    even := BuildLazyIntEvaluator(evenFunc, 0)

    for i := 0; i < 10; i ++ {
        fmt.Printf("%vth even: %v\n", i, even())
    }
}

func BuildLazyEvaluator(evalFunc EvalFunc, initState Any) func() Any {
    retValChan := make(chan Any)
    // value generator
    loopFunc := func() {
        var actState Any = initState
        var retVal Any
        for {
            retVal, actState = evalFunc(actState)
            retValChan <- retVal
        }
    }
    // function that read return value from output channel
    retFunc := func() Any {
        return <-retValChan
    }
    go loopFunc()
    return retFunc
}

func BuildLazyIntEvaluator(evalFunc EvalFunc, initState Any) func() int {
    ef := BuildLazyEvaluator(evalFunc, initState)
    return func() int{
        return ef().(int)
    }
}

func BuildLazyFibonacciEvaluator(initState Any) func() Fibonacci {
    ef := BuildLazyEvaluator(func(initState Any) (Any, Any) {
        f := initState.(Fibonacci)
        f.n ++
        f.p, f.c = f.c, f.c + f.p
        return f, f
    }, Fibonacci{1, 1, 1})
    return func() Fibonacci {
        return ef().(Fibonacci)
    }
}