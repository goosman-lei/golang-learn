package main

import (
    "fmt"
    "bufio"
    "os"
    "io"
)

func main() {
    inputFile, _ := os.Open("exercise-12.7.go")
    outputFile, _ := os.OpenFile("exercise-12.7.go.write", os.O_WRONLY|os.O_CREATE, 0666)
    defer inputFile.Close()
    defer outputFile.Close()

    inputReader := bufio.NewReader(inputFile)
    outputWriter := bufio.NewWriter(outputFile)
    defer outputWriter.Flush()
    for {
        inputString, _, readerError := inputReader.ReadLine()
        if readerError == io.EOF {
            fmt.Println("EOF")
            return
        }
        outputString := string([]byte(inputString)[2:5]) + "\r\n"
        _, err := outputWriter.WriteString(outputString)
        if err != nil {
            fmt.Println(err)
            return
        }
    }
    fmt.Println("Conversion done")
}