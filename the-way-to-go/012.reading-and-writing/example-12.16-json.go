package main

import (
    "fmt"
    "encoding/json"
    "os"
    "log"
    "strings"
)

type Address struct {
    Type string
    City string
    Country string
}

type VCard struct {
    FirstName string
    LastName string
    Address []*Address
    Remark string
}

func (a Address) String() string {
    return fmt.Sprintf("%s %s %s", a.Type, a.City, a.Country)
}
func (v VCard) String() string {
    addrStr := make([]string, len(v.Address))
    for i, v := range v.Address {addrStr[i] = v.String()}
    return fmt.Sprintf("FirstName: %s\nLastName: %s\nRemark: %s\nAddress:\n\t%s\n",
        v.FirstName,
        v.LastName,
        v.Remark,
        strings.Join(addrStr, "\n\t"))
}

func main() {
    pa := &Address{"private", "Aartselaar", "Belgium"}
    wa := &Address{"work", "Boom", "Belgium"}
    vc := VCard{"Jan", "Kersschot&<script>", []*Address{pa, wa}, "none"}
    fmt.Printf("vcard: %v\n", vc)

    js, _ := json.Marshal(vc)
    fmt.Printf("JSON Format: %s\n", js)

    // using an encoder:
    fmt.Println("JOSN Format used encoder:")
    enc := json.NewEncoder(os.Stdout)
    enc.SetEscapeHTML(false)
    enc.SetIndent("", "  ")
    err := enc.Encode(vc)
    if err != nil {
        log.Printf("Error in encoding json: %s\n", err.Error())
    }

    var f interface{}
    err = json.Unmarshal(js, &f)
    if err != nil {
        log.Printf("Error in unmarshal json: %s\n", err.Error())
    }
    fmt.Printf("%v\n", f.(map[string]interface{}))

    var v VCard
    err = json.Unmarshal(js, &v)
    if err != nil {
        log.Printf("Error in unmarshal json: %s\n", err.Error())
    }
    fmt.Printf("%s", v)
}