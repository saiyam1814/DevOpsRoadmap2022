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

## Load Averages
The load average is the average of the load number for a given period of time. It takes into account processes that are

  * Actively running on a CPU.
  * Considered runnable, but waiting for a CPU to become available.
  * Sleeping: i.e. waiting for some kind of resource (typically I/O) to become available.

  The load average can be viewed by running w, top or uptime. We will explain the numbers on the next page.

  ![Load Averages](https://courses.edx.org/assets/courseware/v1/ca05a14d78d8e3bb26b519fe65047a66/asset-v1:LinuxFoundationX+LFS101x+2T2021+type@asset+block/wuptimesuse.png)

  ## Interpreting Load Averages

  The load average is displayed using three numbers (0.45, 0.17, 0.12) is the below screenshot.

  * 0.45: For the last minute the system has been 45% utilized on average
  * 0.17: For the last 5 minutes utilization has been 17%
  * 0.12: For the last 15 minutes utilization has been 12%

  ![Interpreting Load Averages](https://courses.edx.org/assets/courseware/v1/5ad78e82ed03efc7777fad630abed5dd/asset-v1:LinuxFoundationX+LFS101x+2T2021+type@asset+block/woutputrhel.png)

  ## Background and Foreground Processes

Linux supports background and foreground job processing.

You can put a job in the backgorund by suffixing & to the command eg updatedb &.

You can either use CTRL-Z to suspend a foreground job or CTRL-C to terminate a foreground job and can always use the bg and fg commands to run a process in the background and foreground respectively.

![background and Foreground Processes](https://courses.edx.org/assets/courseware/v1/3ff2741d99789599c91efda5c5028150/asset-v1:LinuxFoundationX+LFS101x+2T2021+type@asset+block/bgfgrhel.png)

## Managing Jobs
The jobs utility displays all jobs running in ground. The display show the job ID, state, and command name, as shown here.

jobs -l provides the same information as jobs, and adds the PID of the background jobs.

The background jobs are connected to the terminal window, so, if you log off, the jobs utility will now show the ones started from that window.

![Managing Jobs](https://courses.edx.org/assets/courseware/v1/8dcff92dcec85e717944d972b96d6fcc/asset-v1:LinuxFoundationX+LFS101x+2T2021+type@asset+block/jobsrhel.png)

## The ps Command (System V Style)
ps provides information about currently running processes keyed by PID. If you want a repetitive update of this status.

![The ps Command (System V Style)](https://courses.edx.org/assets/courseware/v1/a910c16bb6f18c4d38e9ff123a6f5e02/asset-v1:LinuxFoundationX+LFS101x+2T2021+type@asset+block/ubuntupsef.png)

![ps Command (BSD Style)](https://courses.edx.org/assets/courseware/v1/8cca52331523da587fab092df4bc7dba/asset-v1:LinuxFoundationX+LFS101x+2T2021+type@asset+block/psbsdrhel.png)

## The Process Tree
pstree displays the processes running on the system in the form of a tree diagram showing the relationship between a process and its parent process and any other processes that it created. Repeated entries of a process are not displayed, and threads are displayed in curly braces.

![The Process tree](https://courses.edx.org/assets/courseware/v1/bce96558c9c9be5c152a567a0c63d392/asset-v1:LinuxFoundationX+LFS101x+2T2021+type@asset+block/ubuntupstree.png)

## top
While a static view of what the system is doing is useful, monitoring the system performance live over time is also valuable.

~[top](https://courses.edx.org/assets/courseware/v1/9eaf15c635ff33e0dbd318a36f295925/asset-v1:LinuxFoundationX+LFS101x+2T2021+type@asset+block/toprhel.png)

## First Line of the top Output 
The first line of the top output displays a quick summary of what is happening is the system, including
  * How long the system has been up
  * How many users are logged on
  * What is the load average.

## Second Line of the top Output
The second line of the top output displays the total number of processes, the number of running, sleeping, stopped, and zombie processes.

## Third line of the top Output
The third line of the top output indicates how the CPU time is being divided between the users (us) and the kernal (sy) by displaying the percentage of CPU time used for each.

## Fourth and Fifth Lines of the top output
The fourth and fifth lines of the top indicate memory usage, which is divided in two categories:
  * Physical memory (RAM) - displayes on line 4.
  * Swap space - displayed on line 5.
Both categories display total memory, used memory and free space.

## Process List of the top Output
Each line in the process list of the top output displays information about a process. By default processes are ordered by highest CPU usage.
  * Process Indentification Number(PID)
  * Process owner (USER)
  * Priority (PR) and nice values (NI)
  * Virtual (VIRT), phydical (RES) and shared memory (SHR)
  * status (S)
  * Percentage of CPU (%CPU) and memory (%MEM) used
  * Execution time (TIME+)
  * Command (COMMAND)

## Interactive Keys with top
|Command|Output|
|:-:|:-:|
|t|Display or hide summary information (rows 2 and 3)|
|m|Display or hide memory information (rows 4 and 5)|
|A|Sort the process list by top resource consumers|
|r|Renice (change the priority of) a specific processes|
|k|kill a specific process|
|f|Enter the top configuration screen|
|o|Interactively select a new sort order in the process list|

## Scheduling Future Processes Using at

![Scheduling Future Processes Using at](https://courses.edx.org/assets/courseware/v1/ec37c00269266c49e55f7a52aab93f9a/asset-v1:LinuxFoundationX+LFS101x+2T2021+type@asset+block/atout.png)

## cron
cron is a time-based scheduling utility program. It can launch routine background jobs at specific times and/or days on an on-going basis cron is driven by a configuration file called `/etc/crontab` (cron table), which contains the various shell commands that need to be run at the properly scheduled times.

typing crontab -e wil open the crontab editor to edit existing jobs or to create new jobs.

|Field|Description|Values|
|:-:|:-:|:-:|
|MIN|Minutes|0 to 59|
|HOUR|Hour field|0 to 23|
|DOM|Day of Month|1-31|
|MON|Month field|1-12|
|DOW|Day of Week|0-6(0=Sunday)|
|CMD|Command|Any command to be executed|

## sleep
sleep suspends execution for at least the specified period of time, which can be given as the number of seconds (the default)

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

## Setting Environment Variables
By default, variables created within a script are only avaible to the current shell; child processes (sub-shells) will not have access to values that have been set or modified.
Allowing child processes to see the values requires use of the export command.

|Task|Command|
|:-:|:-:|
|Show the value of a specific variables|echo $SHELL|
|Export a new variable value|export VARIABLE=value (or VARIABLE=value; export VARIABLE)|
|Add a variable permanently|Edit `~/.bashrc` and add the line export VARIABLE=value    Type source ~/.bashrc or just . ~/.bashrc (dot ~/.bashrc); or just start a new shell by typing bash|

## The HOME Variable
Home is an environment variable that represents the home (or login) directory of the user. cd without arguments will change the current working directory to the value of HOME. 

`cd $HOME` and `cd ~` are completely equivalent statements.

|Command|Explanation|
|:-:|:-:|
|`echo $HOME`|show the value of the HOME environment variable|
|pwd|where are we? Use print working directory to find out|
|cd|Change directory|
|pwd|...takes us back to HOME, as you can now see|

## The PATH Variable
PATH is an ordered list of directories (the path) which is scanned when a command is given to find to appropriate program or script to run. Each directory in path is separated by colons(:). A null (empty) directory name (or ./) indicates the current directory at any given time.

* :path1:path2 - null directory defore the first colon (:)
* path1::path2 - null directory between path1 and path2

![The PATH Variable](https://courses.edx.org/assets/courseware/v1/b955fc762a2b63221bc0a46e1fb9b419/asset-v1:LinuxFoundationX+LFS101x+2T2021+type@asset+block/setpath.png)

## The SHELL Variable
The environment variable SHELL point to the user's default command shell and contains the full pathname to the shell.

## The PS1 Variable and the command line prompt
Prompt statement (PS) is used to customize your prompt string in your terminal windows to display the information you want.

PS1 is the primary prompt variable which controls what your command line prompt look like. The following special characters can be included in PS1.

`\u` - User name
`\h` - Host name
`\w` - Current working directory
`!` - History number of this comand
`\d` - Date

They must be surrounded in single quotes when they are used.

![the ps1 variable and the command line prompt](https://courses.edx.org/assets/courseware/v1/0e41f5477fe686acda77ec04fe717d3d/asset-v1:LinuxFoundationX+LFS101x+2T2021+type@asset+block/ps1.png)


## Recalling Previous Commands
bash keeps track of previously entered commands and statements in a history buffer. You can recall previously used commands simply by using the up and down cursor keys. The information is stored in `~/.bash_history`. If you have multiple terminals open, the commands typed in each session are not saved until the session terminates.

![recalling previous commands](https://courses.edx.org/assets/courseware/v1/0dc7d17f95a156ae24e9ee8a22d98b69/asset-v1:LinuxFoundationX+LFS101x+2T2021+type@asset+block/debianhistory.png)

## Using history environment variables
* HISTFILE
  The location of the history file.
* HISTFILESIZE
  The maximum number of lines in the history file(default 500)
* HISTSIZE
  The maximum number of commands in the history file.
* HISTCONTROL
  How commands are stored
* HISTIGNORE
  which command line can be unsaved

  ![using history evironment variables](https://courses.edx.org/assets/courseware/v1/529cf8bca6c3fff74ad538be6dbab7f2/asset-v1:LinuxFoundationX+LFS101x+2T2021+type@asset+block/history.png)

## Finding and Using Previous Commands
  |Key|Usage|
  |:-:|:-:|
  |Up/Down arrow keys|Browse through the list of commands previously executed|
  |!!(pronounced as bang-bang)|Execute the previous command|
  |CTRL-R|search previously used commands. Press enter to run search command|

## Executing Previous Commands
|Syntax|Task|
|:-:|:-:|
|!|Start a history substitution|
|!$|Refer to the last argument in a line|
|!n|Refer to the nth command line|
|!string|Refer to the most recent command starting with string|

## Keyboard Shortbut
|Keyboard shortcut|task|
|:-:|:-:|
|CTRL-L|Clears the screen|
|CTRL-D|Exits the current shell|
|CTRL-Z|Puts the current process into suspended backgrond|
|CTRL-C|kills the current process|
|CTRL-H|works the same as backspace|
|CTRL-A|Goes to the begining of the line|
|CTRL-W|Deletes the word before the cursor|
|CTRL-U|Deletes from begining of line to cursor position|
|CTRL-E|goes to the end of the line|
|TAB|Auto-completes fies, directories, and binaries|

## File Ownership
In Linux and other UNIX-based operating systems, every file associated with a user who is the owner. Every file is also associated with a group which has an interest in the filea and certain rights, or permissions: read, write, and execute.

|command|Usage|
|:-:|:-:|
|chown|used to change user ownership of a file or directory|
|chgrp|used to change group ownership|
|chmod|used to change the permissions on the file, which can be done separately for owner, group and the rest of the world|

## File Permission Modes and chmod
Files have three kinds of permissions: read (r), write(w), execute(x). These are generally represented as in rwx. These permissions affect three groups of owner: user/owner(u), group(g), and others(o).

As a result you have the following three groups of three permissions:

rwx:  rwx:  rwx
u:    g:    o

There are a number of different ways to use chmod.
```
ls -l somefile
-rw-rw-r-- 1 student student 1601 Mar 9 15:04 somefile
chmod uo+x,g-w somefile
ls -l somefile
-rwxr--r-x 1 student student 1601 Mar 9 15:04 somefile
```

where u stands for user (owner), o stands for other(world), and g stands for group.

* 4 if read permission is desired
* 2 if write permission is desired
* 1 if execute permission is desired
Thus, 7 means read/write/execute, 6 means read/write, and 5 means read/execute

When you apply this to the chmod command, you have to give three digits for each degree of freedom

```
chmod 755 somefile
ls -l somefile
```

![file permission modes and chmod](https://courses.edx.org/assets/courseware/v1/5d60930deeaaca887d468867240fc6e0/asset-v1:LinuxFoundationX+LFS101x+2T2021+type@asset+block/chmodmint.png)

## Example of chown
![chown](https://courses.edx.org/assets/courseware/v1/d99d45386528584f3f861f182577fb1a/asset-v1:LinuxFoundationX+LFS101x+2T2021+type@asset+block/chownrhel7.png)

## Example of chgrp
![chgrp](https://courses.edx.org/assets/courseware/v1/2416814b8977e0048d01804a8319aace/asset-v1:LinuxFoundationX+LFS101x+2T2021+type@asset+block/chgrouprhel7.png)

## Command Line Tools for manipulating Test files
![command line tool for manipulating text files](https://courses.edx.org/assets/courseware/v1/1003ad62cadceacd75a659e16ec77873/asset-v1:LinuxFoundationX+LFS101x+2T2021+type@asset+block/cmdlinetext.png)

## cat
cat is short concatenate. It is often used to read and print files, as well as for simply viewing file contents.
`cat <filename>`

the tac command prints the lines of a file in reverse order.

|Command|Usage|
|:-:|:-:|
|cat file1 file2|Concatenate multiple files and display the output; i.e. the entire content of the first is followed by that of the second file|
|cat file1 file2 > newfile|Combine multiple files and save the output into a new file|
|cat file >> axistingfile|Append a file to the end of an existing file|
|cat > file|Any subsequent lines typed will go into the file, until CTRL-D is typed|
|cat >> file|Any subsequent lines are appended to the file, until CTRL-D is typed.|

## Using cat interactively
![using cat](https://courses.edx.org/assets/courseware/v1/2186f2162bb7f8d6f7fac9004f7d4784/asset-v1:LinuxFoundationX+LFS101x+2T2021+type@asset+block/cateoffedora.png)

Another war to create a file at the terminal is `cat> <filename> <<EOF`. A new file is created and you can type the requried input. To, exit enter EOF at the begining of a line.

*EOF is case sensitive. One can also use another word, such as STOP*

## echo
echo simply displays text. 

`echo string`

echo can be used to display a string on standard output or to place in a new file (using the > operator) or append to an already existing file (using the >> operator).

The -e option, along with the following switches, is used to enable special character sequences, such as the newline character or horizontal tab.
* \n represents newline
* \t represents horizontal tab.

echo is particularly useful for viewing the values of environment variables

|Command|Usage|
|:-:|:-:|
|echo string > newfile|The specified string is placed in a new file|
|echo string >> existringfile|The specified string is appended to the end of an already existing file|
|echo $variable|the contents of the specified environment variable are displayed|

## Working with Large Files
Viewing somefile can be done by typing either of the two following commands
```
less somefile
cat somefile | less
```

By default man pages are sent through the less command. You may have encountered the older more utility which has the same basic function but fewer capabilities i.e. less is more!

## head
head reads the first few lines of each named file (10 by default) and displays it on standard output.

eg
`head -n 5 /etc/default/grub`
you cal also just say:
`head -5 /etc/defaultgrub`
![head](https://courses.edx.org/assets/courseware/v1/a19d53c185f88f384a24c5b0ae6fe03a/asset-v1:LinuxFoundationX+LFS101x+2T2021+type@asset+block/headubuntu.png)

## tail
tail prints the last few lines of each name file and displays the last 10 lines.
tail is especially useful when you are troubleshooting any issue using log fiiles, as you probably want to see the most recent line of output.

![tail](https://courses.edx.org/assets/courseware/v1/e218e0f357f957b78420b93f6dc63aaf/asset-v1:LinuxFoundationX+LFS101x+2T2021+type@asset+block/tailubuntu.png)

## Viewing compressed files
When working working with compresses files, many standard commands cannot be used directly. For many commonly-used file and text manipulation programs, there is also a version especially designed to work directly with compressed files. These associated utilities have the letter 'z' prefixed to their name. For example we have utility programs such as zcat, zless, zdiff, and zgrep

|Command|Description|
|:-:|:-:|
|zcat compressed-file.txt.gz|to veiw a compressed file|
|zless somefile.gz or zmore somefile.gz|to page through a compresses file|
|zgrep -i less somefile.gz|to search inside a compresses file|
|zdiff file1.txt.gz|to search inside a compresses file|
|zdiff file1.txt.gz file2.txt.gz|to compare two compresses files|

*if you run zless on an uncomresses file, it will still work and ignore the decompression stage. *
*There are also quivalent utility programs for other compression methods besides gzip, for example we have bzcat and bzless associated with bzip2 and xzcat and xzless associated with xz.*

## sed and awk
It is very common to create and then repeatedly edit and/or extract contents from a file. Let's learn how to use sed and awk to easily perform such operations.

## sed
It is a powerful text processing tool and is one of the oldest, earliest and most popular UNIX utilities. It is used to modify the contents of a file or input stream, usually placing the contents into a new file or output stream. 

![sed](https://courses.edx.org/assets/courseware/v1/64bebc2555b3777d251a871e72e873d7/asset-v1:LinuxFoundationX+LFS101x+2T2021+type@asset+block/LFS01_ch12_screen_13.jpg)

sed can filter text as well as perform susbtitutions in data streams.

Data from an input source/file (or stream) is taken and moved to wroking space. The entire list of operations/modifications is applied over the data in the working space and the final contents are moved to the standard output space (or stream).

## sed command syntax
|Command|Usage|
|:-:|:-:|
|sed -e command \<filename>|specify editing commands at the command line, operate on file and put the output on standard out.|
|sed -f scripfile \<filename>|specify a scriptfile containing sed commands operate on file and put output on standard out|
|echo "I hate you" \| sed s/hate/love|use sed to filter standard input, putting output on standard out.|

![sed command syntax](https://courses.edx.org/assets/courseware/v1/43267d11e71aba5dee81d8e780b24579/asset-v1:LinuxFoundationX+LFS101x+2T2021+type@asset+block/fedorased.png)

## sed basic operations
We can perform multiple editing and filtering operations with sed.
|Command|Usage|
|:-:|:-:|
|sed s/pattern/replace_string/file|substitute first string occurence in every line|
|sed s/pattern/replace_string/g|substitute all string occurences in every line|
|sed 1,3s/pattern/replace_string/g file| susbstitute all string occurences in a range of lines|
|||
|sed -i s/pattern/replace_string/g file|save chages for string substitution in the same file|