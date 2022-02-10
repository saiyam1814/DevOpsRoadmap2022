package main

import "fmt"

type student struct {
	name string 
	age int
}

func (s student) details() {
	fmt.Println("the name is", s.name)
	fmt.Println("the age is", s.age)
}

func main()  {
	jake := student { 
		name: "Jake Harper", // mind it that we are not using '=' here.
		age: 21,
	}
	jake.details()
}


// for analogy I am writing the python code below 


// class student:
//     def __init__(self, name, age):
//         self.name = name
//         self.age = age

//     def details(self):
//         print("the name is "+self.name)
//         print("the age is ", self.age)

// s1 = student("Anurag",21)
// print(s1.details())
