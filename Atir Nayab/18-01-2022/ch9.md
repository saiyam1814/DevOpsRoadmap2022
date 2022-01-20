# Processes

**Main Highlights**
* Describe what a process is and distinguish between types of processes.
* Enumerate process attributes
* Understand the use of laod averages and other process metrics.
* Manipulate processes by putting them in background and restoring them to foreground.
* Use at, cron, and sleep to schedule processes in the future or pause them.

## What is Process?
A process is simply an instance of one or more related tasks (threads) executing on your computer. It is not the same as a program or a command. A sigle command may actually start several processes sumultanously. Some processes are independent of each other and others are related.

![Processese](https://courses.edx.org/assets/courseware/v1/219d348bf46fa4b3b8c83b3dbdf3fb31/asset-v1:LinuxFoundationX+LFS101x+2T2021+type@asset+block/LFS01_ch16_screen03.jpg)

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