// In general functions only accept the number of arguments that we define. 
// In case of variadic function it can accept any number of arguments
// If the last parameter of a function definition is prefix by ellipsis(...) 
// then it can accept any number of parameters
package main

import "fmt"

func mult(nums...int) int {
	total := 1
	for _, num := range nums {
		total *= num
	}
	return total
}

func main() {
	// println is an example of variadic function
	fmt.Println(mult(1,2,3,4))
	// varidic functions can also be applied to slices
	s := []int {1,10,20}
	fmt.Println(mult(s...)) // need to pass slices alongwith ellipsis
}