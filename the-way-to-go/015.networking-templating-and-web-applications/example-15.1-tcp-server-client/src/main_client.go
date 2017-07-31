package main

import (
    "client"
    "fmt"
    "os"
    "bufio"
    "strings"
    "time"
)

var inputFp *bufio.Reader

func init() {
    inputFp = bufio.NewReader(os.Stdin)
}

func main() {
    ch, err := client.ChatWith("localhost", 8088)
    if err != nil {
        return
    }

    uName := prompt("What's your name:")

    for {
        msg := prompt("What to send to the server? Type Q to QUIT.")
        if msg == "Q" {
            close(ch)
            break
        }
        ch <- uName + ": " + msg
    }

    time.Sleep(1e9)
}

func prompt(msg string) string {
    fmt.Printf("%s\n", msg)
    input, err := inputFp.ReadString('\n')
    if err != nil {
        panic("input error, will shutdown\n")
    }
    return strings.TrimRight(input, "\r\n")
}
