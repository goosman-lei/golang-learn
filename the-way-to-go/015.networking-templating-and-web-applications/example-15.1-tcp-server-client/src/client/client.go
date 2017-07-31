package client

import (
    "log"
    "fmt"
    "os"
    "net"
)

var logger *log.Logger

func init() {
    logger = log.New(os.Stderr, "", log.LstdFlags)
}

func ChatWith(host string, port int) (chan string, error) {
    var ch chan string
    var err error

    conn, err := net.Dial("tcp4", fmt.Sprintf("%s:%d", host, port))
    if err != nil {
        logger.Printf("client connect to %s:%d failed: %s\n", host, port, err.Error())
        return ch, err
    }

    ch = make(chan string, 10)
    go func() {
        defer func() {
            logger.Printf("client[%s => %s] close connect\n", conn.LocalAddr(), conn.RemoteAddr())
            conn.Close()
        }()

        for {
            if msg, ok := <-ch; ok {
                logger.Printf("client[%s => %s] send msg: %s\n", conn.LocalAddr(), conn.RemoteAddr(), msg)
                _, err := conn.Write([]byte(msg))
                if err != nil {
                    logger.Printf("client[%s => %s] send msg fail: %s\n", conn.LocalAddr(), conn.RemoteAddr(), err)
                    break
                }
            } else {
                logger.Printf("client[%s => %s] input channel closed\n", conn.LocalAddr(), conn.RemoteAddr())
                break
            }
        }
    }()
    return ch, err
}
