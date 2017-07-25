package even

import (
    "os"
    "log"
)

func Even(i int) bool {
    return i % 2 == 0
}

func Odd(i int) bool {
    return i % 2 != 0
}

func UdLog(i int) bool {
    logFp, err := os.OpenFile("../../log-frame", os.O_CREATE | os.O_WRONLY | os.O_APPEND, 0755)
    if err != nil {
        return false
    }
    logger := log.New(logFp, "", log.LstdFlags)
    logger.Printf("Index %d\n", i)
    return true
}
