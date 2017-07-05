package main

import "fmt"
import "os"
import "reflect"
import "runtime"

func main() {
	args := make([]interface{}, len(os.Args))
	for i, v := range os.Args {
		args[i] = v
	}

	invoke(demoInfinitiveArgFunc, 1, int32(2), 3, int64(4), 5)

	invoke(demoInfinitiveArgFunc, args...)

	invoke(demoAnonymousFunc)
}

func demoInfinitiveArgFunc(args ...interface{}) {
	for idx, arg := range args {
		fmt.Println("index:", idx, "value:", arg, "type:", reflect.TypeOf(arg))
	}
}

func demoAnonymousFunc() {
	anonymousFuncAdd := func(x, y int) int {
		return x + y
	}
	fmt.Println("3 + 4 =", anonymousFuncAdd(3, 4))
}

func trace() {
	pc, fname, lineno, _ := runtime.Caller(1)
	f := runtime.FuncForPC(pc)
	fmt.Printf("%s:%d %s\n", fname, lineno, f.Name())
}

// 反射调用
func invoke(funcObj interface{}, args ...interface{}) {
	// reflect.Value对象
	funcCallable := reflect.ValueOf(funcObj)
	inParams := make([]reflect.Value, len(args))
	for k, arg := range args {
		inParams[k] = reflect.ValueOf(arg)
	}

	// runtime.Func对象
	fmt.Println("invoke func:", runtime.FuncForPC(reflect.ValueOf(funcObj).Pointer()).Name())

	funcCallable.Call(inParams)

	fmt.Println()
}

func dumpMethods(obj interface{}) {
	objType := reflect.TypeOf(obj)
	for i := 0; i < objType.NumMethod(); i++ {
		fmt.Println(objType.Method(i).Name)
	}
}
