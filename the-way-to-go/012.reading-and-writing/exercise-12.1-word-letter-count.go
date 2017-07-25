package main

import (
    "fmt"
    "bufio"
    "os"
)

func main() {
    input := bufio.NewReader(os.Stdin)

    lineCnt, wordCnt, spaceCnt := 0, 0, 0
    lineChCnt := 0

    inWord := false
    ReadLoop:
    for {
        if ch, err := input.ReadByte(); err == nil {
            switch ch {
                case 'S':
                    break ReadLoop
                case '\n':
                    inWord = false
                    lineChCnt = 0
                case ' ':
                    lineChCnt ++
                    spaceCnt ++
                    inWord = false
                default:
                    if !inWord {
                        wordCnt ++
                    }
                    inWord = true
                    lineChCnt ++
            }
            if lineChCnt == 1 {
                lineCnt ++
            }
        } else {
            fmt.Printf("Read error\n")
        }
    }
    fmt.Printf("lineCnt: %d, wordCnt: %d, spaceCnt: %d\n", lineCnt, wordCnt, spaceCnt)
}