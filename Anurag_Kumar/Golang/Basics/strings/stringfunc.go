package main

import (
	"fmt"
	"strings"
)

func main() {

	s1 := "@@@anurag kumar hello world---"
	cutset := "@-"
	s2 := strings.Trim(s1,cutset)
	fmt.Println(s2)

	words := []string{"python", "c++", "golang", "java", "rust"}
    msg := strings.Join(words, " ")
    fmt.Println(msg)

	str := "hello world hello"
	fmt.Println(strings.Count(str,"hello")) 

}
