# DAY 13

# 27-01-2022

 1. Completed chapter 8 of Introduction to Linux.
 2. Started Chapter 9, Processes , of Introduction to Linux.









>> A process is a simple step made of related tasks that run on our system. 
>> These process require different resources, and Kernel is responsible for allocation of proper share of these resources among processes.

## DIFFERENT TYPES OF PROCESSES

**Process Type**

## Interactive Processes

Need to be started by a user, either at a command line or through a graphical interface such as an icon or a menu selection.
**Example**
**bash, firefox, top**

## Batch Processes

Automatic processes which are scheduled from and then disconnected from the terminal. These tasks are queued and work on a **FIFO**  (**F**irst-**I**n,  **F**irst-**O**ut) basis.
**Example**
**updatedb, ldconfig**

## Daemons

Server processes that run continuously. Many are launched during system startup and then wait for a user or system request indicating that their service is required.
**Example**
**httpd, sshd, libvirtd**

## Threads

Lightweight processes. These are tasks that run under the umbrella of a main process, sharing memory and other resources, but are scheduled and run by the system on an individual basis. An individual thread can end without terminating the whole process and a process can create new threads at any time. Many non-trivial programs are multi-threaded.
**Example**
**firefox, gnome-terminal-server**

## Kernel Threads

Kernel tasks that users neither start nor terminate and have little control over. These may perform actions like moving a thread from one CPU to another, or making sure input/output operations to disk are completed.
**Example**
**kthreadd, migration, ksoftirqd**
