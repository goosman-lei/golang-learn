package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage:", os.Args[0], "<host>")
		os.Exit(1)
	}

	service := os.Args[1]

	conn, err := net.Dial("ip4:icmp", service)
	checkError(err)

	var msg [512]byte
	msg[0] = 8  // echo
	msg[1] = 0  // code 0
	msg[2] = 0  // checksum
	msg[3] = 0  // checksum
	msg[4] = 0  // identifier[0]
	msg[5] = 13 // identifier[1]
	msg[6] = 64 // sequence[0]
	msg[7] = 0  // sequence[1]

	len := 8
	check := checkSum(msg[0:len])
	fmt.Println("checksum:", check)
	msg[2] = byte(check >> 8)
	msg[3] = byte(check & 255)

	_, err = conn.Write(msg[0:len])
	fmt.Println("write")
	checkError(err)

	_, err = conn.Read(msg[0:])
	fmt.Println("read")
	checkError(err)

	fmt.Println("Got response")
	if msg[5] == 13 {
		fmt.Println("Identifier matches")
	}
	if msg[6] == 64 && msg[7] == 0 {
		fmt.Println("Sequence matches")
	}

	os.Exit(0)
}

func checkSum(msg []byte) uint16 {
	var sum uint32
	sum = 0
	n := 0

	// 前2*N字节, 每两字节看做整型求和
	for ; n < len(msg); n += 2 {
		sum += uint32(msg[n])*256 + uint32(msg[n+1])
	}
	// 若为奇数字节, 加最后一字节的值
	if n != len(msg) {
		sum += uint32(msg[n-2])
	}
	// 保留低16位
	sum = (sum >> 16) + (sum & 0xFFFF)
	sum += (sum >> 16)
	// 取反
	return uint16(^sum)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func readFully(conn net.Conn) ([]byte, error) {
	defer conn.Close()

	result := bytes.NewBuffer(nil)
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		result.Write(buf[0:n])
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
	}
	return result.Bytes(), nil
}
