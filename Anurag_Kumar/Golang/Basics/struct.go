package main

import (
	"fmt"
)

type student struct {
	name string
	age int
}

func main() {
	std1 := student{"AnuragKumar",21}
	fmt.Printf("the name of student is %s and age is %d", std1.name, std1.age)
}