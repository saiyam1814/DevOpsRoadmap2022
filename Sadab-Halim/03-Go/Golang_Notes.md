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

### Data Types in Go
- Strings
  - For **textual data**, defined with double quotes, e.g. _"This is a string"_
- Integers
  - Representing **whole numbers**, positive and negative, e.g.: _5, 120, -20_

Each data type can do different things and behaves differently. <br/>

**NOTE:**
1. Go is a statically typed language we need to tell Go Compiler, the data type while declaring the variable
2. Type Inference: BUT, Go can infer the type when we assign a value

Example-1
```
var userName string
var userTickets int
//ask user for their name

username = "Tom"
userTickets = 2
fmt.Printf("User %v booked %v tickets.\n, userName, tickets)
```

Example-2
```
func main(){
  conferenceName := "Go Conference"
  const conferenceTickets = 50
  var remainingTickets = 50
  
  fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n")
}
```

Here, `%T` refers the type of value.

### Getting User Input

**"fmt" package**
Different functions for Formatted Input and Output (I/O)
- Print Messages
- Collect User Input
- Write into a File

```
var firstName string
var lastName string
var email string
var userTickets int

fmt.Println("Enter your first name: ")
fmt.Scan(&firstName)


fmt.Println("Enter your last name: ")
fmt.Scan(&lastName)


fmt.Println("Enter your email address: ")
fmt.Scan(&email)


fmt.Println("Enter number of tickets: ")
fmt.Scan(&userTickets)

fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v", firstName, lastName, userTickets, email)

```

**What is a Pointer?**

A pointer is a variable that points to the memory address of another variable. <br/>
It is also known as _specific variable_ in Go.

![image](https://user-images.githubusercontent.com/74575612/150790609-bb6376d1-0b7a-4338-8f09-07d2176b44a2.png)

![image](https://user-images.githubusercontent.com/74575612/150790764-1c59f9fc-a345-459f-b31c-8601c5acf92b.png)


### Book Ticket Logic
```
package main
import "fmt"

func main() {
  conferenceName := "Go Conference"
  const conferenceTickets int = 50
  var remainingTickets uint = 50
  
  fmt.Printf("Welcome to %v booking application\n", conferenceName)
  fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
  fmt.Println("Get your tickets here to attend.")
  
  var firstName string
  var lastName string
  var email string
  var userTickets uint
  
  fmt.Println("Enter your first name: ")
  fmt.Scan(&firstName)

  fmt.Println("Enter your last name: ")
  fmt.Scan(&lastName)

  fmt.Println("Enter your email address: ")
  fmt.Scan(&email)

  fmt.Println("Enter number of tickets: ")
  fmt.Scan(&userTickets)
  
  remainingTickets = remainingTickets - userTickets

  fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v", firstName, lastName, userTickets, email)
  fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
```

### Arrays in Go

- Data structures to store collection of elements in a single variable
- We want to store the entered user data in some kind of a list
- In array only the same data type can be stored
```
var bookings [50]string
bookings[0] = firstName + " " + lastName
```

```
fmt.Printf("The whole array: %v\n", bookings)
fmt.Printf("The first value: %v\n", bookings[0])
fmt.Printf("Array type: %T\n", bookings)
fmt.Printf("Array length: %v\n", len(bookings))
```

### Slices in Go
- Slice is an abstraction of an Array
- SLices are more flexible and powerful:
  variable-length or get a sub-array of its own
- Slices are also index-based and have a size, but is resized when needed.

append() adds the element(s) at the end of the slice
Grows the slice if a greater capacity is needed and returns the updated slice value.

```
var bookings []string
bookings = append(bookings, firstName + " " + lastName)
```

```
fmt.Printf("The whole slice: %v\n", bookings)
fmt.Printf("The first value: %v\n", bookings[0])
fmt.Printf("Slice type: %T\n", bookings)
fmt.Printf("Slice length: %v\n", len(bookings))
```

### Loops in Go

A loop statement allows us to execute code multiple times, in a loop <br/>
Loops are simplified in Go <br/>
We only have the _"for loop"_

1. Infinite Loop
```
for {
  
}
```
To come out from for loop Enter: `Ctrl + C`

2. For-Each Loop; _Iterating over a list_
**range** iterates over elements for difference data structures

**strings.Fields()**: splits the string with white space as separator
And returns a slice with the split elements

```
firstNames := []string{}
for index, booking := range bookings {
   var names = strings.Fields(booking)
   firstNames = append(firstNames, names[0])
}

fmt.Printf("The first names of bookings are: %v\n", firstNames)
```

**NOTE:**
- Instead of index we can also use **Blank Identifier: _** 
- It is used when we want to ignore a variable
- So with Go we need to make unused variables explicit

Example:
```
firstNames := []string{}
for index, booking := range bookings {
   var names = strings.Fields(booking)
   firstNames = append(firstNames, names[0])
}

fmt.Printf("The first names of bookings are: %v\n", firstNames)
```


## Functions in Go
- Encapsulate code into own container (=function), which logically belongs together.
- Like variable name, we give a function a descriptive name
- Call the function by it's name, whenever we want to execute this block of code.
- Every program has atleas one function, which is the main function
- Function is only executed, when called.
- We can call a function as many times as we want
- A function is also used to reduce code duplication

```
func funName { //function name
   fmt.Println("This is inside function funName.")
}

funcName() // function call
```

Example
```
func main(){
  greetUsers()
}

greetUsers() {
  fmt.Println("Welcome to our conference.")
}
```

**Function Parameters:**
- Information can be passd into function as parameters
- Parameters are also called as arguments

**Return Values:**
- A function can return data as a result
- So a function can take an input and return an output
- In Go we have to define the input and output parameters including its type explicitly
- A Go function can return multiple values

**Package Level Variables:**
- Defined at the top outside all functions
- They can be accessed inside any of the function
- And in all files, which are in the same package

Example
```
func main(){
  greetUsers(conferenceName)
}

greetUsers(confName string) {
  fmt.Println("Welcome to "%v" booking application.\n", confName)
}
```

Book Ticket Logic Example
```
import {
  "fmt"
  "strings"
}

const conferenceTickets int = 50
conferenceName = "Go Conference"
var remainingTickets uint = 50
bookings = []string{}

func main(){
  
  
  greetusers()
  
  for {
    
    firstName, lastName, email, userTickets := getuserInput()
    isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)
    
    if isValidName && isValidEmail && isValidTicketNumber {
      
    bookTicket (userTickets, firstName, lastName, email)
    
    firstNames := getFirstNames()
    fmt.Printf("The first names of bookings are: %v", firstNames)
    
    if remainingTickets == 0 {
      fmt.Println("Our conference is booked out. Come back next year.")
      break
    } else {
        if !isValidName {
          fmt.Println("first name or last name you entered is too short.")
        }
        if !isValidEmail {
          fmt.Println("email address you entered doesn't contain @ sign")
        }
        if !isValidTicketNumber {
          fmt.Println("number of tickets you entered is invalid")
        }
      }
   }
}

func greetUsers() {
   fmt.Printf("Welcome to %v booking applications\n", conferenceName)
   fmt.Printf("We have total of %v tickets and $v are still available.\n", conferenceTickets, remainingTickets)
   fmt.Println("Get your tickets here to attend.")
}

func getFirstNames() []string {
  firstNames := []string{}
    for _, booking := range bookings {
      var name = strings.Fields(booking)
      firstNames = append(firstNames, names[0])
    }
    return firstNames
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
  isValidName := len(firstName) >= 2 && len(lastName) >= 2
  isValidEmail := strings.Contains(email, "@")
  isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
  return isValidName, isValidEmail, isValidTicketNumber
}

func getUserInput() (string, string, string, uint) {
    var firstName string
    var lastName string
    var email string
    var userTickets uint

    fmt.Println("Enter your first name: ")
    fmt.Scan(&firstName)

    fmt.Println("Enter your last name: ")
    fmt.Scan(&lastName)

    fmt.Println("Enter your email address: ")
    fmt.Scan(&email)

    fmt.Println("Enter number of tickets: ")
    fmt.Scan(&userTickets)
    
    return firstName, lastName, email, userTickets
}

func bookTicket (userTickets uint, firstName string, lastName string, email string) {
  remainingTickets = remainingTickets - userTickets
  bookings = append(bookings, firstName + " " + lastName)
      
      
  fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v", firstName, lastName, userTickets, email)
  fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
```

**Local Variables:**
- Defined inside a function or a block
- They can be accessed only inside that function or block of code

**NOTE:**
- Best Practice: Define variable as "local" as possible
- Create the variable where you need it

**More Use Cases of Functions:**
- Group logic that belongs together
- Reuse logic and so reducing duplication of code

## Packages in Go
- A package is a collection of Go files

Example: Filename: helper.go
```
package main

func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
  isValidName := len(firstName) >= 2 && len(lastName) >= 2
  isValidEmail := strings.Contains(email, "@")
  isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
  return isValidName, isValidEmail, isValidTicketNumber
}
```

