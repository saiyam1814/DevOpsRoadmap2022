package main

import (
	"fmt"
	"encoding/json"
)

type Book struct {
	Name string
	Author string
}

func main() {
	lang := Book{Name: "harry potter", Author: "J K Rowling"}
	res, err := json.Marshal(lang)
	// Marshaling basically means encoding
	// unmarshalling means decoding. It means unpacking the JSON data into go.
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(res))
	fmt.Printf("%T\n", string(res)) // string 
	fmt.Printf("%T\n", res) // unit8
}
