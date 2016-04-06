package main

import "fmt"
import "os"

type PersonInfo struct {
	ID      string
	Name    string
	Address string
}

func main() {
	var expectId = os.Args[1]

	var personDB map[string]PersonInfo
	personDB = make(map[string]PersonInfo)

	personDB["12345"] = PersonInfo{"12345", "Tom", "Room 203, ..."}
	personDB["1"] = PersonInfo{"1", "Jack", "Room 301, ..."}

	fmt.Println("personDB: ", personDB)

	person, ok := personDB[expectId]
	if ok {
		fmt.Println("Found person", person.Name, "with ID: ", expectId)
	} else {
		fmt.Println("Did not find person with ID: ", expectId)
	}
}
