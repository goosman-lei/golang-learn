package main

import "fmt"

type Log struct {
    msg string
}

type Customer struct {
    Name string
    log *Log
}

func main() {
    c1 := new(Customer)
    c1.Name = "Barak Obama"
    c1.log = new(Log)
    c1.log.msg = "1 - Yes we can!"
    fmt.Println(c1)

    // shorter:
    c2 :=&Customer{"Barak Obama", &Log{"1 - Yes we can!"}}
    fmt.Println(c2)

    c2.Log().Add("2 - After me the world will be a better place!")
    fmt.Println(c2.log)
    fmt.Println(c2.Log())
}

func (l *Log) Add(s string) {
    l.msg += "\n" + s
}
func (l *Log) String() string {
    return l.msg
}
func (c *Customer) Log() *Log {
    return c.log
}