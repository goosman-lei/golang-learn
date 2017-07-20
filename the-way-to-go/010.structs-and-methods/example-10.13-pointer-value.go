package main

import "fmt"

type B struct {
    thing int
}

func (b *B) change() {
    fmt.Printf("instance pointer in change(): %p\n", b)
    b.thing = 1
}

func (b B) change2() {
    fmt.Printf("instance pointer in change2(): %p\n", &b)
    b.thing = 2
}

func (b B) write() string {
    fmt.Printf("instance pointer in write(): %p\n", &b)
    return fmt.Sprint(b)
}

func main() {
    var b1 B
    fmt.Printf("b1 pointer: %p\n", &b1)
    b1.change()
    b1.change2()
    fmt.Println(b1.write())

    b2 := new(B)
    fmt.Printf("b1 pointer: %p\n", b2)
    b2.change()
    b2.change2()
    fmt.Println(b2.write())
}

/* OUTPUT:
b1 pointer: 0xc42000e258
instance pointer in change(): 0xc42000e258
instance pointer in change2(): 0xc42000e290
instance pointer in write(): 0xc42000e298
{1}
b1 pointer: 0xc42000e2c0
instance pointer in change(): 0xc42000e2c0
instance pointer in change2(): 0xc42000e2c8
instance pointer in write(): 0xc42000e2d0
{1}
 */