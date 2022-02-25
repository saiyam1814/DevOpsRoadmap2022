package main

import (
	"fmt"
)
// Rune is an alias for int 32 in golang
// What the heck is UTF-8 encoding?
func main() {
	s := 'a'
	s_rune := rune(s)
	fmt.Println(s_rune)
	fmt.Printf("Type of s_rune is %T\n", s_rune)
	str := "Anurag"
	str_rune := []rune(str)
	fmt.Println(str_rune)
	fmt.Printf("Type of str_rune is : %T\n", str_rune)
}
