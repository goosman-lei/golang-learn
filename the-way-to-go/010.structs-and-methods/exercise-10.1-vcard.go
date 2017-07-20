package main

import (
    "fmt"
    "bytes"
)

type Address struct {
    address string
}

type VCard struct {
    name string
    addresses []*Address
    birthday string
    photo string
}

func main() {
    add1 := &Address{"BeiJing City, ChangPing District, XiSanQi Street, QiShengJiaYuan"}
    add2 := &Address{"BeiJing City, ChaoYang District, DongSanHuan Street, Center Of TianYuanGang"}
    vcard := &VCard{name: "Guoguo", birthday: "19860911", photo: "http://ice.tec-inf.com", addresses: []*Address{add1, add2}}

    fmt.Printf("%s", vcard)
}

func (this *VCard)String() string{
    var buffer bytes.Buffer
    buffer.WriteString(fmt.Sprintf("Name: %s\nBirthday: %s\nPhotoUrl: %s\n", this.name, this.birthday, this.photo))

    if len(this.addresses) > 0 {
        buffer.WriteString("Address:\n")
    }
    for idx, address := range this.addresses {
        buffer.WriteString(fmt.Sprintf("\t%d: %s\n", idx, address.address))
    }
    return buffer.String()
}