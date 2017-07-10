package main

import "fmt"

func main() {
    season(0)
    season(3)
    season(6)
    season(9)
    season(12)
    season(13)
}

func season(month int) {
    switch month {
        case 1, 2, 3:
            fmt.Printf("%dth month is sprint\n", month)
        case 4, 5, 6:
            fmt.Printf("%dth month is summer\n", month)
        case 7, 8, 9:
            fmt.Printf("%dth month is autumn\n", month)
        case 10, 11, 12:
            fmt.Printf("%dth month is winter\n", month)
        default:
            fmt.Printf("%dth month not exists\n", month)
    }
    return
}