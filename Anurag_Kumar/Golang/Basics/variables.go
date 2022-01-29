package main

import "fmt"

func main() {
	var a int // main difference as compared with other language is that we define datatype later and name of the variable first
	var b int = 12 // declare and assign
	var c,d int = 10,20

	// `:=` is the assignment operator and by using this we can omit defining the datatype 
	e,f := 30,40 // compiler will figure out the data type
	fmt.Println(a,b,c,d,e,f)
}