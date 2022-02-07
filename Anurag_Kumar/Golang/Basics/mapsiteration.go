package main

import "fmt"

func main() {
	students := make(map[string]int)
	students["jake"] = 45
	students["charlie"] = 39
	students["herb"] = 43
	for key,value := range students {
		fmt.Printf("the marks of %s is %d\n", key, value)
	}
}