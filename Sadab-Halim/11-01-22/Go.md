## Completed Go from [Golang Tutorial for Beginners](https://www.youtube.com/watch?v=yyUHQIec83I)

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

### Arrays & Slices

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

**Slices in Go**
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
