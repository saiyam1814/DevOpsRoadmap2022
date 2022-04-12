# Go

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

## Struct data type
- stands for structure
- can hold mixed data types
- it's like a lightweight class, which e.g. doesn't support inheritence

## Type statement - custom types
- The Type keyword creates a new type, with the name you specify

## Concurrency

## Time
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

## Advantages of goroutine
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

Resources:
- 