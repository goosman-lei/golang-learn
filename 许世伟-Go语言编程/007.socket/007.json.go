package main

import "encoding/json"
import "fmt"

type Book struct {
	Title       string
	Authors     []string
	IsPublished bool
	Price       float64
	MapInfo     map[string]int
	privateCode string
}

func main() {
	book := new(Book)
	book.Title = "可复用面向对象软件的基础"
	book.Authors = []string{"Erich Gamma", "Richard Helm", "RalphJohnson"}
	book.IsPublished = true
	book.Price = 30.00
	book.MapInfo = make(map[string]int, 5)
	book.privateCode = "private code"

	book.MapInfo["one"] = 1
	book.MapInfo["two"] = 2
	book.MapInfo["three"] = 3
	book.MapInfo["four"] = 4
	book.MapInfo["five"] = 5

	jsonBytes, _ := json.Marshal(book)
	fmt.Println(string(jsonBytes))

	bookDecoded := new(Book)
	json.Unmarshal(jsonBytes, bookDecoded)
	fmt.Println(bookDecoded)

	var unknownTypeData interface{}
	json.Unmarshal(jsonBytes, &unknownTypeData)
	unknownTypeData, _ = unknownTypeData.(map[string]interface{})
	fmt.Println(unknownTypeData)
}
