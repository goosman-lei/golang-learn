package even

import (
    "testing"
    "log"
    "os"
)

func TestEven(t *testing.T) {
    if !Even(10) {
        t.Log("10 must be even!")
        t.Fail()
    }
    if Even(5) {
        t.Fatal("5 is not even!")
    }
    if Even(7) {
        t.Log("7 is not even!")
        t.Fail()
    }
}

func TestOdd(t *testing.T) {
    if !Odd(11) {
        t.Log("11 must be odd!")
        t.Fail()
    }
    if Odd(10) {
        t.Log("10 is not odd!")
        t.Fail()
    }
}

func BenchmarkEvenWithLog(b *testing.B) {
    logFp, err := os.OpenFile("../../log", os.O_CREATE | os.O_WRONLY | os.O_APPEND, 0755)
    if err != nil {
        b.Log("open log file failed")
        b.Fail()
    }
    logger := log.New(logFp, "", log.LstdFlags)
    for i := 0; i < b.N; i ++ {
        logger.Printf("%d => %t.\n", i, Even(i))
    }
}

func BenchmarkEven(b *testing.B) {
    for i := 0; i < b.N; i ++ {
        Even(i)
    }
}
