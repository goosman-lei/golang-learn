package main

import "fmt"
import "net"
import "time"

func main() {

	go serverListen()

	go clientConnect()

	time.Sleep(10 * time.Second)
}

func serverListen() {
	sAddr := &net.TCPAddr{net.ParseIP("0.0.0.0"), 9999, ""}
	sConn, err := net.ListenTCP("tcp", sAddr)
	defer sConn.Close()
	if err != nil {
		panic(err)
	}

	for {
		cConn, err := sConn.Accept()
		defer cConn.Close()
		if err != nil {
			panic(err)
		}

		fmt.Println("connect from:", cConn.RemoteAddr())
		go serverWorker(cConn)
	}
}

func serverWorker(conn net.Conn) {
	var rBuff, wBuff []byte
	rBuff = make([]byte, 100)

	fmt.Println("server ready to read msg")
	nRead, _ := conn.Read(rBuff)
	fmt.Println("server", string(rBuff[0:nRead]))
	wBuff = []byte("Hello World From Server")
	conn.Write(wBuff)
	fmt.Println("server worker done")
}

func clientConnect() {
	cConn, err := net.Dial("tcp", "127.0.0.1:9999")
	defer cConn.Close()
	if err != nil {
		panic(err)
	}

	var rBuff, wBuff []byte
	rBuff = make([]byte, 100)

	wBuff = []byte("Hello World From Client")

	fmt.Println("client ready to write msg:", string(wBuff))
	_, err = cConn.Write(wBuff)
	nRead, _ := cConn.Read(rBuff)
	fmt.Println("client", string(rBuff[0:nRead]))
	fmt.Println("client done")
}
