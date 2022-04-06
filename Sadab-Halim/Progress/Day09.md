# Linux

## What is Process?
A process is simply an instance of one or more related tasks (threads) executing on your computer. It is not the same as a program or a command. A sigle command may actually start several processes sumultanously. Some processes are independent of each other and others are related.

![Processes](https://courses.edx.org/assets/courseware/v1/219d348bf46fa4b3b8c83b3dbdf3fb31/asset-v1:LinuxFoundationX+LFS101x+2T2021+type@asset+block/LFS01_ch16_screen03.jpg)

## Process Types
A terminal window is a process that runs as long as needed. It allows users to execute programs and access resources in an interactive environment. You can also run programs in the background, which means they become detached from the shell.

|Process Type|Description|Example|
|:-:|:-:|:-:|
|Interactive Processes|Need to be started by a user, either a commadn line or through a graphical interface such as an icon or a menu selection|bash, firefox, top|
|Batch Processes|Automatic processes which are scheduled from and then disconnected from the terminal. These tasks are queued and work on a FIFO (First In, First Out) basis|updatedb, idconfig|
|Daemons|Server processes that run continuously, Many are launched during system startup and then wait for a user or system requrest indication that their service is requried|httpd, sshd, libvirtd|
|Threads|Lightweight processes, These are tasks that run under the umbrella of a main process, sharing memory and other resources|firefox, gnome-terminal-server|
|Kernel threads|Kernel tasks that users neither start nor terminate and have little control over. These may perform actions like moving a thread form one CPU to another|kthreadd, migration, ksoftirqd|

## Process Scheduling and States
A critical kernel function called the scheduler constantly shifts processes on and off the CPU, sharing time according to relative priority, how much time is needed and how much as already been granted to a task.

Whena a process is in so-called running state, it means it is either currently executing instructions on a CPU, or is waiting to be granted a share

Sometimes processes go into what is called a sleep state, generally when they are waiting for something to happen before they can resume.

When a process is terminating. Sometimes, a child process completes, but its parent process has not asked about its state. Amusingly such a process is said to be in a zombie state; it is not really alive, but still shows up in the system's list of processes.

## Process and Thread IDs
At any given time, there are always multiple processes being executed. The operating system keeps track of them by assigning each a unique process ID (PID) number. The PID is used to ctrack process state.

New PIDs are usually assigned in ascending order as processes are born. Thus PID 1 denotes the init process

|ID Type|Description|
|:-:|:-:|
|Process ID (PID)|Unique Process ID number|
|Parent Process ID(PPID)|Process (Parent) that started this process. If the parent dies, the PPID will refer to an adoptive parent; on recent kernels, this is kthreadd which has PPID=2|
|Thread ID (TID)|Thread ID number. This is same as the PID for single-threaded processes. For a multi-threaded process, each Thread shares the same PID, but has a unique TID.|

## Terminating a Process
To terminate a process, you can type `kill -SIGKILL <pid>` or `kill -9 <pid>`

*Only root user can kill other user process*

![Terminating a Process](https://courses.edx.org/assets/courseware/v1/6a71bd8d47df4eaf7e430d8089c632a5/asset-v1:LinuxFoundationX+LFS101x+2T2021+type@asset+block/rhelkill.png)

## User and Group IDs
Many users can access a system simultaneously, and each user can run multiple processes. The operating system identifies the user who starts the process by the Real User ID (RUID) assigned to the user.

![User and Group IDs](https://courses.edx.org/assets/courseware/v1/fbe122ffd13edf336ad978cddb953a7f/asset-v1:LinuxFoundationX+LFS101x+2T2021+type@asset+block/LFS01_ch16_screen07.jpg)

Users can be categorized into various groups. Each group is identified by the Real Group ID (RGID). The access rights of the group are determined by the Effective Group ID (EGID)

We generally talk about about UID and GID.


# Go

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


Resources:
- [Golang Tutorial for Beginners](https://www.youtube.com/watch?v=yyUHQIec83I)