## Pointers

**Agenda**
- creating pointers
- dereferencing pointers
- the new function
- working with nil
- types with internal pointers


```go
var a int = 42
var b *int = $a
fmt.Println(a, *b)
```
Here b is the pointer storing the address of a 

```go
func main(){
  var ms *myStruct
  ms = new(myStruct)
  (*ms).foo = 42
  fmt.Println((*ms).foo)

  // same 
  ms.foo = 42
  fmt.Println(ms.foo)
}
type myStruct struct{
  foo int
}
```

- pointer types use an asterisk(*) as a prefix to type pointed to 
  - *int -a pointer to an integer
- use addressof operator before initializer
   - &myStruct{foo: 42}
- Use the new keyword
  - can't initialize fields at the same time
- types with internal pointers
  - all assignment operations is go are copy operations
  - slices and maps contain internal pointers, so copies point to same underlying data


## function
- basic syntax
```go
  func foo(){

  }
```
- parameters
  - comma delimited list of variables and types
    - func foo(bar string, baz int)
  - Parameters of same type list type once
    - func foo(bar, baz, int)
  - when pointer are passed in the function can change the value in the caller 
    - This is always true for data of slices and maps 
  - Use variadic parameters to send list of same types in 
    - must be last parameter
    - Received as a slice
    - func foo (bar string, baz ...int)
- return values
  - single return values just list type
    - func foo() int
  - Multiple return value list surrounded by  parantheses
    - func foo() (int, error)
    - The (result type, error) paradigm is a very common idiom 
  - Can use named return values
    - initializes returned variable
    - Return using return keyword on its own
  - Can return addresses of local variables
    - Automatically promoted from local memory (stack) to shared memory (heap)
- anonymous functions
  - functions don't have names if they are: 
    - immidiately invoked
    - func(){}()
    - assigned to a variable or passed as an argument to a function
    - a := func(){}; a()
- functions as types
  - can assign functions to variables or use as arguments and return values in functions
  - type signature is like function signature, with no parameter names
    - var f func(string, string, int)(int, error)

- Methods
  - function that executes in context of a type
  - format
    - func (g greater) greet() {
      ...
    }
    - Receiver can be value or pointer
      - value receiver gets copy of type
      - pointer receiver gets pointer to type

## interfaces

**Agenda**
- Basics
- composing interfaces
- type conversion
  - the empty interface
  - type switches
- implementing with values vs. pointers

eg 
```go
func main(){
  var w Writer = ConsoleWriter{}
  w.Write([]byte("Hello Go!"))
}
type Writer interface(){
  // struct store data types whereas interface store method definition
  // interface don't describe data it describe behaviour
  Write([]byte) (int, error)
  type ConsoleWriter struct {}
  func (cw ConsoleWriter) Write(data []byte) (int, error){
    n, err := fmt.Println(string[data])
    return n,err
  }
}
```
- best practices
  - Use manu small interface
    - io.Writer, io.Reader,interface{}
  - Don't export interfaces for types that will be consumed
  - Do export interface for types that will be used by package
  - Design functions and methods to receive interfaces whenever possible 

## goroutine
- creating gorouting
- synchronization
  - waitgroups
  - mutexes
- parallelism
- best practices

- Don't create goroutines in libraries
  - let consumer control concurrency
- when creating a goroutine know how it will end
  - Avoids subtle memory leaks
- check for race conditions at compile time

- summary
  - creating goroutines
    - use go keyword in front of function call
    - when using anonymous functions, pass data as local variables
  - synchromization
    - use sync.Waitgroup to wait for groups of goroutines to complete
    - use sync.Mutex and sync.RWMutex to protext data access
  - parallelism
    - by default, go will use CPU threads, equal to available cores
    - change with runtime.GOMAXPROCS
    - More threads can increase performance, but too many can slow it down

## Channel
 
**Agenda**
- channel basics
- restricting data flow
- buffered channels
- closing channels
- for...range loops with channels
- select statement

eg 
```go
package main()

import (
  "fmt"
  "sync"
)
var wg = sync.WaitGroup{}

func main(){
  ch := make(chan int)
  wg.Add(2)
  go func(){
    i := <- ch
    fmt.Println(i)
    wg.Done()
  }()
  go func(){
    ch <- 42
    wg.Done()
  }()
  wg.wait() 
}
```
- create a channel with make command
  - make(chan int)
- send message into channel
  - ch <- val
- Receive message from channel
  - val := <- ch
- can have multiple senders and receivers
-restricting data flow
  - channel can be cast into send only or receive only versions
  - send-only: chan <- int
  - reveive-only: <- chan int
- buffer channels
  - channels block sender side till receiver is available
  - block receiver side till message is availble
  - can decouple sender and receiver with buffered channels
    - make(chan int, 50)
  - use buffered channels when send and receiver have assymetric loading
- for...range looops with channels
  - use to monitor channel and process messages as they arrive
  - loop exits when channel is closed
- Select statements
  - allows goroutine to monitor several channels at once
  - blocks if all channels block
  - if multiple channels receive value simultaneously, behavior is undefined