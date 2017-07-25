package main

import (
    "fmt"
    "strings"
)

func main() {
    fA(0)
}

func fA(level int) {
    fmt.Printf("%senter fA\n", strings.Repeat("  ", level))
    defer fmt.Printf("%sdefer in fA before fB()\n", strings.Repeat("  ", level))
    fB(level + 1)
    defer fmt.Printf("%sdefer in fA after fB()\n", strings.Repeat("  ", level))
    fmt.Printf("%sexit fA\n", strings.Repeat("  ", level))
}

func fB(level int) {
    fmt.Printf("%senter fB\n", strings.Repeat("  ", level))
    defer fmt.Printf("%sdefer in fB before recover\n", strings.Repeat("  ", level))
    defer func() {
        fmt.Printf("%sdefer in fB in recover\n", strings.Repeat("  ", level))
        recover()
    }()
    defer fmt.Printf("%sdefer in fB after recover\n", strings.Repeat("  ", level))
    defer fmt.Printf("%sdefer in fB before fC()\n", strings.Repeat("  ", level))
    fC(level + 1)
    defer fmt.Printf("%sdefer in fB between fC() and fD()\n", strings.Repeat("  ", level))
    fD(level + 1)
    defer fmt.Printf("%sdefer in fB between fD() and fE()\n", strings.Repeat("  ", level))
    fE(level + 1)
    defer fmt.Printf("%sdefer in fB after fE()\n", strings.Repeat("  ", level))
    fmt.Printf("%sexit fB\n", strings.Repeat("  ", level))
}
func fC(level int) {
    fmt.Printf("%senter fC\n", strings.Repeat("  ", level))
    defer fmt.Printf("%sdefer in fC\n", strings.Repeat("  ", level))
    fmt.Printf("%sexit fC\n", strings.Repeat("  ", level))
}

func fD(level int) {
    fmt.Printf("%senter fD\n", strings.Repeat("  ", level))
    defer fmt.Printf("%sdefer in fD before recover\n", strings.Repeat("  ", level))
    defer func() {
        fmt.Printf("%sdefer in fD in recover\n", strings.Repeat("  ", level))
        recover()
    }()
    defer fmt.Printf("%sdefer in fD after recover\n", strings.Repeat("  ", level))
    defer fmt.Printf("%sdefer in fD before panic\n", strings.Repeat("  ", level))
    panic("panic in fD\n")
    defer fmt.Printf("%sdefer in fD after panic\n", strings.Repeat("  ", level))
    fmt.Printf("%sexit fD\n", strings.Repeat("  ", level))
}

func fE(level int) {
    fmt.Printf("%senter fE\n", strings.Repeat("  ", level))
    defer fmt.Printf("%sdefer in fE before panic\n", strings.Repeat("  ", level))
    panic("panic in fE\n")
    defer fmt.Printf("%sdefer in fE after panic\n", strings.Repeat("  ", level))
    fmt.Printf("%sexit fE\n", strings.Repeat("  ", level))
}

/*
enter fA
  enter fB
    enter fC
    exit fC
    defer in fC
    enter fD
    defer in fD before panic
    defer in fD after recover
    defer in fD in recover
    defer in fD before recover
    enter fE
    defer in fE before panic
  defer in fB between fD() and fE()
  defer in fB between fC() and fD()
  defer in fB before fC()
  defer in fB after recover
  defer in fB in recover
  defer in fB before recover
exit fA
defer in fA after fB()
defer in fA before fB()
*/