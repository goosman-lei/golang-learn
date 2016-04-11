package main

import "net/rpc"
import "fmt"

type Args struct {
	A, B int
}

func main() {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:9999")
	if err != nil {
		panic(err)
	}

	args := Args{7, 8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Arith: %d * %d = %d\n", args.A, args.B, reply)
}
