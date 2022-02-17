## Fields
- Splits the string with white space as separator

## Blank Identifier
- To ignore a varibale you don't want to use
- So with Go you need to make unused variables explicit

## Break
- Terminates the for loop
- And continues with the code right after the loop

## Switch 
- Allows a variable to be tested for equality against a list of values.
- Deafault handles the case if no match is found

## Functions
- Encapsulate code into own container. Which logically belong together!
- Like variable name, you should give a function a descriptive name
- Call the function by its name, whenever you want to want execute this block of code.


## Packages
- Go program are organized into packages
- A package is a collection of Go files
- to run all the files at once `go run .` instead of `go run abc.go xyz.go`

## Exporting a variable
- Make it available for all packages in the app
- Capitalize first letter (pretty easy)

## 3 Levels of Scope
- Local
  - declaration within function
  - declaration within block eg if-else 
- package
  - declaration outside all functions & uppercase first letter. Can be used everywhere across all packages
- global
  - declaration outside all functions & uppercase first letter
  - can be used everywhere across all packages
- variable scope
  - scope is the region of a program, where a defined variable can be accessed

## Maps
- Maps unique keys to values
- You can retrieve the value by using its keys later 
- All keys have the same data type
- All values have the same data type

## struct data type
- stands for structure
- can hold mixed data types
- it's like a lightweight class, which e.g. doesn't support inheritence

## type statement - custom types
- The Type keyword creates a new type, with the name you specify

## Concurrency

## time
- the sleep function stops or blocks the current thread execution for the defined duration

## go keyword
- "go....." starts a new gorouting
- A goroutine is a lightweight thread managed by the Go runtime

## waitgroup
- Waits for the launched goroutine fo finish
- package sync provides basic synchronization functionality
- Add: sets the number of goroutines to wait for (increases the counter by the provided number)
- Wait: Blocks until the WaitGroup counter is 0 
- Done: Decrements the WaitGroup counter by 1 so this is called by the goroutine to indicate that it's finished

## advantages of goroutine
- goroutien
  - Go is using, what's called a green thread
  - Abstraction of an actual thread
  - managed by the go runtime, we are only interacting with these high level goroutines
  - cheaper & lightweight
  - you can run hundreds of thousands or millions of goroutines without affecting the performance.
  - built in functionality for goroutines to talk with each other
- os threads
  - managed by kernel
  - are hardware dependent
  - cost of these are higher
  - higher startup time
  - No easy communication between threads

# Learn Go Programming

## Control Flow
- Defer - Invoke a function but delays it execution to some future point in time.
  - it runs when current function flow is finished and before it returns
  - useful to group open and close functions together
    - be careful in loops
  - it doesn't move at the end but it moves that after the function and that fucntion actually returns 
  - run in LIFO order
  - when using variable dependency it takes the value of the variable when defer is called and not when it runs.
- Panic - occur when program cannot continue at all
  - Dob't use when file can't be opened, unless it is critical
  - Use for unrecoverable events - cannot obtain TCP port for web server 
- Recover - used to recover from panics
  - only useful in deferred functions
  - Current function will not attempt to continue, but higher funtions is call stack will