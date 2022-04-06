# Linux

## Filesystem Varieties
Linux supports a number of native filesystem types.
* ext3
* ext 4
* squashfs
* btrfs
if also offers implementation of filesystems used on other alien operating systems, such as those from:
* Windows(ntfs,vfat)
* SGI (xfs)
* IBM (jfs)
* MacOS(hfs, hfs+)
Many older, legacy filesystems such as FAT are also supported.
The most advanced filesystem types in common use are the journaling varieties: ext4,xfs, btrfs, and jfs. These have amny state of the art features and high performance and are very hard to corrupt accidentlly.

## Linux Partitions
Each filesystem on a Linux system occupies a disk partition. Partition help to organize the contents of disks according to the kind and use of the data contained.

![Linux Partitions: gparted](https://courses.edx.org/assets/courseware/v1/6b82906abb49600a3143f0f3fd8208de/asset-v1:LinuxFoundationX+LFS101x+2T2021+type@asset+block/gparted.png)

## Mount Points
Before you can start using a filesystem, you need to mount it on the filesystem tree at a mount point. This is simply a directory where the filesystem is to be grafted on.

![Mount Points](https://courses.edx.org/assets/courseware/v1/90eea9eba0b63783a8bcf2b85ae8a9e3/asset-v1:LinuxFoundationX+LFS101x+2T2021+type@asset+block/LFS01_ch08_screen06.jpg)

*If you mount filesystem on a non-empty direcotry, the former contents of that directory are covered-up and not accessible until the filesystem is unmounted. Thus, mount points are usually empty directories.*

## Mounting and Unmounting
The mount command is used to attach a filesystem somewhere within the filesystem tree. The basic arguments are the device node and mount point.

**`sudo mount /dev/sda5 /home`**

will attach the filesystem contained in the disk partition associated with the `/dev/sda5` device node, into the filesystem tree at the `/home` mount point.

To unmount the partition
**`sudo unmount /home`**

![Mounting and unmounting](https://courses.edx.org/assets/courseware/v1/18cd65d8ee6e189efd405e7e3c890f2d/asset-v1:LinuxFoundationX+LFS101x+2T2021+type@asset+block/dfmountdebian.png)

## NFS and Network Filesystems
It is often necessary to share data across physical systems which may be either in the same location or anywhere that can be reached by the internet.

![The Client-Server Architecture of NFS](https://courses.edx.org/assets/courseware/v1/b29a567ddecc954ea6440a9e4dedd067/asset-v1:LinuxFoundationX+LFS101x+2T2021+type@asset+block/nfs.png)

Many system administrators mount remote users home directories on a server in order to give them access to the same files and configuration files across multiple client systems. This allows the users to log in to different computers, yet have access to the same files and resources.

The most common such filesystem is named simply NFS (the Network Filesystem)

## NFS on the Server
On the server machine, NFS uses daemons and other system servers are started at the command line by typing:

**sudo systemctl start nfs**

*On RHEL/CentOs 8, the service is called nfs-server, not nfs.*

The text file `/etc/exports` contains the directories and permissions that a host is willing to share with other systems over NFS. A very simple entry in this file may look like the following

**/projects *.example.com(rw)**

This entry allows the directory `/projects` to be mounted using NFS with read and write (rw) permissions and shared with other hosts in the example.com domain.

After modifying the `/etc/exports` file, you can type exportfs -av to notify linux about the directories you are allowing to be remotely mounted using NFS. You can also restart NFS with sudo systemctl restart nfs,

![NFS on the server](https://courses.edx.org/assets/courseware/v1/c9d06cf0b5114c7ff0553aae608e96bd/asset-v1:LinuxFoundationX+LFS101x+2T2021+type@asset+block/exportsnfs.png)

## NFS on the Client
On the client machine, if it is desired to have the remote filesystem mounted automatically uplon system boot, etc/fstab is modified to accomplish this. eg an entry in the client's `/etc/fstab` might look like the following

**servername:/projects /mnt/nfs/projects nfs default 0 0**

![NFS on the Client](https://courses.edx.org/assets/courseware/v1/80a14e19e05a9cfdc8b19f09d20e8e07/asset-v1:LinuxFoundationX+LFS101x+2T2021+type@asset+block/nfsclientubuntu.png)

## Overview of User Home Directories
Each user has a home directory, usually placed under `/home`. The `/root` directory on modern Linux system is no more than the home directory of the root user.

![Home Directores](https://courses.edx.org/assets/courseware/v1/9817790ba352d4047027eb1d61516db5/asset-v1:LinuxFoundationX+LFS101x+2T2021+type@asset+block/Home_directories.png)

## The /bin and /sbin Directories
The `/bin` directory contains executables binaries, essentail commands used to boot the system or in single-user mode, and essentail commands requried by all system users, such as cat, cp, ls, mv, ps and rm.

Likewise the `/sbin` directory is intended for essential binaries related to system adminimistration, such as fsck and ip.

## The /proc Filesystem
Certain filesystem, like the one mounted at `/proc` are called pseudo-filesystems because they have no permanent presence anywhere on the disk.

The `/proc` filesystem contains virtual files (files that exist only in memory) that permit viewing constantly changing kernek data. `/proc` contains files and directories that mimic kernel structures and configuration information. It does not comtain real files, but runtime system information.

Some important entries in `/proc` are: 
`/proc/cpuinfo`
`proc/interrupts`
`proc/meminfo`
`proc/mounts`
`proc/partitions`
`proc/version`

![The /proc Filessystem](https://courses.edx.org/assets/courseware/v1/5851a953799d1db46c17d156b1cd23bc/asset-v1:LinuxFoundationX+LFS101x+2T2021+type@asset+block/lsproc.png)

## The /dev Directory
The `/dev` directory contains device nodes. a type of pseudo-file used by most hardware and software devices, except for network devices. This directory is
* Empty on the disk partition when it is not mounted.
* Contains entries which are created by the udev sytem, which creates and manages device nodes on linux, creating them dynamically when devices are found.
eg
  1. `/dev/sda1` (first partition on the first hard disk)
  1. `/dev/1p1` (second printer)
  1. `/dev/random` (a source of random numbers)

  ![The /dev Directory](https://courses.edx.org/assets/courseware/v1/ee00318b4e056829ec5580f3f8c6ca10/asset-v1:LinuxFoundationX+LFS101x+2T2021+type@asset+block/lsdev.png)

  ## The /var Directory
  The `/var` directory contains files that are expected to change in size and content as the system is running (var stands for variable)
  eg
    * system log files `/var/log`
    * Package and database files: `/var/lib`
    * print queues: `/var/spool`
    * Temporary files `/var/tmp`
![The /var Directory](https://courses.edx.org/assets/courseware/v1/2d840d9232739d72bb6a2af07308a46d/asset-v1:LinuxFoundationX+LFS101x+2T2021+type@asset+block/lsvar.png)

The `/var` directory may be put on its own filesystem so that growth of the files can be accomodated and any exploding file sizes do not fatally affect the system.

![The /var Directory](https://courses.edx.org/assets/courseware/v1/948dafcdc47f674bd2c0b5c1560ebb7c/asset-v1:LinuxFoundationX+LFS101x+2T2021+type@asset+block/varfolders.png)

## Creating Aliases
You can create customized commands or modify the behavior of already existing ones by creating aliases. Most often, these aliases are placed in your `~/.bashrc` file so they are available to any command shells you create unalias removes an alias.

![creating Aliases](https://courses.edx.org/assets/courseware/v1/97491d062822787b87a74f33ea868847/asset-v1:LinuxFoundationX+LFS101x+2T2021+type@asset+block/aliassuse.png)

## Basics of Users and Groups
All Linux users are assigned a unique user ID (uid), which is just an integer; normal users start with a uid of 1000 or greater.

Linux uses groups for organizing users. Groups are collection of accounts with certain shared permissions. Control of group membership is administered through the `/etc/group` file. which shows a list of groups and their members. By default every user belongs to a default or primary group.

Users also have one or more group IDs (gid), including a default one which is the same as user ID. These numbers are associated with names through the files `/etc/passwd` and `/etc/group`. 

![basics and Users and Groups](https://courses.edx.org/assets/courseware/v1/03549a0189644137de64f426a69442c3/asset-v1:LinuxFoundationX+LFS101x+2T2021+type@asset+block/etc_group_passwd.png)

## Adding and Removing Users
Adding a new user is done with useradd and removing an existing user is done with userde1. In the simplest form, an account for the new user bjmoose would be done with:

**sudo useradd bjmoose**

which by default, sets the home directory to `/home/bjmoose` populates  it with some basic files (copied from `/etc/ske1`) and adds a line to `/etc/passwd` such as

`bjmoose:x:1002:1002::/home/bjmoose:bin/bash`

and sets the default shell to `/bin/bash`. Removing a user account is as easy as typing userdel bjmoose.

![adding and removing users](https://courses.edx.org/assets/courseware/v1/1387735f26c0ae0b377390c4c9dd9e7a/asset-v1:LinuxFoundationX+LFS101x+2T2021+type@asset+block/useradd.png)

## Adding and Removing Groups
Adding a new group is done with groupadd:

`sudo /usr/sbin/groupadd anewgroup`

the group can be removed with:

`sudo /usr/sbin/groupdel anewgroup`

Adding a user to an already group is done with usermod.

`groups rjsquirrel`
`bjmoose : rjsquirrel`

and then add the new group

`sudo /usr/sbin/usermod -a -G anewgroup rjsquirrel`
`groups rjsquirrel`
`rjsquirrel: rjsquirrel anewgroup`

These utilities update `/etc/group` as necessary. Make sure to use the -a option, for append, so as to avoid removing already existing groups. groupmod can be used to change group properties, such as the Group ID (gid) with the -g option or its name with then -n option.

Removing a user from the group is somewhat trickier. The -G option to usermod must give a complete list of groups. Thus is you do:

`sudo /usr/sbin/usermod -G rjsquirrel rjsquirrel`
`groups rjsquirrel`
`sjsquirrel : rjsquirrel`

ont the rjsquirrel group will be left

![Adding and Removing Groups](https://courses.edx.org/assets/courseware/v1/388d78b23f2b6d93f8f39cbecf6194b1/asset-v1:LinuxFoundationX+LFS101x+2T2021+type@asset+block/newgroupsuse.png)

## The root Account
The root account is very powerful and has full access to the system. Other opertaing systems often call thsi the administrator account; in linux if is often called the superuser account.

You can use sudo to assign more limited privilieges to user accounts:
* Only on a temprary basis.
* Only for a specific subset of commands.

## su and sudo
when assigning elevated privileges, you can use the command su (switch or substitute user) to launch a new shell running as another user. It is almost always a bad practise to use su to become root. Resulting errors can include deletion of vital files from the system and security breaches.

Granting privileges using sudo is less dangerous and is preferred. By default sudo must be enabled on a per-user basis

## Environment Variables
Environment variables are quantities that have specific values which may be utilized by the command shell, such as bash or other utilities and applications. Some environment variables are given preset values by the system.

![Environment](https://courses.edx.org/assets/courseware/v1/1604db32728f9bb80765461155886f79/asset-v1:LinuxFoundationX+LFS101x+2T2021+type@asset+block/envsetexport.png)

## Setting Environment Variables
By default, variables created within a script are only available to the current shell; child processes (sub-shells) will not have access to values that have been set or modified.

|task|Command|
|:-:|:-:|
|Show the value of a specific variable|echo $SHELL|
|export a new variable value|export VARIABLE=value (or VARIABLE=value; export VARIABLE)|
|add a variable|Edit `~/.bashrs` and the line `export VARIABLE=value`|
|||
|||
|||
|||

# Go

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

### Arrays in Go

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

### Slices in Go
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

Resources:
- [Golang Tutorial for Beginners](https://www.youtube.com/watch?v=yyUHQIec83I)