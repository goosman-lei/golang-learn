package main

import "sync"
import "net"
import "net/http"
import "net/rpc"
import "fmt"
import "errors"

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func main() {
	var wg sync.WaitGroup

	arith := new(Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()

	l, e := net.Listen("tcp", ":9999")
	if e != nil {
		fmt.Println("listen error:", e)
	}

	wg.Add(1)
	go http.Serve(l, nil)

	wg.Wait()
}
