package main

import (
    "fmt"
    "time"
    "strings"
    "sort"
)

type Msg struct {
    filter int
    num int
}

func generate(ch chan int) {
    for i := 2; i <= 200 ; i ++ {
        ch <- i
    }
    close(ch)
}

func filter(in <-chan int, out chan<- int, prime int, dumpCh chan<- Msg) {
    for {
        if i, ok := <-in; i % prime != 0 && ok {
            dumpCh <- Msg{prime, i}
            out <- i
        } else if !ok {
            close(out)
            break;
        }
    }
}

func dumpResult(dumpCh <-chan Msg) {
    mapping := make(map[int][]int)
    for {
        if msg, ok := <-dumpCh; ok {
            _, exists := mapping[msg.filter]
            if !exists {
                mapping[msg.filter] = make([]int, 0, 10)
            }
            mapping[msg.filter] = append(mapping[msg.filter], msg.num)
        } else {
            break;
        }
    }

    var mappingKeys sort.IntSlice = make([]int, 0, len(mapping))
    for k, _ := range mapping {
        mappingKeys = append(mappingKeys, k)
    }
    mappingKeys.Sort()

    strArr := make([]string, 2)
    for _, k := range mappingKeys {
        strArr[0] = fmt.Sprintf("%s% 4d", strArr[0], k)
        strArr[1] = strArr[1] + "----"
        for i, n := range mapping[k] {
            if len(strArr) <= i + 2 {
                strArr = append(strArr, "")
            }
            strArr[i + 2] = fmt.Sprintf("%s% 4d", strArr[i + 2], n)
        }
    }
    fmt.Printf("%s\n", strings.Join(strArr, "\n"))
}

func main() {
    ch := make(chan int)
    dumpCh := make(chan Msg)
    go generate(ch)
    go dumpResult(dumpCh)

    for {
        prime, ok := <-ch
        if !ok {
            break;
        }
        ch1 := make(chan int)
        go filter(ch, ch1, prime, dumpCh)
        ch = ch1
    }
    close(dumpCh)
    time.Sleep(1e9)
}