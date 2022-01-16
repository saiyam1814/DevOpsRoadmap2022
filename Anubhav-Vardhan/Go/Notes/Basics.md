``` Go
package main

import "fmt"

func main() {
	fmt.Println("Hi there!")
}

/*
	Questions:
	1. How do we run the code in our project?
		We use the go command line tool

	2. What does package main means?
	3. What does import fmt means?
	4. What is that func thing?
	5. How is the main.go file organized?
*/

/* 1
Go CLI:

1. go build		=> compiles a bunch of go source code files
2. go run		=> compiles and executes one or more files
3. go fmt		=> formats all the code in each file in the current directory
4. go install	=> compiles and "installs" a package
5. go get		=> download the raw source code of someone else's package
6. go test		=> Runs any test associated with current project
*/

/* 2
Package == Project == Workspace

Types of Packages:
1. Executable:
	Generates a file that we can run
2. Reusable
	Code used as "helpers". Good place to put reusable logic

The name of the package decides which type of package we are making. "main"
is used to make executable type package. Any other name would not execute
the package and will be reusable package only (that can be used as a dependency)
(helper code)

Every executable package must have a function inside it called main as well
*/

/* 3
import statements can be used to enable a variety of functionality from
other packages both from standard go library and packages created by other developers
Visit https://pkg.go.dev/ to find packages for go
*/

/* 4
just like functions in other languages

func: tells go we're about to declare a function
main: sets the name of the function
(): list of argumanets to pass the function
{}: function body => calling the function runs this code
*/

/* 5
- package main 					=> Package declaration
- import "fmt"				    => Import other packages that we need
- func main() {
	fmt.Println("hi there")		=> Declare funtion, tells go to do things
  }
*/
```