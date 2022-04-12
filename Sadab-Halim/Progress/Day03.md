# Linux

**The Boot Process:** The Linux Boot Process is the procedure for initializing the system. It consists of everything that happens from when the computer power is first switched on until the user interface is fully operation.

1. Power ON
2. BIOS
3. Master Boot Record (MBR) also known as first sector of hard disk
4. Boot Loader (eg GRUB)
5. Kernal (linux OS)
6. Initial RAM Disk - initramfs image
7. /sbin/init (parent process)
8. Command Shell using getty
9. X Window System(GUI)

**BOIS:** The First Step Starting x86 based Linux System involves a number of steps. When the computer is power on the Basic Input/Output System (BIOS) initializes the hardware, including server and keyboard and tests the main memory. The process is called POST(Power On Self Test)

1. Power ON
2. BIOS (Basic Input/Output System)
3. Initializes screen keyboard and tests main memory

The Bios is stored on ROM chip on Mother Board. After this the remainder of the boot process is controlled by the Operating System(OS)

Master Boot Record (MBR) and Boot Loader Once the post is complete the system control passes from the BIOS to the Boot Loader. The boot loader is usually stored on one of the hardisks in the system, either in the boot sector or EFI parition.
* MBR
  - Partition 1 Searches for GRUB and loads into RAM
  - Partition 2
  - Partition 3

### Boot Loader In Action

The Boot Loader has two distict stages:

- For system using the BIOS/MBR method, the boot loader resides at the first sector of the hard disk, also know as Master Boot Record (MBR)
  The size of MBR is 512 Bytes. In this stage, the boot loader examines the partition table and finds a bootable partition. 
  Once it finds a bootable partition, it then searches for the second stage boot loader for example GRUB and loads it into RAM manager data to determine which UEFI application is to be launched and from where the UEFI application.

    - System
    - Bootable Hard Disk
      - Kernal
      - Inital RAM Disk
      
- The Second stage boot loader reside under /boot. 
  A splash screen is displayed, which allows us to choose which Operating System (OS) to boot. 
  After choosing the OS the boot loader loads the kernal of the selected operating system into RAM and passes control to it kernals are almost always compressed. 
  So, its first job is to uncompress itself. After this, it will check and analyze the system hardware and initialize any hardware device drivers built into the kernal.
  
### Initial RAM Disk
The initramfs filesystem image contains programs and binary files that performs all actions needed to mount porper filesystem, like providing kernal functionality for the needed filesystem and device drivers for mass storage controllers with a facility call udev (user device), which is responsible for figuring out which devices are present, locating the device drivers they to operate properly and loading them. After the root filesystem has been found, it is checked for errors and mounted.

![image](https://user-images.githubusercontent.com/74575612/150344398-69fe1901-2457-4311-abfa-5149e7a94661.png)

### Text Mode Login
Near the end of the boot process, init starts a no.of text mode login prompts. These enable you to type your username, followed by your password, and to eventually get a command shell. However, if you are running a system graphical login interface, you will not see this at first

![image](https://user-images.githubusercontent.com/74575612/150344500-7621dd96-2c41-4ea7-bace-f3a7c423cd5a.png)

### Linux Kernel
The Boot Loader loads both the kernal and initial RAM-based file system (initramfs) into memory, so it can be used directly by the kernal.

When the kernal is loaded in RAM, it immidiately initalizes and configures the computer memory and also configures all the hardware attached to the system. This includes all processors I/O subsystems, storage devices etc. The kernal also leads some mecessary user space applications.

![image](https://user-images.githubusercontent.com/74575612/150344633-0c8139a9-955d-4032-b397-dd12ee306e97.png)

### sbin/init
Once the kernal has setup all its hardware and mounted the root filesystem, the kernal runs /sbin/init. This then becomes the initial process, which then starts other processes to get the system running. Besides the system init is reponsible for keeping the system running and for shutting it down cleanly.

![image](https://user-images.githubusercontent.com/74575612/150344767-7673a20f-bb2c-4207-809b-8b567932b86e.png)

One systemd command (systemctl) is used for most basic tasks.

- $ sudo systemctl start|stop|restart httpd.service (apache web server)
- abling or disbaling a system service form starting up at system boot. sudo systemctl enable|disable httpd.service

### Linux Filesystem Basics
Different types of file systems supported by Linux:
- nventional disk filesystems: ext3, ext4, XFS, Btrfs, JFS, NTFS, vfat, exfat, etc.
- ash storage filesystems: ubifs, jffs2, yaffs, etc,
- tabase filesystems
- Special purpose filesystems: proofs, sysfs, tmpfs, squashfs, debugfs, fuse etc.

### Partitions and Filesystems
**Parition:** is a physically contiguous section of a disk, or what appears to be so in some advanced setups.
**Filesystem:** is a method of storing/finding files on a hard disk (usually in a partition).

### The FileSystem Hierarchy Standard
Linux systems store their important files according to a standard layout called the Filesystem Hierarchy standard (FHS).

Linux uses the '/' character to separate paths (unlike Windows, which uses '\'), and does not have drive letters.
![image](https://user-images.githubusercontent.com/74575612/150346586-a05bd43a-f3ac-4731-82ec-fc6e78f583ae.png)

All Linux filesystem names are case-sensitive, so `boot` `/Boot`, `/BOOT/` represent three different directories.

### Linux Basic Commands

- `ls` shows all the files and directories
- `cd` is used for changing directory
- `echo` is used to print anything on the terminal
- `clear` is used to clear the terminal
- `mkdir` is used to create a new directory
- `rmdir` is used to remove a directory
- `mv` is used to move the file _(same as copy)_
- `rm` is used to remove files
- `pwd` is used to print the current working directory
- `whoami` gives the current user 
- `sort` will sort the file
- `cat` is used to display the content of the file
- `touch` is used to create a file
- `cp srcfile destfile` is used to copy the file to destination directory
- `man` command shows the manual page of the command.
- `su` is used to switch to the root user
- `sudo` is used to access the privileges
- `sudo bash` is used to switch to the root user
- `sudo useradd demo` is used to create a user with name demo
- `sudo passwd demo` is used to set a password for the demo user
- `sudo userdel demo` is used to delete the user demo
- `sudo groupadd geeks` will create a group, a group can have multiple users.
- `sudo groupdel geeks` will delete the group
- `grep words filename` to search texts in a particular file
- `tr` it translate characters
- `df` checks the system usage
- `tail` displays content from last 
- `diff` compares file line by line
- `locate` is used to find files by name
- `find` walks a file hierarchy
- `history` shows histories of all the commands given
- `regex` is a pattern for a matching string that follows same pattern
- `alias` instructs the shell to replace one string with another string while executing the commands
- `wget` is used to download files from internet
- `top` is used to show the Linux processes
- `uname` prints basic system information
- `zip` is used to compress the files and also used as file package utlity
- `unzip` extracts all files from the specified ZIP archive to the current directory
- `hostname` used to obtain the DNS and set the system's hostname
- `lspcu` fetches the CPU architecture information
- `free` gives information about used and unused memory usage and swap memory of a system
- `vmstat` reports various bits of system information
- `id` is used to find out user, group names and numeri ID's
- `getent` helps a user get entries in a number of important text files
- `nslookup` is used to find the address record for a domain
- `netstat` is used for troubleshooting and configuration
- `sed` performs editing operations on text coming from standar input or a file
- `cut` used for cutting sections from each line of a file
- `htop` allows the user to interactivity monitor the system's vital resources or server's processes in real time
- `ps aux` is a tool to monitor processes running on our Linux systems


### Few More Linux Commands
- `chown` to change the ownership of the file/directory
- `chmod` to change the access permissions of the files/directories
   - `0` --> no permissions
   - `1` --> executes permissions
   - `2` --> writes permission
   - `4` --> reads permission
   - `7` --> reads, writes and executes all three
- `chmod 777 demo.txt` will give the readm write and execute permission to user, groups any other users.
- `lsof` list of open files and details associated with them
- `lsof -u username` to see the list of the files opened by a particular user
- `id` is used to find out user and the group names and numeric IDs of the current user or any other other
- `tar -cvf filename source-folder` is used to zip of `.tar` format
- `tar -xvf tarfile` to unzip a file.


Resources:
- [Introduction to Linux](https://www.edx.org/course/introduction-to-linux)
- [DevOps Prerequisite Course](https://www.youtube.com/watch?v=Wvf0mBNGjXY)