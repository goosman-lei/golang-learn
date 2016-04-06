package main

import "fmt"

func main() {
    var (
        age int
        nickName string
    )
    _, _, age, nickName = GetName()

    fmt.Printf("age: %d, nickName: %s\n", age, nickName)

    _, _, n_age, n_nickName := GetName()

    fmt.Printf("age: %d, nickName: %s\n", n_age, n_nickName)

}

func GetName() (firstName, lastName string, age int, nickName string){
    return "May", "Chan", 18, "Chibi Maruko"
}