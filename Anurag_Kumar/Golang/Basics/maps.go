// it is similar to dictionary in python
// It is used for the purpose of storing key-value pairs
// syntax
// make(map[datatype of key]datatype of value)

// for demonstration we will create a map of students with their marks

package main

import "fmt"

func main() {
	students := make(map[string]int)
	fmt.Println(students) // map[] as nothing is assigned now
	students["jake"] = 45
	students["charlie"] = 39
	students["herb"] = 43
	fmt.Println(students)
	// also possible to initialise along with declration
	pens := map[string]int {
		"cello": 10,
		"linc": 20,
	}
	fmt.Println(pens)
}
