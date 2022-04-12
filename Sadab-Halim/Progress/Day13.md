# Go

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

Resources:
- [Golang Tutorial for Beginners](https://www.youtube.com/watch?v=yyUHQIec83I)