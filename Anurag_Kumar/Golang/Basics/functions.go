package main

import "fmt"

// function syntax 
// func <function name> return type {
//	function body
// }

func addition(a int, b int) int {
	return a+b
}

func subtract(a int, b int) int {
	return a-b
}

func multiplication(a int, b int) int {
	return a*b
}

func division(a int, b int) int {
	return a/b
}

func main() {
	fmt.Println(addition(2,5))
	fmt.Println(subtract(4,2))
	fmt.Println(multiplication(5,6))
	fmt.Println(division(9,3))
	fmt.Println("Exited")
}