# Golang

Golang was developed by Google in 2007 and opensourced in 2009. With the increasing computation power like multicore processors and cloud applications, most of the programming languages does not able to take full advantage of multi-threading and make it more efficient, that’s why go is developed.

**What is Multi-Threading?**

Multi-Threading is the process of executing multiple threads simultaneously. Now what does thread means?

Thread means processes. For an instance you are using google drive to upload, navigate and download files all doing at the same time, these processes are happening together in parallel and the whole process is called multi-threading.

Now there are some challenges of multi-threading, like in a ticket booking app if there's only one ticket left and two users are booking at the same time it should book ticket of only one user not two. Go along with many other languages like C++ and JAVA has built-in concurrency support mechanism.

Still, the problem exists like the code will become more complex or it becomes more expensive and slow, These are the problem Go try to address and solve.

### Advantages of Go

Go was designed to run on multiple cores and built to support concurrency.

Concurrency on Go is cheap and easy.

Good for running on scaled and distributed systems.

Go is a combination of simplicity and readable programming language like python and Efficient and safe like C++.

Go is consistent across different operating systems.

### Data Types in Go

The two most common and basic data types are string and integers

**Arrays**:

`var nawaz = [50]string{ }`

only the same data types can be stored

**Slices:**

Its an abstraction of array

We don't need to specify the size at the beginning

It automatically expands itself with the input data

**Map:**

`var userData = make(map[string]string)` (its an empty map)

it is defined in the [key]value pair.

`userData [”firstname”] = firstname`

`userData[”lastName”]= lastName`

We cant mix different datatypes as values in Go.

**type userData struct:**

stands for structure

can hold mixed data types unlike map

```jsx
struct {

firstName string

lastName string

email string

numberOfTickets uint

isOptedForNewsletter bool

}
```

Sprintf = stores the print in a variable but didn't print unless asked.

`time.Sleep(10 * time.Second)` = sleeps for 10 seconds

go = used at the start of the function to make it concurrent in a very cheap way. It starts a goroutine, which is a lightweight thread managed by the GO runtime.

`var wg = sync.WaitGroup{}` (a package that waits for launched goroutines to finish)
