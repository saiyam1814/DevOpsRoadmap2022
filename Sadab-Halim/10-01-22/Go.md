## Continued Learning Go from [Golang Tutorial for Beginner](https://www.youtube.com/watch?v=yyUHQIec83I)

### Data Types in Go
- Strings
  - For **textual data**, defined with double quotes, e.g. _"This is a string"_
- Integerrs
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



