## Started Learning Go from [Golang Tutorial for Beginners](https://www.youtube.com/watch?v=yyUHQIec83I)

**Why Go, another language?**
-	Infrastructure changed a lot
  -	Scalable & Distributed
  -	Dynamic
  -	More Capacity
-	Existing programming languages did not fully take advantage of it
-	Example: Downloading, Navigation, Uploading at Parralel
-	Example: Watching, Commenting..
-	Example: Multiple users editing the same document
-	Example: Multiple users booking at the same time, prevent double boking

**Concurrency** is about dealing with lots of things at once

Built-In Concurrency Mechanism Langs: C++, Java
   - Complex code
   - Expensive and slow

Go was designed to _run on multiple cores and built to support concurrency_. <br/>
Concurrency in Go is cheap and easy.

**Main Use Case of Go:**
-	For performant applications
-	Running on scaled, distributed systems

**Characteristics of Go:**
-	Simple and readable syntax of a dynamically typed language like Python
-	Efficiency and safety of a lower-level, statically typed language like C++
-	Server-Side or Backend Language
  -	Microservices
  -	Web Applications
  -	Database Services
-	Simple Syntax
-	Fast build time, start up and run
-	Requires few resources
-	Compiled Language; Compiles into single binary (machine code)
-	Faster than interpreted language like Python
-	Consistent across different OS
-	Simplicity, Speed
-	DevOps, SRE

 
`go mod init <module path>`
- Creates a new module
- Module path can correspond to a repository you plan to publish your module to (e.g. github.com/sadab/booking-app)
- Initializes a go.mod file
- Describes the module: with name/module path and go version used in the program
- The module path is also the import path (e.g. import “github.com/sadab/booking-app”)

**NOTE:** 
1. All our code must belong to a package
2. The first statement in Go file must be “package…”
3. Execution always start from the main function

`go run <file name> =` compiles and runs the code

**Variables:**
- Variables are used to store values
- Like containers for values
- Give the variable a name & reference everywhere in the app

conferenceName  = “Go  Conference”

Welcome to out Go Conference booking applications!
Printing ticket for Go Conference…
Thank you for booking a ticket for the Go Conference!

**Benefits:**
- Update the value only once!
- Makes our app more flexible

**NOTE:** 
1. Go Compiler Errors to enforce better code quality
2. Variable names must be used
3. Imported packages must be used

**Constants** are like variables, except that their values cannot be changed.

**Print Formatted Data:** <br/>
fmt.printf(“Some text with a variable %s”, myVariable)

It takes a template string that contains the text that needs to be formatted <br/>
Plus some placeholder that tells the fmt functions how to format the variable passed in.

Example: <br/>
```
package main
import "fmt"

func main() {
   var conferenceName = "Go Conference"
   const conferenceTickets = 50
   var remainingTickets = 50
   
   fmt.Printf("Welcome to %v booking application\n", conferenceName)
   fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available")
   fmt.Println("Get your tickets here to attend.");
}
```

Here, `%v` is the format specifier and to use it we use Printf statement.






