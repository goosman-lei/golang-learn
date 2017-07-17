package main

import (
    "fmt"
    "bytes"
    "os"
)

func main() {
    inputFp := os.Stdin
    ioBuffer := new(bytes.Buffer)
    tmpBuff := []byte{0}

    for ioBuffer.Len() < 10 {
        inputFp.Read(tmpBuff)
        ioBuffer.WriteByte(tmpBuff[0])
    }

    fmt.Printf(ioBuffer.String())
}