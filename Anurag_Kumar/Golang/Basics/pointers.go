package main

import (
	"fmt"
)

func main() {
	a := 4
	fmt.Println("The value of a is",a)
	fmt.Println("The address of a is ",&a)
	fmt.Printf("The square of %d is %d \n",a, square(a))
	// read &a is address of a
	// * is used as the derefrencing operator 
	p := &a // p is assigned to store the address of a
	fmt.Println(p)
	fmt.Println(*p) // derefrencing

}

func square(i int) int {
	return i*i
}
