package main

import (
	"fmt"
	"strings"
)

func main() {
	n := "charlie-jake-alan"
	w := "Anurag"
	fmt.Println(strings.Repeat(w+" ",5)) // repeats the given string 5 times
	fmt.Println(strings.Contains(n,"Anurag")) // returns the boolean value if string contains the latter value
	fmt.Println(strings.Split(n, "-")) // 
	fmt.Println(strings.ToLower(n)) // convert the string to lowercase
	fmt.Println(strings.ToUpper(n)) // convert the string to uppercase
	
}
