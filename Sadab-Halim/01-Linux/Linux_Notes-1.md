# Linux

### There are Three Major Linux Distribution :
- Debian
  - Ubuntu
    - Linux Mint
- RHEL
  - Fedora
    - Cent OS
    - Oracle Linux
- SUSE
  - Open SUSE
- Other Distributions

### Key Facts About The Debian Family:
- The Debian family is upstream for Ubuntu, and Ubuntu is upstream for Linux Mint and others.
- It uses the DPKG-based APT package manager (using apt, apt-get, apt-cache, etc.
- Ubuntu has been widely used for cloud deployments.

### Key Facts About The Red Hat Family (RHEL):
- RHEL is widely used by enterprises which host their own systems.
- It uses the yum and dnf RPM-based yum package managers.
- It supports hardware platforms such as Intel x86, Arm, Itanium, PowerPC, and IBM System z.

### Key Facts About The SUSE Family:
- SLES is widely used in retail and many other sectors.
- It uses the RPM-based zypper package manager.
- It includes the YaST (Yet Another Setup Tool) application for system administration purposes.


Linux is an open source operating system that is initally developed on and for Intel X86 based personal computers.

### Linux Terminologies:
- Kernal: Glu between hardware and application
- Distribution: Collection of software making up a linux based OS eg RHEL, Fedora, Ubunut.
- Boot Loader: Program that Boots the Operating System.
- Service: A program that runs as background process eg httpd, nfsd, ntpd, ftpd and named.
- Filesystem: Method for storing and organizing files eg ext3, ext4, FAT, XFS, NTFS, and Btrfs
- X Window System: Graphical Subsystem on nearly all linux system.
- Desktop Environment: Graphical User Interface on top of the Operating System eg GNOME, XDE, Xfce and Fluxbox.
- Command Line: Interface for typing commands on top of the Operating System.
- Shell: Command Line Interpreter that interprets the command line and instructs the Operating System to perform any necessary tasks and commands eg Bash, tcsh, zsh
- Linux Distributions: The Linux Distribution consists of the kernal plus a no. of other software tools for file related operations, user management and software package management.

### Distribution Roles:
- Linux Kernal
  - Support services: commercial, community
  - Documentation for commands Application services
  - Libraries Utilities
  - Application
  - Package update, upgrade kernal and driver patches
- Linux Support and Services
  - Commercial support and services
    - Oracle
    - RHEL
    - SLES
    - Ubuntu Commercial Support
  - Community Support and Services
    - Cent OS
    - OpenSUSE
    - Ubuntu Community Support

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

### Graphical User Interface
Wec can use either a Command Line Interface (CLI) or a Graphical User Interface (GUI) when using linux.

### X Window System
In a Linux desktop system, the X Window System is loaded as one of the final steps in the boot process. If is often just called X.

A service called the Display Manager keeps track of the displays being provided and loads the X server. The display manager also handles graphical logins and starts the appropriate desktop environment after a user logs in.

### GNOME Desktop Environment
GNOME is a popular desktop environment with an easy-to-use graphical user interface. It is bundled as the default desktop environment for most Linux distributions, including Red Hat Enterprise Linux (RHEL), Fedora, CentOs, SUSE Linux Enterprise, Ubuntu and Debian.

gnome-tweaks
There is a standard utility, gnome-tweaks, which exposes many more setting options. It also permits you to easily install extensions by external parties.

gnome-tweaks

Find the lastest modified file in `/var/log` <br/>
trash files are stored in `~/.local/share/Trash`

### Linux Package Managers

**Red Hat Package Manager (RPM):**
- `rpm -i telnet.rpm` install package
- `rpm -e telnet.rpm` uninstall package
- `rpm -q telnet.rpm` query package

**YUM:**
YUM Package Manager user RPM underneath.
- `yum isntall ansible` install package

YUM searches software repositories that acts as warehouses that cotains RPM package files. These repositores can be local, available on the internet or in cloud.

When we try to install a package using YUM, YUM searches the repositories finds the required packages and dependencies and installs all of them in the right order.

![image](https://user-images.githubusercontent.com/74575612/150744613-8aec7737-62a7-49d4-bf49-7e40ac0215b7.png)

**How does YUM find where a particular package is located?**
The information about the repository in a configuration file at path `/etc/yum.repos.d` directory.

![image](https://user-images.githubusercontent.com/74575612/150745125-eb82b26c-272e-4aec-b5a2-2f166436ee76.png)

**YUM Repos**
- `yum repolist`
![image](https://user-images.githubusercontent.com/74575612/150745332-cbd40de0-e4bc-4e3f-8824-a4b20c347acc.png)

- `yum list ansible` lists the mentioned package
- `yum remove ansible` removes the mentioned package
- `yum --showduplicates list ansible` to list all available packages
- `yum install ansible-2.4.2.0` to install a particular version of the package

### Linux Services Configuration

- `service httpd start` or `systemctl start httpd` Start HTTPD service
- `systemctl stop httpd` Stop HTTPD service
- `systemctl status httpd` Check HTTP service status
- `systemctl enable httpd` Configure HTTPD to start at startup
- `systemctl disable httpd` Configure HTTPD to not start at startup


### Debian Packaging
dpkg is the underlying package manager for these systems. It can instal, remove, and build packages. Unlike higher-level package management systems.

Package Management in the Debain Family System

For Debian-based systems, the higher-level package management system is the Advanced Package Tool (APT) system of utilities.

### Red Hat Package Manager (RPM)
Red Hat Package Manager (RPM) is the other package management system popular linux distributions. It was developed by Red Hat, and adopted by a number of other disttibutions, including SUSE/openSUSE, Mageia, CentOS, Oracle Linux, and others.

### openSUSE's YaST S oftware Mangement
The Yet another Setup Tool (YaST) software manager is similar to other graphical package managers. It is an RPM-based application.

### VI Editor
VI Editor has two modes of operations: 
- Command Mode
- Insert Mode

When we open a file in VI Editor by default we are in **Command Mode**.
In the **Command Mode** we can issue commands to the editor such as to copy and paste lines, delete lines or to naviage between lines, etc. But we cannot write content to the file in Command Mode.

To write contents to the file we must switch to **Insert Mode**.

_To Switch to Insert Mode_, type `i`

Now, we can modify the modify the file content as per our wish.

_To switch back to Command Mode from Insert Mode_, press `Esc`

### VI Editor Command Mode

![image](https://user-images.githubusercontent.com/74575612/150750561-e30121c6-bbab-4f2c-bd61-df384c52c514.png)

![image](https://user-images.githubusercontent.com/74575612/150750659-a0287b61-00e9-4798-8020-3b024c72714d.png)

- `/` to find in the text editor, for example if we want to find _of_ then, we will write the command as `/of`

To move the cursor to the rest of occurences in find, press `n` key.

# Git

### Git Cheat Sheet

**Setup:** Configuring user information used across all local repositories

- `git config --global user.name "[your_name]"` set a name 
- `git config --global user.email "[your_email]"` set an email address

**Setup & INIT:** Configuring user information, initializing and cloning repositories

- `git init` initialize an existing directory as a Git repository
- `git clone [url]` retrieve an entire repository from a hosted location via URL

**Stage & SnapShot:** Working with snapshots and the Git staging area

- `git status` show modified files in working directory, staged for your next commit
- `git add [file]` add a file as it looks now to your next commit (stage)
- `git reset [file]` unstage a file while retaining the changes in working directory
- `git diff` diff of what is changed but not staged
- `git diff --staged` diff of what is staged but not yet committed
- `git commit -m "[descriptive message]"` commit your staged content as a new commit snapshot

**Branch & Merge:** Isolating work in branches, changing context, and interchanging changes

- `git branch` list your branches. a*will appear next to the currently active branch
- `git branch [branch-name]` create a new branch at the current commit
- `git checkout` switch to another branch and check it out into the current one
- `git merge [branch]` merge the specified branch's history into the current one
- `git log` show all commits in the current branch's history

**Inspect & Compare:** Examining logs, diffs and object information

- `git log` show the commit history for the currently active branch
- `git log branchB..branchA` show the commits on branchA that are not on branchB
- `git log --follow [file]` show the commits that changed file, even across renames
- `git diff branchB..branchA` show the diff of what is in branchA that is not in branchB
- `git show [SHA]` show any object in Git in human-readable format

**Tracking Path Changes:** Versioning file removes and path changes

- `git rm [file]` delete the file from project and stage the removal for commit
- `git mv [existing-path] [new-path]` change an existing file path and stage the move
- `git log --start -M` show all commit logs with indication of any paths that moved

**Share & Update:** Retrieving updates from another repository and updating local repos

- `git remote add [alias] [url]` add a git URL as an alias
- `git fetch [alias]` fetch down all the branches form that Git remote
- `git merge [alias]/[branch]`merge a remote branch into your current branch to bring it up to date
- `git push [alias] [branch]` transmit local branch commits to the remote repository branch
- `git pull` fetch and merge any commits from the tracking remote branch

**Rewrite History:** Rewriting branches, updating commits and clearing history

- `git rebase [branch]` apply any commits of current branch ahead of specified one
- `git reset --hard [commit]` clear staging area, rewrite working tree from specfied commit

**Temporary Commits:** Temporarily store modified, tracked files in order to change branches

- `git stash` save modified and staged changes
- `git statsh list` list stack-order of staged file changes
- `git stash pop` write working from top of stash stack
- `git stash drop` discard the changes from top of stash stack

**Ignoring Patterns:** Preventing unintentional staging or commiting of files

- `log/
   *.notes
   pattern*/`
   save a file with desired pattern as .gitignore with either direct string matches or wildcard globs.

- `git config --global core.excludefile [file]` system wide ignore patterm for all local repositories


### Internet Applications

- Linux offers a wide variety of Internet applications, such as web browsers, email clients, online media applications, and others.
- Web browsers supported by Linux can be either graphical or text-based, such as Firefox, Google Chrome, Epiphany, w3m, lynx, and others.
- Linux supports graphical email clients, such as Thunderbird, Evolution, and Claws Mail, and text mode email clients, such as Mutt and mail.
- Linux systems provide many other applications for performing Internet-related tasks, such as Filezilla, XChat, Pidgin, and others. Most Linux distributions offer LibreOffice to create and edit different kinds of documents.
- Linux systems offer entire suites of development applications and tools, including compilers and debuggers.
- Linux systems offer a number of sound players including Amarok, Audacity, and Rhythmbox.
- Linux systems offer a number of movie players, including VLC, MPlayer, Xine, and Totem.
- Linux systems offer a number of movie editors, including Kino, Cinepaint, Blender among others.
- The GIMP (GNU Image Manipulation Program) utility is a feature-rich image retouching and editing tool available on all Linux distributions.
- Other graphics utilities that help perform various image-related tasks are eog, Inkscape, convert, and Scribus.

![image](https://user-images.githubusercontent.com/74575612/150855701-8f29e19a-a66f-46a3-b095-c5c947984b0e.png)



### Intro to Command Line
There is a saying, "graphical user interface make easy tasks easier, while command line interfaces make difficult tasks possible".
- No GUI overhead is incurred.
- Virtually any and every task can be accomplished while sitting at the command line.
- You can implement scripts for often-used (or easy=to=forget) tasks and series of procedures.
- You can sign into remote machines anywhere on the internet.
- You can initiate graphical applications directly from the command line instead of hunting through menus.
- While graphical tools amy vary among Linux distributions, the command line interface does not.

By default, on GNOME desktop environments, the gnome-terminal applications is used to emulate a text-mode terminal in a windows.
- xterm
- lnsole (default on KDE)
- terminatior

### Some Basic Utilities
- `cat` used to type out a file (or combine files)
- `head` used to show the first few lines of a file.
- `tail` used to show the last few lines of a file.
- `man` used to view documentation

### Sudo
`sudo` allows users to run programs usign the security privileges of another user, generally root (superuser)

### Steps for setting Up and Running sudo
If your system does not already have sudo set up and enabled,

1. At the command line prompt type su and press Enter, type your password. You should end up with a different looking prompt, often ending '#' eg su password: #
2. Now create a configuration file to enable your user account to use sudo. Typically this file is created in the /etc/sudoers.d directory with the name of the file the same as your username. eg # echo "student All=(ALL) ALL" > `/etc/sudoers.d/student`
3. Finally, some Linux distributions will complain if you do not also change permissions on the file by doing: # chmod 440 `/etc/sudoers/d/student`

### Virtual Terminals (VT)
VT are console sessions that use the entire display and keyboard outside of a graphical environment. Such terminals are considered "virtual" because, although there can be multiple active terminals, only one terminal remains visible at a time. A VT is not quite the same as a command line terminal window.

### Turning On or Off the Graphical Desktop

`sudo systemctl stop gdm`
`sudo systemctl start gdm`

