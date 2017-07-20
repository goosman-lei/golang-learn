package main

import (
    "fmt"
)

type Base struct {}
func (Base) Magic() {
    fmt.Printf("Base Magic\n")
}

func (this *Base) MoreMagic() {
    this.Magic()
    this.Magic()
}

type Voodoo struct {
    Base
}

func (Voodoo) Magic() {
    fmt.Printf("Voodoo Magic\n")
}

func main() {
    v := new(Voodoo)
    v.Magic()
    v.MoreMagic()
}