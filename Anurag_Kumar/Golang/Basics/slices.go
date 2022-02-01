package main
// array holds the data itself 
// slice holds a pointer to the array that holds the data
// length -> number of element present
// capacity -> maximum element that 
import "fmt"

func main() {
	var a [5]int = [5]int {10,20,30,40,50}
	var s []int = a[1:4] // creates a slice from a[1] to a[3]
	fmt.Println(s)
	
	// we can use make keyword to make slices

	var sl []int
	sl = make([]int, 6) 
	fmt.Println(sl) // all the values is set to 0 by default
}
