package main

import "fmt"

func main() {
	var a [5]int // declare an array of length 5. Array length is part of the type
	a[0]=12 // if we will not define other element and then go will set it to zero(default)
	var b = [5]int{10,20,30,40,50} // declare and initialize 
	c := [...]int{11,22,33,44} // ... means ellipsis in go. In this case compiler figures out the array length
	fmt.Println(a) 
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(len(c)) // length of the array c is determined by the compiler 
}