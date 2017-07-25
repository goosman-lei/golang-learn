package main

import (
    "fmt"
    "os"
    "flag" // command line options parser
)

var NewLine = flag.Bool("n", false, "print on new line")
var T = flag.Bool("t", false, "test option t")

const (
    Space = " "
    Newline = "\n"
)

func main() {
    flag.PrintDefaults()
    flag.Parse()

    fmt.Printf("NewLine %v\n", *NewLine)
    fmt.Printf("T %v\n", *T)

    var s string = ""
    for i := 0; i < flag.NArg(); i ++ {
        fmt.Printf("Arg[%d]: %s\n", i, flag.Arg(i))
        if i > 0 {
            s += Space
        }
        s += flag.Arg(i)
    }

    if *NewLine {
        s += Newline
    }

    os.Stdout.WriteString(s)
}