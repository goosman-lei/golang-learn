package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
)

func main() {
    inputFile, inputErr := os.Open("exercise-12.1-word-letter-count.go")
    if inputErr != nil {
        fmt.Printf("An error occurred on opening the inputFile\n" + "Does the file exists?\n" + "Have you got access to it?\n")
        return
    }
    defer inputFile.Close()

    inputReader := bufio.NewReader(inputFile)
    for {
        inputString, readerError := inputReader.ReadString('\n')
        if readerError == io.EOF {
            return
        }
        fmt.Printf("> %s", inputString)
    }
}